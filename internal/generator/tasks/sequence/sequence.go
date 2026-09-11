package sequence

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"

	"github.com/google/uuid"
)

// Task defines the list of the task
type Task struct {
	sync.RWMutex
	TaskName     string     `json:"taskName"`
	ID           string     `json:"id"`
	CreateTime   time.Time  `json:"createdTime"`
	EditTime     time.Time  `json:"editTime"`
	AllSentences []Sentence `json:"allSentences"`
	Config       *RawConfig `json:"config"`
	taskConfig   *TaskConfig
	wavDir       string
}

// NewTask chreates new task
func NewSequenceTask(
	taskName string,
	config *RawConfig,
	wavDir string,
) (*Task, error) {
	newTaskID := uuid.NewString()
	processedConfig, err := config.ToTaskConfig()
	if err != nil {
		return nil, err
	}
	return &Task{
		TaskName:     taskName,
		Config:       config,
		taskConfig:   processedConfig,
		ID:           newTaskID,
		CreateTime:   time.Now(),
		EditTime:     time.Now(),
		AllSentences: []Sentence{},
		wavDir:       wavDir,
	}, nil
}

// LoadSequenceTaskFromFile loads sequence task from file
func LoadSequenceTaskFromFile(
	fileName string,
) (*Task, error) {
	var Result Task
	data, err := os.ReadFile(
		fileName,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(
		data,
		&Result,
	)
	if err != nil {
		return nil, err
	}
	return &Result, nil
}

// GenerateFileName generates filename of metadata for this task
func (task *Task) GenerateFileName(
	targetDir string,
) string {
	return fmt.Sprintf(
		"%s/%s_%s_%d.json",
		targetDir,
		task.TaskName,
		task.ID,
		task.CreateTime.UnixNano(),
	)
}

// SaveFile saves the task into file
func (task *Task) SaveFile(
	targetDir string,
) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}
	filePath := task.GenerateFileName(
		targetDir,
	)
	err = os.WriteFile(
		filePath,
		data,
		0644,
	)
	return err
}

// ConvertToTasks converts sentence to tasks of this task
func (task *Task) ConvertToTasks() ([]tasks.Task, error) {
	result := []tasks.Task{}
	for _, sentence := range task.AllSentences {
		newTask, err := sentence.ToTask(
			task.wavDir,
			task.TaskName,
			audio.DefaultAudioInfo,
			task.taskConfig.DefaultSpeed,
			task.taskConfig.Characters,
			task.taskConfig.TaskLanguage,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, newTask)
	}
	return result, nil
}

// AddSentence adds one single sentence
func (task *Task) AddSentence(
	sentence Sentence,
) {
	task.Lock()
	defer task.Unlock()
	task.AllSentences = append(task.AllSentences, sentence)
}
