package cmd

import (
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "node server",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
