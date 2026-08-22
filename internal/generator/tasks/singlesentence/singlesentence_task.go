package singlesentence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/aquestalk2"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

// Task defines the task of generating a single sentence
type Task struct {
	ID           string            `json:"id"`
	TaskName     string            `json:"taskName"`
	Text         string            `json:"text"`
	TaskLanguage language.Language `json:"taskLanguage"`
	Speed        int               `json:"speed"`
	CreateTime   time.Time         `json:"createTime"`
	EditTime     time.Time         `json:"editTime"`
	CharacterID  *string           `json:"characterID"`
	PhontName    *string           `json:"phontName"`
	ResultFile   *string           `json:"resultFile"`
}

// NewSingleSentenceTask creates new single sentence task
func NewSingleSentenceTask(
	text string,
	characterID *string,
	phontName *string,
	speed int,
	taskName string,
	taskLanguage language.Language,
) (*Task, error) {
	if phontName == nil && characterID == nil {
		return nil, errors.New(
			"you have to choose at least one way to generate the audio",
		)
	}
	id := uuid.NewString()
	return &Task{
		ID:           id,
		Text:         text,
		CreateTime:   time.Now(),
		EditTime:     time.Now(),
		CharacterID:  characterID,
		PhontName:    phontName,
		Speed:        speed,
		ResultFile:   nil,
		TaskName:     taskName,
		TaskLanguage: taskLanguage,
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

// SaveFile saves the file in the target directory.
func (task *Task) SaveFile(
	targetDir string,
) (string, error) {
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
	task.EditTime = time.Now()
	fileName := fmt.Sprintf(
		"%s/%s_%s_%d.wav",
		targetDir,
		task.TaskName,
		task.ID,
		task.EditTime.Unix(),
	)
	processedText, err := language.ConvertText(
		task.Text,
		task.TaskLanguage,
	)
	if err != nil {
		return err
	}
	var phontPath string
	if task.CharacterID != nil {
		characterList := characters.CharacterList.GetData()
		character, exists := characterList[*task.CharacterID]
		if !exists {
			return fmt.Errorf(
				"character %s does not exists",
				*task.CharacterID,
			)
		}
		phontPath, err = PhontFile(
			phontsDir,
			character.PhontName,
		)
		if err != nil {
			return err
		}
	} else if task.PhontName != nil {
		phontPath, err = PhontFile(
			phontsDir,
			*task.PhontName,
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
		task.Speed,
		phontPath,
		fileName,
		processedText,
	)
	err = generator.GenerateWav()
	if err != nil {
		return err
	}
	task.ResultFile = &fileName
	return nil
}
