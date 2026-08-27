package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/empty"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

// PrepareGenerateByPhont validates inputs and creates a task without generating audio.
func PrepareGenerateByPhont(params *GenerateByPhontParams) (*singlesentence.Task, error) {
	if singlesentence.Manager.HasTask(params.TaskName) {
		return nil, fmt.Errorf("task %s already exists", params.TaskName)
	}
	exists := IsPhontExists(params.PhontName)
	if !exists {
		return nil, fmt.Errorf("no such phont %s", params.PhontName)
	}
	phontName := params.PhontName
	return singlesentence.NewSingleSentenceTask(
		params.Text,
		nil,
		&phontName,
		params.Speed,
		params.TaskName,
		language.ToLanguage(params.Language),
		characters.CharacterList,
	)
}

// GenerateByPhont converts text, generates a wav via AquesTalk2, and registers the task.
func GenerateByPhont(
	ctx context.Context,
	params *GenerateByPhontParams,
) (*GenerateResult, error) {
	task, err := PrepareGenerateByPhont(params)
	if err != nil {
		return nil, err
	}
	err = task.Generate(
		ctx,
		filePathForProg.PhontsDir,
		filePathForProg.ResultDir,
	)
	if err != nil {
		return nil, err
	}
	resultFile := task.GetResultFile()
	if resultFile == nil {
		return nil, errors.New("result file path is nil after generation")
	}
	taskFile, err := RegisterGeneratedTask(task)
	if err != nil {
		return nil, err
	}
	return &GenerateResult{
		ResultFile: *resultFile,
		TaskFile:   taskFile,
	}, nil
}

// PrepareGenerateByCharacter validates inputs and creates a task without generating audio.
func PrepareGenerateByCharacter(
	params *GenerateByCharacterParams,
) (*singlesentence.Task, error) {
	if singlesentence.Manager.HasTask(params.TaskName) {
		return nil, fmt.Errorf("task %s already exists", params.TaskName)
	}
	exists := IsCharacterExists(params.CharacterName)
	if !exists {
		return nil, fmt.Errorf("no such character %s", params.CharacterName)
	}
	characterName := params.CharacterName
	return singlesentence.NewSingleSentenceTask(
		params.Text,
		&characterName,
		nil,
		params.Speed,
		params.TaskName,
		language.ToLanguage(params.Language),
		characters.CharacterList,
	)
}

// GenerateByCharacter converts text, generates a wav via AquesTalk2, and registers the task.
func GenerateByCharacter(
	ctx context.Context,
	params *GenerateByCharacterParams,
) (*GenerateResult, error) {
	task, err := PrepareGenerateByCharacter(params)
	if err != nil {
		return nil, err
	}
	err = task.Generate(
		ctx,
		filePathForProg.PhontsDir,
		filePathForProg.ResultDir,
	)
	if err != nil {
		return nil, err
	}
	resultFile := task.GetResultFile()
	if resultFile == nil {
		return nil, fmt.Errorf("result file path is nil after generation")
	}
	taskFile, err := RegisterGeneratedTask(task)
	if err != nil {
		return nil, err
	}
	return &GenerateResult{
		ResultFile: *resultFile,
		TaskFile:   taskFile,
	}, nil
}

// GenerateEmpty generates empty wav file.
// The generated audio will have the same format with the original audio.
func GenerateEmpty(
	ctx context.Context,
	length float64,
	originalAudioPath string,
) (*string, error) {
	format, errGetAudioInfo := audio.GetAudioInfo(
		originalAudioPath,
	)
	if errGetAudioInfo != nil {
		return nil, errGetAudioInfo
	}
	newEmptyTask, errNewEmptyTask := empty.NewEmptyTask(
		length,
		format,
	)
	if errNewEmptyTask != nil {
		return nil, errNewEmptyTask
	}
	errGenerate := newEmptyTask.Generate(
		ctx,
		"",
		filePathForProg.WavsDir,
	)
	if errGenerate != nil {
		return nil, errGenerate
	}
	if newEmptyTask.ResultFile == nil {
		return nil, errors.New(
			"the result file is possibly not generated",
		)
	}
	return newEmptyTask.ResultFile, nil
}
