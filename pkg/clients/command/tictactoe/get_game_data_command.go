package tictactoe

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/kertox662/game-site/pkg/clients/command"
	"github.com/kertox662/game-site/proto/games/tictactoe"
)

type GetGameDataCommand struct {
	command.BaseCommand
	tttClientCmd
}

func (cgc *GetGameDataCommand) Execute(ctx context.Context, args []string) ([]string, error) {
	messages := make([]string, 0)

	if len(args) != 1 {
		return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
	}

	gameId := args[0]
	req := &connect.Request[tictactoe.GetGameDataRequest]{Msg: &tictactoe.GetGameDataRequest{GameId: gameId}}
	messages = append(messages, fmt.Sprintf("Getting game data for game %s", gameId))
	resp, err := cgc.protoClient.GetGameData(ctx, req)
	if err != nil {
		return messages, err
	}

	board := dataToBoard(resp.Msg.GetData())
	messages = append(messages, boardToString(board))
	messages = append(messages, fmt.Sprintf(
		"Player count: %d\n Current player: %d\n",
		resp.Msg.Metadata.PlayerCount,
		resp.Msg.Metadata.CurrentTurn,
	))

	metadata, ok := command.GetStateData[gameMetadata](cgc.State(), gameDataKey)
	if !ok {
		metadata = gameMetadata{}
	}
	metadata.CurrentPlayer = resp.Msg.Metadata.CurrentTurn
	metadata.PlayerCount = resp.Msg.Metadata.PlayerCount
	command.SetStateData[gameMetadata](cgc.State(), gameDataKey, metadata)

	data := gameData{
		Board: board,
	}
	command.SetStateData[gameData](cgc.State(), gameId, data)

	return messages, nil
}

func (cgc *GetGameDataCommand) Help() string {
	return "args: <max players> <board size> <connect target>"
}

func dataToBoard(board *tictactoe.Board) [][]int32 {
	result := make([][]int32, 0)
	for _, row := range board.GetCells() {
		result = append(result, row.GetCells())
	}
	return result
}

func boardToString(board [][]int32) string {
	result := ""
	for _, row := range board {
		for _, cell := range row {
			result += fmt.Sprintf("%d ", cell)
		}
		result += "\n"
	}
	return result
}
