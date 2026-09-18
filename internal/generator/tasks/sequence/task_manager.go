package sequence

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// TaskManager manages the task
type TaskManager struct {
	sync.RWMutex
	Tasks    map[string]string `json:"tasks"`
	filePath string
}

// NewTaskManager creates new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make(map[string]string),
	}
}

// Load loads manager from file
func (manager *TaskManager) Load(
	targetDir string,
	targetFilePath string,
) error {
	manager.Lock()
	defer manager.Unlock()
	filePath := fmt.Sprintf(
		"%s/%s",
		targetDir,
		targetFilePath,
	)
	manager.filePath = filePath
	_, err := os.Stat(
		filePath,
	)
	if err != nil {
		if os.IsNotExist(err) {
			data, err := json.Marshal(
				manager,
			)
			if err != nil {
				return err
			}
			err = os.WriteFile(
				filePath,
				data,
				0644,
			)
		}
		return err
	}
	data, err := os.ReadFile(
		filePath,
	)
	if err != nil {
		return err
	}
	err = json.Unmarshal(
		data,
		manager,
	)
	return err
}

// AddNewTask adds new task
func (manager *TaskManager) AddNewTask(
	taskName string,
	fileName string,
) {
	manager.Lock()
	defer manager.Unlock()
	manager.Tasks[taskName] = fileName
}

// GetAllTasks gets all the tasks for manager
func (manager *TaskManager) GetAllTasks() map[string]string {
	manager.RLock()
	defer manager.RUnlock()
	result := manager.Tasks
	return result
}

// Save saves the manager to file
func (manager *TaskManager) Save() error {
	manager.Lock()
	defer manager.Unlock()
	_, err := json.Marshal(
		manager,
	)
	if err != nil {
		return err
	}
	return nil
}
