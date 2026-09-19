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
func NewWorkSpace(
	manager *TaskManager,
) *WorkSpace {
	return &WorkSpace{
		tasks: map[string]*Task{},
	}
}

// LoadTaskToWorkSpace loads task to work space.
// Returns ID of this task and error.
func (workspace *WorkSpace) LoadTaskToWorkSpace(
	task *Task,
) (string, error) {
	workspace.Lock()
	defer workspace.Unlock()
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
// return: task id -> sequence info
func (workspace *WorkSpace) ShowAllTasks() map[string]*SequenceInfo {
	workspace.RLock()
	defer workspace.RUnlock()
	result := map[string]*SequenceInfo{}
	for id, task := range workspace.tasks {
		result[id] = task.ToPreviewInfo()
	}
	return result
}

// AddSentence adds sentence to certain task
func (workspace *WorkSpace) AddSentence(
	taskID string,
	newSentence *Sentence,
) *SequenceInfo {
	workspace.Lock()
	defer workspace.Unlock()
	workspace.tasks[taskID].AddSentence(
		newSentence,
	)
	result := workspace.tasks[taskID].ToPreviewInfo()
	return result
}

// InsertSentence inserts sentence to certain task after idx
func (workspace *WorkSpace) InsertSentence(
	taskID string,
	newSentence *Sentence,
	idx int,
) *SequenceInfo {
	workspace.Lock()
	defer workspace.Unlock()
	workspace.tasks[taskID].InsertSentence(
		newSentence,
		idx,
	)
	result := workspace.tasks[taskID].ToPreviewInfo()
	return result
}

// FinishAndSave saves the task in the workspace
func (workspace *WorkSpace) FinishAndSave(
	taskID string,
	targetDir string,
) (string, error) {
	workspace.Lock()
	defer workspace.Unlock()
	path, err := workspace.tasks[taskID].SaveFile(
		targetDir,
	)
	if err != nil {
		return "", err
	}
	delete(workspace.tasks, taskID)
	return path, nil
}
