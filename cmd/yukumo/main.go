//go:build windows
// +build windows

package main

import (
	"context"

	"github.com/yukumo-group/yukumo-script/pkg/characters"
	"github.com/yukumo-group/yukumo-script/pkg/example"
	"github.com/yukumo-group/yukumo-script/pkg/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/pkg/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
	"github.com/yukumo-group/yukumo-script/pkg/utils/logger"
)

var cliLogger = logger.NewLogger(
	"CLI",
	nil,
)

// Initialize directories for storing data
func init() {
	// Initialize directories
	utils.InitializeDirectory(utils.PhontsDir)
	utils.InitializeDirectory(utils.ResultDir)
	utils.InitializeDirectory(utils.WavsDir)
	utils.InitializeDirectory(utils.DatasDir)
	utils.InitializeDirectory(utils.ExampleDir)
	utils.InitializeDirectory(utils.ImagesDir)
	utils.InitializeDirectory(utils.TaskDir)
	utils.InitializeDirectory(utils.SingleSentenceDir)
	dir, err := phontsmanager.GetAllPhonts(utils.PhontsDir)
	if err != nil {
		cliLogger.Error(err.Error())
		panic(err)
	}
	ctx := context.Background()
	// Generate examples
	err = example.GenerateExampleWin(
		ctx,
		utils.ExampleDir,
		utils.PhontsDir,
		dir,
	)
	if err != nil {
		cliLogger.Error(err.Error())
		panic(err)
	}
	// Initialize phont name 2 file name
	err = phontsmanager.InitializePhontNameToFileName(
		utils.PhontsDir,
	)
	if err != nil {
		cliLogger.Error(err.Error())
		panic(err)
	}
	// Config the character list
	characters.CharacterList.SetTargetFile(
		utils.DatasDir,
		utils.CharactersFile,
	)
	err = characters.CharacterList.ReadData()
	if err != nil {
		cliLogger.Error(err.Error())
		panic(err)
	}
	// Config the task list
	singlesentence.Manager.SetTargetFile(
		utils.TaskDir,
		utils.SingleSentenceTasksFile,
	)
	err = singlesentence.Manager.ReadData()
	if err != nil {
		cliLogger.Error(err.Error())
		panic(err)
	}
}

// Main process
func main() {
	Execute()
}
