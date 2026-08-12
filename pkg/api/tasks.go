package api

import (
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
)

// GetAllTasks gets all the tasks
func GetAllTasks() map[string]string {
	return singlesentence.Manager.GetAllTasks()
}
