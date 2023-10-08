package tictactoe

import (
	"context"
	"fmt"

	"github.com/kertox662/game-site/pkg/clients/command"
)

type PrintGameCommand struct {
	command.BaseCommand
}

func (pgc *PrintGameCommand) Execute(ctx context.Context, args []string) ([]string, error) {
	messages := make([]string, 0)
	if len(args) != 1 {
		return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	gameId := args[0]
	game, ok := command.GetStateData[gameData](pgc.State(), gameId)
	if !ok {
		return nil, fmt.Errorf("games are not loaded")
	}

	cols := make([]string, len(game.Board)+1)
	cols[0] = ""
	for i := range game.Board {
		cols[i] = fmt.Sprintf("%d", i)
	}
	gameTable := command.NewTable(cols)
	for i := range game.Board {
		row := make([]string, len(game.Board)+1)
		row[0] = fmt.Sprintf("%d", i)
		for j := range game.Board {
			row[j+1] = fmt.Sprintf("%d", game.Board[i][j])
		}
		gameTable.AddRow(row)
	}

	messages = append(messages, gameTable.String())
	return messages, nil
}

func (pgc *PrintGameCommand) Help() string {
	return "args:"
}
