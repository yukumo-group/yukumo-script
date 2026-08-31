package singlesentence

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"maps"

	"github.com/yukumo-group/yukumo-script/internal/characters"
)

// TaskManager creates new task.
// This links task metadata to the task name.
type TaskManager struct {
	sync.RWMutex `json:"-"`
	Tasks        map[string]string `json:"tasks"`
	fileName     string
}

// NewTaskManager creates new task manager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make(map[string]string),
	}
}

// Manager manages the tasks
var Manager = NewTaskManager()

// SetTargetFile sets the target file to save data
func (manager *TaskManager) SetTargetFile(
	folder string,
	file string,
) {
	manager.Lock()
	defer manager.Unlock()
	manager.fileName = fmt.Sprintf(
		"%s/%s",
		folder,
		file,
	)
}

// Save saves the manager file
func (manager *TaskManager) Save() error {
	manager.RLock()
	defer manager.RUnlock()
	data, errMarshal := json.Marshal(manager)
	if errMarshal != nil {
		return errMarshal
	}
	errWrite := os.WriteFile(
		manager.fileName,
		data,
		0644,
	)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// ReadData reads the data in the file
func (manager *TaskManager) ReadData() error {
	manager.Lock()
	defer manager.Unlock()
	_, errExist := os.Stat(manager.fileName)
	if errExist != nil {
		if os.IsNotExist(errExist) {
			if manager.Tasks == nil {
				manager.Tasks = make(map[string]string)
			}
			data, errMarshal := json.Marshal(manager)
			if errMarshal != nil {
				return errMarshal
			}
			return os.WriteFile(manager.fileName, data, 0644)
		}
		return errExist
	}
	data, errRead := os.ReadFile(manager.fileName)
	if errRead != nil {
		return errRead
	}
	err := json.Unmarshal(data, manager)
	return err
}

// DeleteTask deletes certain task
func (manager *TaskManager) DeleteTask(
	taskName string,
) error {
	manager.Lock()
	_, exists := manager.Tasks[taskName]
	if !exists {
		manager.Unlock()
		return fmt.Errorf(
			"Task with name %s does not exists",
			taskName,
		)
	}
	delete(manager.Tasks, taskName)
	manager.Unlock()
	return manager.Save()
}

// NewTask creates new task
func (manager *TaskManager) NewTask(
	taskName string,
	fileName string,
) error {
	manager.Lock()
	_, exists := manager.Tasks[taskName]
	if exists {
		manager.Unlock()
		return fmt.Errorf(
			"Task with name %s already exists",
			taskName,
		)
	}
	manager.Tasks[taskName] = fileName
	manager.Unlock()
	return manager.Save()
}

// GetAllTasks gets all the tasks
func (manager *TaskManager) GetAllTasks() map[string]string {
	manager.RLock()
	defer manager.RUnlock()
	return maps.Clone(manager.Tasks)
}

// HasTask checks if certain task with task name exists
func (manager *TaskManager) HasTask(
	taskName string,
) bool {
	manager.RLock()
	defer manager.RUnlock()
	_, exists := manager.Tasks[taskName]
	return exists
}

// GetTask gets the task with certain task name
func (manager *TaskManager) GetTask(
	taskName string,
	charactersManager *characters.Characters,
) (*Task, error) {
	manager.RLock()
	defer manager.RUnlock()
	thisTaskfilePath, exists := manager.Tasks[taskName]
	if !exists {
		return nil, fmt.Errorf(
			"task with name %s does not exists",
			taskName,
		)
	}
	return NewSingleSentenceTaskFromFile(
		thisTaskfilePath,
		charactersManager,
	)
}
