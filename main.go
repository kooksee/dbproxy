package main

import (
	"github.com/kooksee/dbproxy/cmd"
	"github.com/kooksee/dbproxy/internal/cnst"
	"github.com/kooksee/go-assert"
	"github.com/kooksee/kweb"
	"os"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.InitFilesCmd,
		cmd.VersionCmd,
		cmd.ServerCmd,
	)

	assert.MustNotError(kweb.PrepareBaseCmd(rootCmd, cnst.EnvPrefix,
		os.ExpandEnv(cnst.CurPath)).Execute())
}
