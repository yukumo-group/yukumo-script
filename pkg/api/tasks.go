package api

import (
	"fmt"
	"maps"
	"slices"
	"time"

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
	if err := singlesentence.Manager.NewTask(task.GetTaskName(), taskFile); err != nil {
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

// RandomTaskName creates a random task name for the task name
func RandomTaskName(
	taskType string,
) string {
	return fmt.Sprintf(
		"New_%s_Task_%d",
		taskType,
		time.Now().Unix(),
	)
}

// GetResultFileForSingleSentenceTask
func GetResultFileForSingleSentenceTask(
	taskName string,
) (string, error) {
	thisTask, err := singlesentence.Manager.GetTask(
		taskName,
	)
	if err != nil {
		return "", err
	}
	resultFile := thisTask.GetResultFile()
	if resultFile == nil {
		return "", fmt.Errorf(
			"task %s does not have result file",
			taskName,
		)
	}
	return *resultFile, nil
}
