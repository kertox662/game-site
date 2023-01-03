package server

import (
	"net/http"

	"github.com/kertox662/game-site/pkg/games/tictactoe"
	"github.com/kertox662/game-site/proto/games/tictactoe/tictactoeconnect"
)

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()

	tttSrv := tictactoe.NewServer()
	tttPath, tttHandler := tictactoeconnect.NewTicTacToeServiceHandler(tttSrv)
	mux.Handle(tttPath, tttHandler)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
