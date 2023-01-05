package tictactoe

import (
	"testing"

	"github.com/kertox662/game-site/proto/games/tictactoe"
	"github.com/stretchr/testify/assert"
)

func TestGameToProto(t *testing.T) {
	testBoard := board([][]int{
		{1, 2, 1, 1},
		{2, 3, 3, 2},
		{3, 2, 1, 3},
		{2, 3, 1, 1},
	})

	tests := []struct {
		name     string
		game     *game
		expected *tictactoe.GetGameDataResponse
	}{
		{
			name: "Empty",
			game: &game{
				board:         newBoard(3),
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
			expected: &tictactoe.GetGameDataResponse{
				Data: &tictactoe.Board{
					Dimension: 3,
					Cells: []*tictactoe.BoardRow{
						{
							Length: 3,
							Cells:  []int32{0, 0, 0},
						}, {
							Length: 3,
							Cells:  []int32{0, 0, 0},
						}, {
							Length: 3,
							Cells:  []int32{0, 0, 0},
						},
					},
				},
				PlayerCount: 2,
				Players:     []string{"1", "2"},
				CurrentTurn: 1,
				Winner:      0,
			},
		},
		{
			name: "TestBoard",
			game: &game{
				board:         testBoard,
				boardSize:     4,
				maxPlayers:    5,
				playerCount:   3,
				connectTarget: 3,
				currentTurn:   2,
				winner:        EMPTY_PLAYER,
			},
			expected: &tictactoe.GetGameDataResponse{
				Data: &tictactoe.Board{
					Dimension: 4,
					Cells: []*tictactoe.BoardRow{
						{
							Length: 4,
							Cells:  []int32{1, 2, 1, 1},
						}, {
							Length: 4,
							Cells:  []int32{2, 3, 3, 2},
						}, {
							Length: 4,
							Cells:  []int32{3, 2, 1, 3},
						}, {
							Length: 4,
							Cells:  []int32{2, 3, 1, 1},
						},
					},
				},
				PlayerCount: 3,
				Players:     []string{"1", "2", "3"},
				CurrentTurn: 2,
				Winner:      0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.game.toProto())
		})
	}
}

func TestGameMakeMove(t *testing.T) {
	testBoard := newBoard(5)
	testBoard.makeMove(2, 2, 1)

	tests := []struct {
		name   string
		game   *game
		player int
		row    int
		col    int
		err    error
	}{
		{
			name: "No error",
			game: &game{
				board:         newBoard(5),
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
			player: 1,
			row:    2,
			col:    2,
			err:    nil,
		},
		{
			name: "Game is over",
			game: &game{
				board:         newBoard(3),
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   1,
				winner:        2,
			},
			player: 1,
			row:    3,
			col:    3,
			err:    ErrGameConcluded,
		},
		{
			name: "Incorrect Turn",
			game: &game{
				board:         newBoard(3),
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   2,
				winner:        EMPTY_PLAYER,
			},
			player: 1,
			row:    3,
			col:    3,
			err:    ErrIncorrectPlayerTurn,
		},
		{
			name: "Non-empty space",
			game: &game{
				board:         testBoard,
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
			player: 1,
			row:    2,
			col:    2,
			err:    ErrNonEmptySpace,
		},
		{
			name: "Out of bounds move 1",
			game: &game{
				board:         newBoard(3),
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
			player: 1,
			row:    3,
			col:    3,
			err:    ErrOutOfBounds,
		},
		{
			name: "Out of bounds move 2",
			game: &game{
				board:         newBoard(3),
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
			player: 1,
			row:    -1,
			col:    -2,
			err:    ErrOutOfBounds,
		},
		{
			name: "Out of bounds move 3",
			game: &game{
				board:         newBoard(3),
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
			player: 1,
			row:    -1,
			col:    2,
			err:    ErrOutOfBounds,
		},
		{
			name: "Out of bounds move 4",
			game: &game{
				board:         newBoard(3),
				boardSize:     3,
				maxPlayers:    2,
				playerCount:   2,
				connectTarget: 3,
				currentTurn:   1,
				winner:        EMPTY_PLAYER,
			},
			player: 1,
			row:    1,
			col:    -1,
			err:    ErrOutOfBounds,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			err := test.game.makeMove(test.player, test.row, test.col)
			assert.ErrorIs(test.err, err)
			if err == nil {
				assert.Equal(testBoard, test.game.board)
			}
		})
	}

}
