package tictactoe

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/kertox662/game-site/pkg/clients/command"
	"github.com/kertox662/game-site/proto/games/tictactoe"
)

type PlayMoveCommand struct {
	command.BaseCommand
	tttClientCmd
}

func (cgc *PlayMoveCommand) Execute(ctx context.Context, args []string) ([]string, error) {
	messages := make([]string, 0)

	if len(args) != 4 {
		return nil, fmt.Errorf("expected 4 arguments, got %d", len(args))
	}

	gameId := args[0]
	player, err := parseInt32(args[1])
	if err != nil {
		return messages, err
	}

	row, err := parseInt32(args[2])
	if err != nil {
		return messages, err
	}
	col, err := parseInt32(args[3])
	if err != nil {
		return messages, err
	}

	req := &connect.Request[tictactoe.MakeMoveRequest]{
		Msg: &tictactoe.MakeMoveRequest{
			GameId: gameId,
			Player: player,
			Move: &tictactoe.Move{
				Row: row,
				Col: col,
			},
		},
	}
	messages = append(messages, fmt.Sprintf("Making move for game %s", gameId))
	_, err = cgc.protoClient.MakeMove(ctx, req)
	if err != nil {
		return messages, err
	}

	messages = append(messages, "Move successful")
	return messages, nil
}

func (cgc *PlayMoveCommand) Help() string {
	return "args: <gameId> <player> <x> <y>"
}
