package sequence

import (
	"fmt"
	"sync"
)

// WorkSpace defines the work space for the user to edit the sequence
type WorkSpace struct {
	sync.RWMutex
	tasks map[string]*Task
}

// NewWorkSpace creates new work space
func NewWorkSpace() *WorkSpace {
	return &WorkSpace{
		tasks: map[string]*Task{},
	}
}

// LoadFileToWorkSpace loads file to work space.
// Returns ID of this task and error.
func (workspace *WorkSpace) LoadFileToWorkSpace(
	fileName string,
) (string, error) {
	workspace.Lock()
	defer workspace.Unlock()
	task, err := LoadSequenceTaskFromFile(
		fileName,
	)
	if err != nil {
		return "", err
	}
	id := task.GetTaskName()
	_, exists := workspace.tasks[id]
	if exists {
		return "", fmt.Errorf(
			"%s task already exists in the workspace",
			id,
		)
	}
	workspace.tasks[id] = task
	return id, nil
}

// ShowAllTasks shows all the tasks
func (workspace *WorkSpace) ShowAllTasks() map[string]*SequenceInfo {
	workspace.RLock()
	defer workspace.RUnlock()
	result := map[string]*SequenceInfo{}
	for id, task := range workspace.tasks {
		result[id] = task.ToPreviewInfo()
	}
	return result
}
