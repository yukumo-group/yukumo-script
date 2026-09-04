package sequence

import (
	"fmt"
	"sync"
)

// ConfigManager defines the manager for configuration
type ConfigManager struct {
	sync.RWMutex
	Data     map[string]string
	filePath string
}

// NewConfigManager creates new config manager for sequence tasks
func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		Data: make(map[string]string),
	}
}

// SetConfigPath sets the path for storing configurations
func (manager *ConfigManager) SetConfigPath(
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
