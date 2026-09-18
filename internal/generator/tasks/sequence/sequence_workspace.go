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

// LoadFileToWorkSpace loads file to work space
func (workspace *WorkSpace) LoadFileToWorkSpace(
	fileName string,
) error {
	workspace.Lock()
	defer workspace.Unlock()
	task, err := LoadSequenceTaskFromFile(
		fileName,
	)
	if err != nil {
		return err
	}
	id := task.GetTaskName()
	_, exists := workspace.tasks[id]
	if exists {
		return fmt.Errorf(
			"%s task already exists in the workspace",
			id,
		)
	}
	workspace.tasks[id] = task
	return nil
}
