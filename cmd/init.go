package cmd

import (
	"github.com/spf13/cobra"
)

// InitFilesCmd initialises a fresh Tendermint Core instance.
var InitFilesCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize config",
	Run:   initFiles,
}

func initFiles(cmd *cobra.Command, args []string) {
}
