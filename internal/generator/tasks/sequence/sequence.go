package sequence

import (
	"sync"
	"time"

	"github.com/yukumo-group/yukumo-script/internal/characters"
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
	characters  *characters.Characters
}

// NewTask chreates new task
func NewSequenceTask(
	TaskName string,
) *Task {
	return &Task{
		TaskName: TaskName,
	}
}
