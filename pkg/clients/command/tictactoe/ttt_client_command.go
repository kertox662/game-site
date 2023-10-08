package tictactoe

import "github.com/kertox662/game-site/proto/games/tictactoe/tictactoeconnect"

type tttClientCmd struct {
	protoClient tictactoeconnect.TicTacToeServiceClient
}

func (c *tttClientCmd) WithProtoClient(client tictactoeconnect.TicTacToeServiceClient) {
	c.protoClient = client
}
