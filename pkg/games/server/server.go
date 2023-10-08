package server

import (
	"net/http"
	"time"

	"github.com/kertox662/game-site/pkg/games/tictactoe"
	"github.com/kertox662/game-site/proto/games/tictactoe/tictactoeconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()

	tttSrv := tictactoe.NewServer()
	tttPath, tttHandler := tictactoeconnect.NewTicTacToeServiceHandler(tttSrv)
	mux.Handle(tttPath, tttHandler)

	return &http.Server{
		Addr:              addr,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 5 * time.Second,
	}
}
