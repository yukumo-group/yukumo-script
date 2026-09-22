package api

import (
	"errors"

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

// ListAllTasks lists all the tasks stored
// This shows all the task ids
func ListAllTasks() []string {
	tasks := sequence.TasksManager.GetAllTasks()
	result := []string{}
	for i := range tasks {
		result = append(result, i)
	}
	return result
}

// AddOrInsertSentenceTo adds or inserts sentence to workspace.
// If insert is set to true, you cannot pass nil to idx.
// If insert is set to true, the sentence will be inserted to the position after idx.
// If insert is set to false, the sentence will directly be added to the end of all sentences
func AddOrInsertSentenceTo(
	taskID string,
	sentence *sequence.Sentence,
	insert bool,
	idx *int,
) (*sequence.SequenceInfo, error) {
	if insert {
		if idx == nil {
			return nil, errors.New(
				"you cannot pass a nil to idx if insert is set to true",
			)
		}
		info := sequence.MainWorkSpace.InsertSentence(
			taskID,
			sentence,
			*idx,
		)
		return info, nil
	}
	info := sequence.MainWorkSpace.AddSentence(
		taskID,
		sentence,
	)
	return info, nil
}

// ListAllTasksInWorkspace lists all tasks (map[`taskID`]`info“) that are in work space.
// You can use this function to refresh the workspace.
func ListAllTasksInWorkspace() map[string]*sequence.SequenceInfo {
	return sequence.MainWorkSpace.ShowAllTasks()
}
