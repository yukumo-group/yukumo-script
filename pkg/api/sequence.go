package api

import (
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
)

// InitSequenceTaskConfigManager intialises config manager for sequence task
func InitSequenceTaskConfigManager() {
	sequence.ConfManager.SetConfigFilePath(
		filePathForProg.TaskDir,
		filePathForProg.ConfigManagerFile,
	)
}

// InitSequenceTaskManager initialises the task manager for sequence tasks
func InitSequenceTaskManager() error {
	return sequence.TasksManager.Load(
		filePathForProg.SequenceDir,
		filePathForProg.SequenceTaskManagerFile,
	)
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

// GetAllSequenceConfigs gets all the configurations for sequence task
func GetAllSequenceConfigs() []string {
	return sequence.ConfManager.GetAllConfigNames()
}

// GetSequenceConfig gets certain configuration for sequence task
func GetSequenceConfig(
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

// ListAllTasks lists all the tasks
// This shows all the task ids
func ListAllTasks() []string {
	tasks := sequence.TasksManager.GetAllTasks()
	result := []string{}
	for i := range tasks {
		result = append(result, i)
	}
	return result
}
