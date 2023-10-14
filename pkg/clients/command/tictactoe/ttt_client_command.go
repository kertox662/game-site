package tictactoe

import (
	"strconv"

	"github.com/kertox662/game-site/proto/games/tictactoe/tictactoeconnect"
)

type tttClientCmd struct {
	protoClient tictactoeconnect.TicTacToeServiceClient
}

func (c *tttClientCmd) WithProtoClient(client tictactoeconnect.TicTacToeServiceClient) {
	c.protoClient = client
}

func parseInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int32(i), err
}
