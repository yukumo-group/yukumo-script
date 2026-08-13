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
	ID         string                          `json:"id"`
	CreatedAt  time.Time                       `json:"time"`
	AllTasks   map[int]tasks.Task              `json:"allTasks"`
	Characters map[string]characters.Character `json:"characters"`
}
