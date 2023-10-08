package tictactoe

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/kertox662/game-site/pkg/clients/command"
	"github.com/kertox662/game-site/proto/games/tictactoe"
)

type ListGamesCommand struct {
	command.BaseCommand
	tttClientCmd
}

func (cgc *ListGamesCommand) Execute(ctx context.Context, args []string) ([]string, error) {
	messages := make([]string, 0)

	if len(args) != 0 {
		return nil, fmt.Errorf("expected 0 arguments, got %d", len(args))
	}

	req := &connect.Request[tictactoe.ListGamesRequest]{Msg: &tictactoe.ListGamesRequest{}}
	resp, err := cgc.protoClient.ListGames(ctx, req)
	if err != nil {
		return nil, err
	}

	messages = append(messages, fmt.Sprintf("Retrieved %d games", len(resp.Msg.GetGames())))
	currentData, ok := command.GetStateData[map[string]gameMetadata](cgc.State(), gameDataKey)
	if !ok {
		currentData = make(map[string]gameMetadata)
	}
	for _, game := range resp.Msg.GetGames() {
		currentData[game.GetId()] = gameMetadata{
			CurrentPlayer: game.CurrentTurn,
			PlayerCount:   game.PlayerCount,
			MaxPlayers:    game.MaxPlayers,
			BoardSize:     game.BoardSize,
			ConnectTarget: game.ConnectToWin,
		}
	}
	command.SetStateData[map[string]gameMetadata](cgc.State(), gameDataKey, currentData)
	return messages, nil
}

func (cgc *ListGamesCommand) Help() string {
	return "args:"
}
