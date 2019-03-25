package cmd

import (
	"fmt"
	"github.com/kooksee/dbproxy/version"
	"github.com/spf13/cobra"
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version")
		fmt.Println("app version", version.Version)
		fmt.Println("app commit version", version.CommitVersion)
		fmt.Println("app build version", version.BuildVersion)
	},
}
