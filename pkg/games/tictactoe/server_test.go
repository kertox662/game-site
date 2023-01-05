package tictactoe

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	tttproto "github.com/kertox662/game-site/proto/games/tictactoe"
	"github.com/kertox662/game-site/proto/games/tictactoe/tictactoeconnect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeMockServerAndClient() (*Server, tictactoeconnect.TicTacToeServiceClient) {
	tttServer := NewServer()
	path, handler := tictactoeconnect.NewTicTacToeServiceHandler(tttServer)
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	server := &http.Server{
		Addr:              "localhost:8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	mock := newMockClient(server, tttServer)
	client := tictactoeconnect.NewTicTacToeServiceClient(mock, "http://localhost:8080")
	return tttServer, client
}

type mockClient struct {
	tttSrv *Server
	server *http.Server
}

func newMockClient(server *http.Server, tttSrv *Server) *mockClient {
	return &mockClient{
		tttSrv: tttSrv,
		server: server,
	}
}

func (c *mockClient) Do(req *http.Request) (*http.Response, error) {
	rec := httptest.NewRecorder()
	c.server.Handler.ServeHTTP(rec, req)
	return rec.Result(), nil
}

func TestServerCreateGame(t *testing.T) {
	assert := assert.New(t)
	tttSrv, client := makeMockServerAndClient()

	tests := []struct {
		name            string
		expectedGame    *game
		expectedErrCode connect.Code
	}{
		{
			name: "TestGame1",
			expectedGame: &game{
				maxPlayers:    2,
				playerCount:   2,
				board:         newBoard(3),
				boardSize:     3,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
		},
		{
			name: "TestGame2",
			expectedGame: &game{
				maxPlayers:    5,
				playerCount:   5,
				board:         newBoard(5),
				boardSize:     5,
				connectTarget: 4,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
		}, {
			name: "NonPositiveBoardSize",
			expectedGame: &game{
				maxPlayers:    1,
				connectTarget: 1,
				boardSize:     0,
			},
			expectedErrCode: connect.CodeInvalidArgument,
		}, {
			name: "NonPositiveConnectTarget",
			expectedGame: &game{
				maxPlayers:    1,
				connectTarget: 0,
				boardSize:     1,
			},
			expectedErrCode: connect.CodeInvalidArgument,
		}, {
			name: "NonPositiveMaxPlayers",
			expectedGame: &game{
				maxPlayers:    0,
				connectTarget: 1,
				boardSize:     1,
			},
			expectedErrCode: connect.CodeInvalidArgument,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expectedGame := test.expectedGame
			req := &tttproto.CreateGameRequest{
				MaxPlayers:    int32(expectedGame.maxPlayers),
				BoardSize:     int32(expectedGame.boardSize),
				ConnectTarget: int32(expectedGame.connectTarget),
			}

			resp, err := client.CreateGame(context.Background(), &connect_go.Request[tttproto.CreateGameRequest]{
				Msg: req,
			})
			if err != nil {
				if test.expectedErrCode == 0 {
					t.Fatal(err)
				}
				connErr := &connect.Error{}
				require.ErrorAs(t, err, &connErr)
				assert.Equal(test.expectedErrCode, connErr.Code())
				return
			}

			gameId := resp.Msg.GetGameId()
			assert.Contains(tttSrv.games, gameId)
			assert.Equal(tttSrv.games[gameId], expectedGame)
		})
	}
	assert.Equal(len(tttSrv.games), 2)
}

func TestServerGetGame(t *testing.T) {
	assert := assert.New(t)
	tttSrv, client := makeMockServerAndClient()
	testBoard := board([][]int{
		{1, 2, 1, 1},
		{2, 3, 3, 2},
		{3, 2, 1, 3},
		{2, 3, 1, 1},
	})

	tests := []struct {
		name            string
		id              string
		game            *game
		expectedPlayers []string
		expectedErrCode connect.Code
	}{
		{
			name: "TestGame1",
			id:   "1",
			game: &game{
				maxPlayers:    2,
				playerCount:   2,
				board:         newBoard(3),
				boardSize:     3,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
			expectedPlayers: []string{"1", "2"},
		},
		{
			name: "TestGame2",
			id:   "2",
			game: &game{
				maxPlayers:    5,
				playerCount:   5,
				board:         testBoard,
				boardSize:     4,
				connectTarget: 4,
				currentTurn:   1,
				winner:        3,
			},
			expectedPlayers: []string{"1", "2", "3", "4", "5"},
		},
		{
			name:            "NonExistentGame",
			id:              "3",
			expectedErrCode: connect.CodeNotFound,
		},
	}

	for _, test := range tests {
		if test.game != nil {
			tttSrv.games[test.id] = test.game
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := &tttproto.GetGameDataRequest{
				GameId: test.id,
			}

			resp, err := client.GetGameData(context.Background(),
				&connect_go.Request[tttproto.GetGameDataRequest]{Msg: req},
			)
			if err != nil {
				if test.expectedErrCode == 0 {
					t.Fatal(err)
				}
				connErr := &connect.Error{}
				require.ErrorAs(t, err, &connErr)
				assert.Equal(test.expectedErrCode, connErr.Code())
				return
			}
			game := test.game
			assert.Equal(int32(game.playerCount), resp.Msg.GetPlayerCount())
			assert.Equal(int32(game.currentTurn), resp.Msg.GetCurrentTurn())
			assert.Equal(int32(game.winner), resp.Msg.GetWinner())
			assert.Equal(int32(game.boardSize), resp.Msg.GetData().Dimension)
			assert.Equal(test.expectedPlayers, resp.Msg.GetPlayers())

			data := resp.Msg.GetData()
			for i, row := range game.board {
				for j, val := range row {
					assert.Equal(int32(val), data.Cells[i].Cells[j])
				}
			}
		})
	}
}

func TestServerMakeMove(t *testing.T) {
	assert := assert.New(t)
	tttSrv, client := makeMockServerAndClient()

	testBoard := func() board {
		return board([][]int{
			{0, 1, 0, 0},
			{0, 2, 0, 0},
			{0, 0, 1, 0},
			{0, 2, 0, 0},
		})
	}

	tests := []struct {
		name            string
		board           board
		expectedBoard   board
		winner          int
		currentTurn     int
		move            *tttproto.Move
		gameId          string
		expectedErrCode connect.Code
	}{
		{
			name:  "TestMove1",
			board: testBoard(),
			expectedBoard: board([][]int{
				{0, 1, 0, 0},
				{0, 2, 0, 0},
				{0, 1, 1, 0},
				{0, 2, 0, 0},
			}),
			winner:      EMPTY_PLAYER,
			currentTurn: 1,
			gameId:      "1",
			move:        &tttproto.Move{Row: 2, Col: 1},
		}, {
			name:  "TestMove2",
			board: testBoard(),
			expectedBoard: board([][]int{
				{0, 1, 0, 1},
				{0, 2, 0, 0},
				{0, 0, 1, 0},
				{0, 2, 0, 0},
			}),
			winner:      EMPTY_PLAYER,
			currentTurn: 1,
			gameId:      "1",
			move:        &tttproto.Move{Row: 0, Col: 3},
		}, {
			name:            "OutOfBounds",
			board:           testBoard(),
			winner:          EMPTY_PLAYER,
			currentTurn:     1,
			move:            &tttproto.Move{Row: -1, Col: 3},
			gameId:          "1",
			expectedErrCode: connect.CodeFailedPrecondition,
		}, {
			name:            "IncorrectPlayer",
			board:           testBoard(),
			winner:          EMPTY_PLAYER,
			currentTurn:     2,
			move:            &tttproto.Move{Row: 0, Col: 3},
			gameId:          "1",
			expectedErrCode: connect.CodeFailedPrecondition,
		}, {
			name:            "GameAlreadyOver",
			board:           testBoard(),
			winner:          1,
			currentTurn:     1,
			move:            &tttproto.Move{Row: 0, Col: 3},
			gameId:          "1",
			expectedErrCode: connect.CodeFailedPrecondition,
		}, {
			name:            "ExistingMove",
			board:           testBoard(),
			winner:          EMPTY_PLAYER,
			currentTurn:     1,
			move:            &tttproto.Move{Row: 0, Col: 1},
			gameId:          "1",
			expectedErrCode: connect.CodeFailedPrecondition,
		}, {
			name:            "NotFound",
			board:           testBoard(),
			winner:          EMPTY_PLAYER,
			currentTurn:     1,
			move:            &tttproto.Move{Row: 0, Col: 1},
			gameId:          "2",
			expectedErrCode: connect.CodeNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tttSrv.games["1"] = &game{
				maxPlayers:    2,
				playerCount:   2,
				board:         test.board,
				boardSize:     4,
				connectTarget: 3,
				currentTurn:   test.currentTurn,
				winner:        test.winner,
			}

			_, err := client.MakeMove(context.Background(),
				&connect_go.Request[tttproto.MakeMoveRequest]{Msg: &tttproto.MakeMoveRequest{
					Move:   test.move,
					GameId: test.gameId,
					Player: 1,
				}},
			)
			if err != nil {
				if test.expectedErrCode == 0 {
					t.Fatal(err)
				}
				connErr := &connect.Error{}
				require.ErrorAs(t, err, &connErr)
				assert.Equal(test.expectedErrCode, connErr.Code())
				return
			}

			assert.Equal(test.expectedBoard, tttSrv.games["1"].board)
		})
	}
}
