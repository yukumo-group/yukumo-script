package sequence

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Task defines the list of the task
type Task struct {
	sync.RWMutex
	TaskName     string     `json:"taskName"`
	ID           string     `json:"id"`
	CreatedTime  time.Time  `json:"createdTime"`
	EditTime     time.Time  `json:"editTime"`
	AllSentences []Sentence `json:"allSentences"`
	Config       *RawConfig `json:"config"`
	taskConfig   *TaskConfig
}

// NewTask chreates new task
func NewSequenceTask(
	taskName string,
	config *RawConfig,
) (*Task, error) {
	newTaskID := uuid.NewString()
	processedConfig, err := config.ToTaskConfig()
	if err != nil {
		return nil, err
	}
	return &Task{
		TaskName:     taskName,
		Config:       config,
		taskConfig:   processedConfig,
		ID:           newTaskID,
		CreatedTime:  time.Now(),
		EditTime:     time.Now(),
		AllSentences: []Sentence{},
	}, nil
}

// LoadSequenceTaskFromFile loads sequence task from file
func LoadSequenceTaskFromFile(
	fileName string,
) (*Task, error) {
	var Result Task
	data, err := os.ReadFile(
		fileName,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(
		data,
		&Result,
	)
	if err != nil {
		return nil, err
	}
	return &Result, nil
}

// AddSentence adds one single sentence
func (task *Task) AddSentence(
	sentence Sentence,
) {
	task.Lock()
	defer task.Unlock()
	task.AllSentences = append(task.AllSentences, sentence)
}
