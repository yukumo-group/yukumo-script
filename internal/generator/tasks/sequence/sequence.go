package sequence

import (
	"sync"
	"time"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
)

// Task defines the list of the task
type Task struct {
	sync.RWMutex
	TaskName    string             `json:"taskName"`
	ID          string             `json:"id"`
	CreatedTime time.Time          `json:"createdTime"`
	EditTime    time.Time          `json:"editTime"`
	AllTasks    map[int]tasks.Task `json:"allTasks"`
	Config      *RawConfig         `json:"config"`
	taskConfig  *TaskConfig
}

// NewTask chreates new task
func NewSequenceTask(
	taskName string,
	config *RawConfig,
) (*Task, error) {
	processedConfig, err := config.ToTaskConfig()
	if err != nil {
		return nil, err
	}
	return &Task{
		TaskName:   taskName,
		Config:     config,
		taskConfig: processedConfig,
	}, nil
}
