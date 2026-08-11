//go:build windows

package api

import (
	"context"
	"fmt"

	"github.com/yukumo-group/yukumo-script/pkg/characters"
	"github.com/yukumo-group/yukumo-script/pkg/example"
	"github.com/yukumo-group/yukumo-script/pkg/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
)

// Init initializes runtime dirs, examples, phont map, characters, and tasks (Windows).
func Init() error {
	InitRuntimeDirs()

	dir, err := phontsmanager.GetAllPhonts(utils.PhontsDir)
	if err != nil {
		return err
	}
	if err := example.GenerateExampleWin(
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

// GenerateByPhont converts text, generates a wav via AquesTalk2, and registers the task.
func GenerateByPhont(params GenerateByPhontParams) (*GenerateByPhontResult, error) {
	task, err := PrepareGenerateByPhont(params)
	if err != nil {
		return nil, err
	}
	phontPath, err := PhontPath(*task.PhontName)
	if err != nil {
		return nil, err
	}
	if err := task.GenerateWin(phontPath, utils.ResultDir); err != nil {
		return nil, err
	}
	if task.ResultFile == nil {
		return nil, fmt.Errorf("result file path is nil after generation")
	}
	taskFile, err := RegisterGeneratedTask(task)
	if err != nil {
		return nil, err
	}
	return &GenerateByPhontResult{
		ResultFile: *task.ResultFile,
		TaskFile:   taskFile,
	}, nil
}
