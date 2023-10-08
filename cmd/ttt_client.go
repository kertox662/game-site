package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kertox662/game-site/pkg/clients/console"
	"github.com/kertox662/game-site/proto/games/tictactoe/tictactoeconnect"
	"github.com/spf13/cobra"
)

var tttClientCmd = &cobra.Command{
	Use:   "ttt",
	Short: "Runs the tic-tac-toe client",
	Run: func(cmd *cobra.Command, args []string) {
		client := tictactoeconnect.NewTicTacToeServiceClient(
			http.DefaultClient,
			"http://localhost:8080",
		)
		fmt.Println("Running Tic-Tac-Toe client...")
		c := console.NewConsole(client, os.Stdin, os.Stdout)
		c.Run(cmd.Context())
	},
}
