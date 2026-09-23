package api

import (
	"errors"
	"fmt"
	"os"

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

// ExportConfig exports certain configuration for sequence task.
// You do not need to add suffix in fileName
func ExportConfig(
	configName string,
	targetDir string,
	fileName *string,
) (string, error) {
	configuration, err := sequence.ConfManager.GetConfig(
		configName,
	)
	if err != nil {
		return "", err
	}
	var resultFileName string
	if fileName == nil {
		fetchedFileName, exists := configuration.GetConfigName()
		if !exists {
			return "", errors.New(
				"you cannot export a config without setting config name",
			)
		}
		resultFileName = fetchedFileName
	} else {
		resultFileName = *fileName
	}
	return configuration.ToYAML(
		targetDir,
		resultFileName,
	)
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

// InitializeDefaultSequenceConfig initializes default config for sequences
func InitializeDefaultSequenceConfig() error {
	name, err := sequence.DefaultConfig.GenerateYAMLFileName()
	if err != nil {
		return err
	}
	filePath := fmt.Sprintf(
		"%s/%s",
		filePathForProg.ConfigDir,
		name,
	)
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			configName, exists := sequence.DefaultConfig.GetConfigName()
			if !exists {
				return errors.New(
					"default config must have a config name",
				)
			}
			_, err = sequence.DefaultConfig.ToYAML(
				filePathForProg.ConfigDir,
				configName,
			)
			return err
		}
		return err
	}
	sequence.DefaultConfig, err = sequence.ReadRawConfig(
		filePath,
	)
	return err
}
