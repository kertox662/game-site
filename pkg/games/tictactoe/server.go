package tictactoe

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	"github.com/kertox662/game-site/proto/games/tictactoe"
	"github.com/kertox662/game-site/proto/games/tictactoe/tictactoeconnect"
	"go.melnyk.org/hufid"
)

var _ tictactoeconnect.TicTacToeServiceHandler = (*Server)(nil)

type Server struct {
	tictactoeconnect.UnimplementedTicTacToeServiceHandler
	games map[string]*game
}

// Function to create a new server
func NewServer() *Server {
	return &Server{
		games: make(map[string]*game),
	}
}

// CreateGame implements connect interface
// It returns a game id for the game or an error message
func (s *Server) CreateGame(ctx context.Context,
	req *connect.Request[tictactoe.CreateGameRequest]) (*connect.Response[tictactoe.CreateGameResponse], error) {
	msg := req.Msg

	//If any game-building values are given as negative numbers, send out an error and no response
	if msg.BoardSize <= 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("non-positive board size"))
	} else if msg.ConnectTarget <= 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("non-positive connect target"))
	} else if msg.MaxPlayers <= 0 {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("non-positive max players"))
	}

	// If there are no errors, create a new game id
	id := string(hufid.NewUniqID(4))
	s.games[id] = newGame(
		int(msg.BoardSize),
		int(msg.MaxPlayers),
		int(msg.ConnectTarget),
	)

	// This returns the game id
	return &connect.Response[tictactoe.CreateGameResponse]{
		Msg: &tictactoe.CreateGameResponse{
			GameId: id,
		},
	}, nil
}

// GetGameData implements connect interface
// It checks to see if it has recieved a valid game id
func (s *Server) GetGameData(ctx context.Context,
	req *connect.Request[tictactoe.GetGameDataRequest]) (*connect.Response[tictactoe.GetGameDataResponse], error) {
	msg := req.Msg
	g, ok := s.games[msg.GameId]

	//If there is not a valid game id, then return an error message
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, nil)
	}

	return &connect.Response[tictactoe.GetGameDataResponse]{
		Msg: g.toProto(),
	}, nil
}

// MakeMove implements connect interface
// It checks if the given parameters describe a valid move. If it does, it makes that move.
// If it doesn't, it returns an error
func (s *Server) MakeMove(ctx context.Context,
	req *connect.Request[tictactoe.MakeMoveRequest]) (*connect.Response[tictactoe.MakeMoveResponse], error) {
	msg := req.Msg
	g, ok := s.games[msg.GameId]

	//If there is not a valid game id, then return an error message
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, nil)
	}

	// Try to make the move, if it doesn't work, return an error
	if err := g.makeMove(int(msg.Player), int(msg.Move.Row), int(msg.Move.Col)); err != nil {
		return nil, connect.NewError(connect.CodeFailedPrecondition, err)
	}

	return &connect.Response[tictactoe.MakeMoveResponse]{}, nil
}
