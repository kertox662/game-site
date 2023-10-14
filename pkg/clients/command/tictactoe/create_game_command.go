package tictactoe

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/kertox662/game-site/pkg/clients/command"
	"github.com/kertox662/game-site/proto/games/tictactoe"
)

type CreateGameCommand struct {
	command.BaseCommand
	tttClientCmd
}

func (cgc *CreateGameCommand) Execute(ctx context.Context, args []string) ([]string, error) {
	messages := make([]string, 0)

	if len(args) != 3 {
		return nil, fmt.Errorf("expected 3 arguments, got %d", len(args))
	}

	maxPlayers, err := parseInt32(args[2])
	if err != nil {
		return nil, fmt.Errorf("invalid max players: %w", err)
	}

	boardSize, err := parseInt32(args[2])
	if err != nil {
		return nil, fmt.Errorf("invalid board size: %w", err)
	}

	connectTarget, err := parseInt32(args[2])
	if err != nil {
		return nil, fmt.Errorf("invalid connect target: %w", err)
	}

	req := &connect.Request[tictactoe.CreateGameRequest]{
		Msg: &tictactoe.CreateGameRequest{
			MaxPlayers:    maxPlayers,
			BoardSize:     boardSize,
			ConnectTarget: connectTarget,
		},
	}

	messages = append(messages,
		fmt.Sprintf("Creating game with max players %d, board size %d, and connect target %d",
			maxPlayers, boardSize, connectTarget),
	)

	resp, err := cgc.protoClient.CreateGame(ctx, req)
	if err != nil {
		return messages, err
	}

	messages = append(messages, fmt.Sprintf("Game created with id %s", resp.Msg.GetGameId()))

	metadata := gameMetadata{
		CurrentPlayer: 0,
		PlayerCount:   0,
		MaxPlayers:    maxPlayers,
		BoardSize:     boardSize,
		ConnectTarget: connectTarget,
	}

	existing, ok := command.GetStateData[map[string]gameMetadata](cgc.State(), knownGamesKey)
	if !ok {
		existing = make(map[string]gameMetadata)
	}
	existing[resp.Msg.GetGameId()] = metadata
	// setData(cgc.State(), knownGamesKey, existing)

	return messages, nil
}

func (cgc *CreateGameCommand) Help() string {
	return "args: <max players> <board size> <connect target>"
}
