package api

import (
	"fmt"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

// GenerateByPhont converts text, generates a wav via AquesTalk2, and registers the task.
func GenerateByPhont(params *GenerateByPhontParams) (*GenerateByPhontResult, error) {
	task, err := PrepareGenerateByPhont(params)
	if err != nil {
		return nil, err
	}
	phontPath, err := PhontPath(*task.PhontName)
	if err != nil {
		return nil, err
	}
	if err := task.Generate(phontPath, utils.ResultDir); err != nil {
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

// PrepareGenerateByPhont validates inputs and creates a task without generating audio.
func PrepareGenerateByPhont(params *GenerateByPhontParams) (*singlesentence.Task, error) {
	if singlesentence.Manager.HasTask(params.TaskName) {
		return nil, fmt.Errorf("task %s already exists", params.TaskName)
	}
	_, exists := phontsmanager.PhontNameToFileName.GetValue(params.PhontName)
	if !exists {
		return nil, fmt.Errorf("no such phont %s", params.PhontName)
	}
	processedText, err := language.ConvertText(
		params.Text,
		language.ToLanguage(params.Language),
	)
	if err != nil {
		return nil, err
	}
	phontName := params.PhontName
	return singlesentence.NewSingleSentenceTask(
		processedText,
		nil,
		&phontName,
		params.Speed,
		params.TaskName,
	)
}
