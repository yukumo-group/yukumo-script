package sequence

import (
	"sync"
)

// WorkSpace defines the work space for the user to edit the sequence
type WorkSpace struct {
	sync.RWMutex
	tasks map[string]Task
}

// NewWorkSpace creates new work space
func NewWorkSpace() *WorkSpace {
	return &WorkSpace{
		tasks: map[string]Task{},
	}
}

// LoadFileToWorkSpace loads file to work space
func LoadFileToWorkSpace(
	fileName string,
) error {
	return nil
}
