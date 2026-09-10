package sequence

import (
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

// AddSentence adds one single sentence
func (task *Task) AddSentence(
	sentence Sentence,
) {
	task.Lock()
	defer task.Unlock()
	task.AllSentences = append(task.AllSentences, sentence)
}
