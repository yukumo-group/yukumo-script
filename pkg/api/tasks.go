package api

import (
	"maps"
	"slices"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
)

// GetAllTasks gets all the tasks
func GetAllTasks() map[string]string {
	return singlesentence.Manager.GetAllTasks()
}

// ListTasks returns registered single-sentence task names.
func ListTasks() []string {
	return slices.Collect(maps.Keys(singlesentence.Manager.GetAllTasks()))
}

// RegisterGeneratedTask saves task metadata and registers it in the manager.
func RegisterGeneratedTask(task *singlesentence.Task) (string, error) {
	taskFile, err := task.SaveFile(filePathForProg.SingleSentenceDir)
	if err != nil {
		return "", err
	}
	if err := singlesentence.Manager.NewTask(task.TaskName, taskFile); err != nil {
		return "", err
	}
	return taskFile, nil
}

// InitTaskManager configures and loads the single-sentence task manager.
func InitTaskManager() error {
	singlesentence.Manager.SetTargetFile(
		filePathForProg.TaskDir,
		filePathForProg.SingleSentenceTasksFile,
	)
	return singlesentence.Manager.ReadData()
}
