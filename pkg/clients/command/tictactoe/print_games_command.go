package tictactoe

import (
	"context"
	"fmt"

	"github.com/kertox662/game-site/pkg/clients/command"
)

type PrintGamesCommand struct {
	command.BaseCommand
}

func (pgc *PrintGamesCommand) Execute(ctx context.Context, args []string) ([]string, error) {
	games, ok := command.GetStateData[map[string]gameMetadata](pgc.State(), gameDataKey)
	if !ok {
		return []string{}, nil
	}

	messages := make([]string, 0)
	gameTable := command.NewTable([]string{
		"Game Id", "Current Player", "Player Count", "Max Players", "Board Size", "Connect Target",
	})
	for id, metadata := range games {
		err := gameTable.AddRow([]string{
			id,
			fmt.Sprintf("%d", metadata.CurrentPlayer),
			fmt.Sprintf("%d", metadata.PlayerCount),
			fmt.Sprintf("%d", metadata.MaxPlayers),
			fmt.Sprintf("%d", metadata.BoardSize),
			fmt.Sprintf("%d", metadata.ConnectTarget),
		})
		if err != nil {
			return nil, err
		}
	}
	messages = append(messages, gameTable.String())

	return messages, nil
}

func (pgc *PrintGamesCommand) Help() string {
	return "args:"
}
