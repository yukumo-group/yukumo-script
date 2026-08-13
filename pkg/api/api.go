package api

import (
	"context"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/example"
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
)

// Init initializes runtime dirs, examples, phont map, characters, and tasks.
func Init() error {
	InitializePathesByConst()
	InitRuntimeDirs()

	dir, err := phontsmanager.GetAllPhonts(filePathForProg.PhontsDir)
	if err != nil {
		return err
	}
	if err := example.GenerateExamples(
		context.Background(),
		filePathForProg.ExampleDir,
		filePathForProg.PhontsDir,
		dir,
	); err != nil {
		return err
	}
	if err := InitPhontMap(); err != nil {
		return err
	}

	characters.CharacterList.SetTargetFile(
		filePathForProg.DataDir,
		filePathForProg.CharactersFile,
	)
	if err := characters.CharacterList.ReadData(); err != nil {
		return err
	}
	return InitTaskManager()
}

// InitRuntimeDirs creates the runtime directories used by CLI and clib.
func InitRuntimeDirs() {
	utils.InitializeDirectory(filePathForProg.RuntimeDir)
	utils.InitializeDirectory(filePathForProg.PhontsDir)
	utils.InitializeDirectory(filePathForProg.ResultDir)
	utils.InitializeDirectory(filePathForProg.WavsDir)
	utils.InitializeDirectory(filePathForProg.DataDir)
	utils.InitializeDirectory(filePathForProg.ExampleDir)
	utils.InitializeDirectory(filePathForProg.ImagesDir)
	utils.InitializeDirectory(filePathForProg.TaskDir)
	utils.InitializeDirectory(filePathForProg.SingleSentenceDir)
}

// InitializePathesByConst initializes the pathes according to the constants in utils
func InitializePathesByConst() {
	filePathForProg.RuntimeDir = utils.RuntimeDir
	filePathForProg.PhontsDir = utils.PhontsDir
	filePathForProg.ResultDir = utils.ResultDir
	filePathForProg.WavsDir = utils.WavsDir
	filePathForProg.DataDir = utils.DataDir
	filePathForProg.ExampleDir = utils.ExampleDir
	filePathForProg.ImagesDir = utils.ImagesDir
	filePathForProg.TaskDir = utils.TaskDir
	filePathForProg.SingleSentenceDir = utils.SingleSentenceDir
	filePathForProg.SingleSentenceTasksFile = utils.SingleSentenceTasksFile
	filePathForProg.CharactersFile = utils.CharactersFile
	filePathForProg.ConfDir = utils.ConfDir
}
