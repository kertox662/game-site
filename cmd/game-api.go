package cmd

import (
	"log"

	"github.com/kertox662/game-site/pkg/games/server"
	"github.com/spf13/cobra"
)

const defaultAddr = ":8080"

var gameAPICmd = &cobra.Command{
	Use:   "game",
	Short: "Runs the games API",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.NewServer(defaultAddr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	},
}
