package cmd

import (
	"github.com/spf13/cobra"
)

var gameAPICmd = &cobra.Command{
	Use:   "game",
	Short: "Runs the games API",
}

func init() {
	apiCmd.AddCommand(gameAPICmd)
}
