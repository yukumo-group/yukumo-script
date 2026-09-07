package sequence

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"sync"

	"github.com/yukumo-group/yukumo-script/pkg/utils/osoperation"
)

// ConfigManager defines the manager for configuration
type ConfigManager struct {
	sync.RWMutex
	Data     map[string]string
	filePath string
}

// NewConfigManagerFromFile loads new config manager from json file
func NewConfigManagerFromFile(
	filePath string,
) (*ConfigManager, error) {
	newConfigManager := ConfigManager{}
	fileData, err := os.ReadFile(
		filePath,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(
		fileData,
		&newConfigManager,
	)
	if err != nil {
		return nil, err
	}
	newConfigManager.filePath = filePath
	return &newConfigManager, nil
}

// NewConfigManager creates new config manager for sequence tasks
func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		Data: make(map[string]string),
	}
}

// SetManagerFilePath sets the manager path for storing configurations
func (manager *ConfigManager) SetConfigFilePath(
	dir string,
	fileName string,
) {
	manager.Lock()
	defer manager.Unlock()
	manager.filePath = fmt.Sprintf(
		"%s/%s",
		dir,
		fileName,
	)
}

// AddRawConfig adds raw configuration to the manager
func (manager *ConfigManager) AddRawConfig(
	rawTask *RawConfig,
	confDir string,
) error {
	manager.Lock()
	defer manager.Unlock()
	configName, exists := rawTask.GetConfigName()
	if !exists {
		return errors.New(
			"you cannot directly add a raw config without config name",
		)
	}
	_, exists = manager.Data[configName]
	if exists {
		return fmt.Errorf(
			"config named %s already exists",
			configName,
		)
	}
	fileName, err := rawTask.GenerateYAMLFileName()
	if err != nil {
		return err
	}
	savedFilePath, err := rawTask.ToYAML(
		confDir,
		fileName,
	)
	if err != nil {
		return err
	}
	manager.Data[configName] = savedFilePath
	return nil
}

// AddFile adds new file to it
func (manager *ConfigManager) AddFile(
	rawConfigFilePath string,
	confDir string,
) error {
	manager.Lock()
	defer manager.Unlock()
	configData, err := ReadRawConfig(
		rawConfigFilePath,
	)
	if err != nil {
		return err
	}
	configName, exists := configData.GetConfigName()
	if !exists {
		return errors.New(
			"no task name is included in this config",
		)
	}
	_, exists = manager.Data[configName]
	if exists {
		return fmt.Errorf(
			"config named %s already exists",
			configName,
		)
	}
	err = osoperation.CopyFile(
		rawConfigFilePath,
		confDir,
		configName,
		"yaml",
	)
	if err != nil {
		return err
	}
	fileName, err := osoperation.GetNewFilePath(
		confDir,
		configName,
		"yaml",
	)
	if err != nil {
		return err
	}
	manager.Data[configName] = fileName
	return nil
}

// SaveFile saves the data in the file
func (manager *ConfigManager) SaveFile() error {
	manager.Lock()
	defer manager.Unlock()
	data, err := json.Marshal(manager)
	if err != nil {
		return err
	}
	err = os.WriteFile(
		manager.filePath,
		data,
		0644,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetConfig gets the config from file managed by this manager
func (manager *ConfigManager) GetConfig(
	configName string,
) (*RawConfig, error) {
	yamlFileName, exists := manager.Data[configName]
	if !exists {
		return nil, fmt.Errorf(
			"%s config name does not exists",
			configName,
		)
	}
	resultRawConfig, err := ReadRawConfig(yamlFileName)
	if err != nil {
		return nil, err
	}
	return resultRawConfig, nil
}

// GetData gets the data from config manager
func (manager *ConfigManager) GetData() map[string]string {
	manager.RLock()
	defer manager.RUnlock()
	return maps.Clone(manager.Data)
}

// GetAllConfigNames gets all the config names
func (manager *ConfigManager) GetAllConfigNames() []string {
	manager.RLock()
	defer manager.RUnlock()
	configNamesList := []string{}
	for configName := range manager.Data {
		configNamesList = append(configNamesList, configName)
	}
	return configNamesList
}
