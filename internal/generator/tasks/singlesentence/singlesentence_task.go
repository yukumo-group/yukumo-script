package singlesentence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/aquestalk2"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

// Task defines the task of generating a single sentence
type Task struct {
	sync.RWMutex
	ID                string            `json:"id"`
	TaskName          string            `json:"taskName"`
	Text              string            `json:"text"`
	TaskLanguage      language.Language `json:"taskLanguage"`
	Speed             int               `json:"speed"`
	CreateTime        time.Time         `json:"createTime"`
	EditTime          time.Time         `json:"editTime"`
	CharacterID       *string           `json:"characterID"`
	PhontName         *string           `json:"phontName"`
	ResultFile        *string           `json:"resultFile"`
	charactersManager *characters.Characters
}

// NewSingleSentenceTask creates new single sentence task
func NewSingleSentenceTask(
	text string,
	characterID *string,
	phontName *string,
	speed int,
	taskName string,
	taskLanguage language.Language,
	charactersManager *characters.Characters,
) (*Task, error) {
	if phontName == nil && characterID == nil {
		return nil, errors.New(
			"you have to choose at least one way to generate the audio",
		)
	}
	id := uuid.NewString()
	return &Task{
		ID:                id,
		Text:              text,
		CreateTime:        time.Now(),
		EditTime:          time.Now(),
		CharacterID:       characterID,
		PhontName:         phontName,
		Speed:             speed,
		ResultFile:        nil,
		TaskName:          taskName,
		TaskLanguage:      taskLanguage,
		charactersManager: charactersManager,
	}, nil
}

// NewSingleSentenceTaskFromFile gets single sentence task from file
func NewSingleSentenceTaskFromFile(
	fileName string,
) (*Task, error) {
	var result Task
	data, errRead := os.ReadFile(fileName)
	if errRead != nil {
		return nil, errRead
	}
	errJSON := json.Unmarshal(data, &result)
	if errJSON != nil {
		return nil, errJSON
	}
	return &result, nil
}

// GenerateFileName generates filename of metadata for this task
func (task *Task) GenerateFileName(
	targetDir string,
) string {
	return fmt.Sprintf(
		"%s/%s_%s_%d.json",
		targetDir,
		task.TaskName,
		task.ID,
		task.CreateTime.Unix(),
	)
}

// GenerateWavName generates name for the result wav file
func (task *Task) GenerateWavFileName(
	targetDir string,
) string {
	return fmt.Sprintf(
		"%s/%s_%s_%d.wav",
		targetDir,
		task.TaskName,
		task.ID,
		task.EditTime.Unix(),
	)
}

// SaveFile saves the file in the target directory.
// Returns file path and error.
func (task *Task) SaveFile(
	targetDir string,
) (string, error) {
	task.Lock()
	defer task.Unlock()
	// Marshal
	marshalResult, errMarshal := json.Marshal(task)
	if errMarshal != nil {
		return "", errMarshal
	}
	// Write file
	fileName := task.GenerateFileName(targetDir)
	task.EditTime = time.Now()
	errWrite := os.WriteFile(
		fileName,
		marshalResult,
		0644,
	)
	if errWrite != nil {
		return "", errWrite
	}
	return fileName, nil
}

// Generate synthesizes the wav file via AquesTalk2.
func (task *Task) Generate(
	ctx context.Context,
	phontsDir string,
	targetDir string,
) error {
	task.Lock()
	task.EditTime = time.Now()
	fileName := task.GenerateWavFileName(targetDir)
	taskText := task.Text
	taskLanguage := task.TaskLanguage
	taskCharacterID := task.CharacterID
	characterManager := task.charactersManager
	phontName := task.PhontName
	speed := task.Speed
	task.Unlock()
	processedText, err := language.ConvertText(
		taskText,
		taskLanguage,
	)
	if err != nil {
		return err
	}
	var phontPath string
	if taskCharacterID != nil {
		if characterManager == nil {
			return errors.New(
				"the character list is nil, which is not allowed",
			)
		}
		characterList := characterManager.GetData()
		character, exists := characterList[*taskCharacterID]
		if !exists {
			return fmt.Errorf(
				"character %s does not exists",
				*taskCharacterID,
			)
		}
		phontPath, err = tasks.PhontFile(
			phontsDir,
			character.PhontName,
		)
		if err != nil {
			return err
		}
	} else if phontName != nil {
		phontPath, err = tasks.PhontFile(
			phontsDir,
			*phontName,
		)
		if err != nil {
			return err
		}
	} else {
		return errors.New(
			"phont name and character id cannot both not exists",
		)
	}
	generator := aquestalk2.NewGenerator(
		speed,
		phontPath,
		fileName,
		processedText,
	)
	err = generator.GenerateWav()
	if err != nil {
		return err
	}
	task.Lock()
	task.ResultFile = &fileName
	task.Unlock()
	return nil
}

// GetTaskName gets the name of the task
func (task *Task) GetTaskName() string {
	task.RLock()
	defer task.RUnlock()
	return task.TaskName
}

// GetResultFile gets result file for certain task
func (task *Task) GetResultFile() *string {
	task.RLock()
	defer task.RUnlock()
	return task.ResultFile
}
