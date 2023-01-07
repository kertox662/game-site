package cmd

import (
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Runs an API for the game site",
}

func init() {
	apiCmd.AddCommand(gameAPICmd)
}
