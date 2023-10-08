package cmd

import (
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Runs a game client",
}

func init() {
	clientCmd.AddCommand(tttClientCmd)
}
