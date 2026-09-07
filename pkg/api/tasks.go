package api

import (
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
)

// InitSequenceTaskConfigManager intialises config manager for sequence task
func InitSequenceTaskConfigManager() {
	sequence.ConfManager.SetConfigFilePath(
		filePathForProg.TaskDir,
		filePathForProg.ConfigManagerFile,
	)
}

// GetAllTasks gets all the single sentence tasks
func GetAllTasks() map[string]string {
	return singlesentence.Manager.GetAllTasks()
}

// ListTasks returns registered single-sentence task names.
func ListTasks() []string {
	return slices.Collect(maps.Keys(singlesentence.Manager.GetAllTasks()))
}

// RegisterGeneratedTask saves single sentence task metadata and registers it in the manager.
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
		time.Now().UnixNano(),
	)
}

// GetResultFileForSingleSentenceTask
func GetResultFileForSingleSentenceTask(
	taskName string,
) (string, error) {
	thisTask, err := singlesentence.Manager.GetTask(
		taskName,
		characters.CharacterList,
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

// GetTask gets single sentence task from manager
func GetTask(
	taskName string,
) (*TaskInfo, error) {
	task, err := singlesentence.Manager.GetTask(
		taskName,
		characters.CharacterList,
	)
	if err != nil {
		return nil, err
	}
	result := SingleSentenceTaskToTaskInfo(
		task,
	)
	return result, nil
}

// AddSequenceTaskConfigFromFile adds config for sequence task from file.
// The fileName can be any path in the computer as it will be copied to conf manager
func AddSequenceTaskConfigFromFile(
	fileName string,
) error {
	err := sequence.ConfManager.AddFile(
		fileName,
		filePathForProg.ConfigDir,
	)
	return err
}

// GetAllConfigs gets all the configurations for sequence task
func GetAllConfigs() []string {
	return sequence.ConfManager.GetAllConfigNames()
}

// GetConfig gets certain configuration for sequence task
func GetConfig(
	configName string,
) (*sequence.RawConfig, error) {
	configuration, err := sequence.ConfManager.GetConfig(
		configName,
	)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}
