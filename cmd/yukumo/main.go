package main

import (
	"github.com/yukumo-group/yukumo-script/pkg/api"
	"github.com/yukumo-group/yukumo-script/pkg/utils/logger"
)

var cliLogger = logger.NewLogger(
	"CLI",
	nil,
)

// Initialize directories and shared runtime state
func init() {
	if err := api.Init(); err != nil {
		cliLogger.Error(err.Error())
		panic(err)
	}
}

// Main process
func main() {
	Execute()
}
