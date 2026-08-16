package singlesentence

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/yukumo-group/yukumo-script/internal/generator/aquestalk2"
)

// Task defines the task of generating a single sentence
type Task struct {
	ID          string    `json:"id"`
	TaskName    string    `json:"taskName"`
	Text        string    `json:"text"`
	Speed       int       `json:"speed"`
	CreateTime  time.Time `json:"createTime"`
	EditTime    time.Time `json:"editTime"`
	CharacterID *string   `json:"characterID"`
	PhontName   *string   `json:"phontName"`
	ResultFile  *string   `json:"resultFile"`
}

// NewSingleSentenceTask creates new single sentence task
func NewSingleSentenceTask(
	text string,
	characterID *string,
	phontName *string,
	speed int,
	taskName string,
) (*Task, error) {
	if phontName == nil && characterID == nil {
		return nil, errors.New(
			"you have to choose at least one way to generate the audio",
		)
	}
	id := uuid.NewString()
	return &Task{
		ID:          id,
		Text:        text,
		CreateTime:  time.Now(),
		EditTime:    time.Now(),
		CharacterID: characterID,
		PhontName:   phontName,
		Speed:       speed,
		ResultFile:  nil,
		TaskName:    taskName,
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
	phontPath string,
	targetDir string,
) error {
	fileName := fmt.Sprintf(
		"%s/%s_%s_%d.wav",
		targetDir,
		task.TaskName,
		task.ID,
		task.CreateTime.Unix(),
	)
	generator := aquestalk2.NewGenerator(
		task.Speed,
		phontPath,
		fileName,
		task.Text,
	)
	err := generator.GenerateWav()
	task.EditTime = time.Now()
	task.ResultFile = &fileName
	return err
}
