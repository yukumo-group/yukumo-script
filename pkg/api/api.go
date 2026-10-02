package api

import (
	"context"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/example"
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
)

func init() {
	InitializePathesByConst()
}

// Init initializes runtime dirs, examples, phont map, characters, and tasks.
func Init() error {
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

	InitSequenceTaskConfigManager()

	characters.CharacterList.SetTargetFile(
		filePathForProg.DataDir,
		filePathForProg.CharactersFile,
	)
	if err := characters.CharacterList.ReadData(); err != nil {
		return err
	}
	characters.CharacterList.CleanData()
	err = InitializeLanguageConfig()
	if err != nil {
		return err
	}
	err = InitSequenceTaskManager()
	if err != nil {
		return err
	}
	err = InitializeDefaultSequenceConfig()
	if err != nil {
		return err
	}
	err = InitTaskManager()
	return err
}

// InitRuntimeDirs creates the runtime directories used by CLI and clib.
func InitRuntimeDirs() {
	utils.InitializeDirectory(filePathForProg.RuntimeDir)
	utils.InitializeDirectory(filePathForProg.AssetsDir)
	utils.InitializeDirectory(filePathForProg.PhontsDir)
	utils.InitializeDirectory(filePathForProg.ResultDir)
	utils.InitializeDirectory(filePathForProg.WavsDir)
	utils.InitializeDirectory(filePathForProg.DataDir)
	utils.InitializeDirectory(filePathForProg.ExampleDir)
	utils.InitializeDirectory(filePathForProg.ImagesDir)
	utils.InitializeDirectory(filePathForProg.TaskDir)
	utils.InitializeDirectory(filePathForProg.SingleSentenceDir)
	utils.InitializeDirectory(filePathForProg.SequenceDir)
	utils.InitializeDirectory(filePathForProg.ConfigDir)
}

// InitializePathesByConst initializes the pathes according to the constants in utils
func InitializePathesByConst() {
	filePathForProg.SetPathes(
		WithRuntimeDir(
			utils.RuntimeDir,
		),
		WithPhontsDir(
			utils.PhontsDir,
		),
		WithResultDir(
			utils.ResultDir,
		),
		WithWavsDir(
			utils.WavsDir,
		),
		WithDataDir(
			utils.DataDir,
		),
		WithExampleDir(
			utils.ExampleDir,
		),
		WithImagesDir(
			utils.ImagesDir,
		),
		WithTaskDir(
			utils.TaskDir,
		),
		WithSingleSentenceDir(
			utils.SingleSentenceDir,
		),
		WithSingleSentenceTasksFile(
			utils.SingleSentenceTasksFile,
		),
		WithCharactersFile(
			utils.CharactersFile,
		),
		WithConfPath(
			utils.ConfPath,
		),
		WithSequenceDir(
			utils.SequenceDir,
		),
		WithConfigDir(
			utils.ConfigDir,
		),
		WithConfigManagerFile(
			utils.ConfigManagerFile,
		),
		WithPolyphonicsManagerFile(
			utils.PolyphonicsManagerFile,
		),
		WithSequenceTaskManagerFile(
			utils.SequenceTaskManagerFile,
		),
		WithAssetsDir(
			utils.AssetsDir,
		),
		WithDictFilePath(
			utils.DictZhSFile,
		),
	)
}

// InitializePathesByCostum allows the user to use their own file structure
func InitializePathesByCostum(
	opts ...Option,
) {
	filePathForProg.SetPathes(
		opts...,
	)
}
