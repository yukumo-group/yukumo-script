package api

import (
	"errors"
	"fmt"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
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

// GenerateByPhont converts text, generates a wav via AquesTalk2, and registers the task.
func GenerateByPhont(
	params *GenerateByPhontParams,
) (*GenerateResult, error) {
	task, err := PrepareGenerateByPhont(params)
	if err != nil {
		return nil, err
	}
	if task.PhontName == nil {
		return nil, errors.New(
			"Phont name cannot be nil when generating",
		)
	}
	phontPath, err := PhontPath(*task.PhontName)
	if err != nil {
		return nil, err
	}
	if err := task.Generate(phontPath, filePathForProg.ResultDir); err != nil {
		return nil, err
	}
	if task.ResultFile == nil {
		return nil, errors.New("result file path is nil after generation")
	}
	taskFile, err := RegisterGeneratedTask(task)
	if err != nil {
		return nil, err
	}
	return &GenerateResult{
		ResultFile: *task.ResultFile,
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
	processedText, err := language.ConvertText(
		params.Text,
		language.ToLanguage(params.Language),
	)
	if err != nil {
		return nil, err
	}
	characterName := params.CharacterName
	return singlesentence.NewSingleSentenceTask(
		processedText,
		&characterName,
		nil,
		params.Speed,
		params.TaskName,
	)
}

// GenerateByCharacter converts text, generates a wav via AquesTalk2, and registers the task.
func GenerateByCharacter(
	params *GenerateByCharacterParams,
) (*GenerateResult, error) {
	task, err := PrepareGenerateByCharacter(params)
	if err != nil {
		return nil, err
	}
	if task.CharacterID == nil {
		return nil, errors.New(
			"The character ID (character name) cannot be nil",
		)
	}
	phontName, errGetPhontName := GetPhontNameByCharacterName(
		*task.CharacterID,
	)
	if errGetPhontName != nil {
		return nil, errGetPhontName
	}
	phontPath, errGetPhontPath := PhontPath(phontName)
	if errGetPhontPath != nil {
		return nil, errGetPhontPath
	}
	if err := task.Generate(phontPath, filePathForProg.ResultDir); err != nil {
		return nil, err
	}
	if task.ResultFile == nil {
		return nil, fmt.Errorf("result file path is nil after generation")
	}
	taskFile, err := RegisterGeneratedTask(task)
	if err != nil {
		return nil, err
	}
	return &GenerateResult{
		ResultFile: *task.ResultFile,
		TaskFile:   taskFile,
	}, nil
}
