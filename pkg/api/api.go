package api

import (
	"context"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/example"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
)

// Init initializes runtime dirs, examples, phont map, characters, and tasks.
func Init() error {
	InitRuntimeDirs()

	dir, err := phontsmanager.GetAllPhonts(utils.PhontsDir)
	if err != nil {
		return err
	}
	if err := example.GenerateExamples(
		context.Background(),
		utils.ExampleDir,
		utils.PhontsDir,
		dir,
	); err != nil {
		return err
	}
	if err := InitPhontMap(); err != nil {
		return err
	}

	characters.CharacterList.SetTargetFile(
		utils.DataDir,
		utils.CharactersFile,
	)
	if err := characters.CharacterList.ReadData(); err != nil {
		return err
	}
	return InitTaskManager()
}

// InitRuntimeDirs creates the runtime directories used by CLI and clib.
func InitRuntimeDirs() {
	utils.InitializeDirectory(utils.RuntimeDir)
	utils.InitializeDirectory(utils.PhontsDir)
	utils.InitializeDirectory(utils.ResultDir)
	utils.InitializeDirectory(utils.WavsDir)
	utils.InitializeDirectory(utils.DataDir)
	utils.InitializeDirectory(utils.ExampleDir)
	utils.InitializeDirectory(utils.ImagesDir)
	utils.InitializeDirectory(utils.TaskDir)
	utils.InitializeDirectory(utils.SingleSentenceDir)
}

// InitPhontMap loads phont name → file mappings from PhontsDir.
func InitPhontMap() error {
	return phontsmanager.InitializePhontNameToFileName(utils.PhontsDir)
}

// InitTaskManager configures and loads the single-sentence task manager.
func InitTaskManager() error {
	singlesentence.Manager.SetTargetFile(
		utils.TaskDir,
		utils.SingleSentenceTasksFile,
	)
	return singlesentence.Manager.ReadData()
}
