package main

import (
	"yukumo-script-cmd/internal/welcome"

	"github.com/spf13/cobra"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

var cmdLogger = api.NewLogger("CMD")

// rootCMD defines the root command
var rootCMD = &cobra.Command{
	Use:   "yukumo",
	Short: "yukumo is a program that can generate yukumo audio",
	Long: `
Yukumo is a simple and flexible program that can generate yukumo audio without the need for network connection. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		welcome.CLIWelcome()
	},
}

// Execute executes the command
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		panic(err.Error())
	}
}
