package sequence

import (
	"sync"
)

// ConfigManager defines the manager for configuration
type ConfigManager struct {
	sync.RWMutex
	Data map[string]string
}
