package cmd

import (
	"fmt"
	"github.com/kooksee/dbproxy/internal/app"
	"github.com/kooksee/dbproxy/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.PersistentFlags().String("log_level", "debug", "Log level")
}

var RootCmd = &cobra.Command{
	Use:     "app",
	Short:   "app service",
	Version: version.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("root")
		return viper.Unmarshal(app.DefaultApp().Cfg)
	},
}
