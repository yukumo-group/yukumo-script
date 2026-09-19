package sequence

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"

	"golang.org/x/sync/errgroup"
)

// Task defines the list of the task
type Task struct {
	sync.RWMutex
	TaskName       string      `json:"taskName"`
	ID             string      `json:"id"`
	CreateTime     time.Time   `json:"createdTime"`
	EditTime       time.Time   `json:"editTime"`
	AllSentences   []*Sentence `json:"allSentences"`
	Config         *RawConfig  `json:"config"`
	SentenceAudios []string    `json:"sentenceAudios"`
	taskConfig     *TaskConfig
	wavDir         string
}

// NewTask chreates new task
func NewSequenceTask(
	taskName string,
	config *RawConfig,
	wavDir string,
) (*Task, error) {
	newTaskID := uuid.NewString()
	if config == nil {
		return nil, errors.New(
			"you cannot pass the config as nil",
		)
	}
	processedConfig, err := config.ToTaskConfig()
	if err != nil {
		return nil, err
	}
	return &Task{
		TaskName:       taskName,
		Config:         config,
		taskConfig:     processedConfig,
		ID:             newTaskID,
		CreateTime:     time.Now(),
		EditTime:       time.Now(),
		AllSentences:   []*Sentence{},
		SentenceAudios: []string{},
		wavDir:         wavDir,
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
) (string, error) {
	task.Lock()
	defer task.Unlock()
	data, err := json.Marshal(task)
	if err != nil {
		return "", err
	}
	filePath := task.GenerateFileName(
		targetDir,
	)
	err = os.WriteFile(
		filePath,
		data,
		0644,
	)
	return filePath, err
}

// ConvertToTasks converts sentence to tasks of this task
func (task *Task) ConvertToTasks() ([]tasks.Task, error) {
	task.RLock()
	defer task.RUnlock()
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
	sentence *Sentence,
) {
	task.Lock()
	defer task.Unlock()
	task.EditTime = time.Now()
	task.AllSentences = append(task.AllSentences, sentence)
}

// InsertSentence inserts one sentence after the idx index into the task
func (task *Task) InsertSentence(
	sentence *Sentence,
	idx int,
) {
	task.Lock()
	defer task.Unlock()
	task.EditTime = time.Now()
	if idx < len(task.AllSentences)-1 {
		task.AllSentences = append(
			task.AllSentences[:idx+1],
			append(
				[]*Sentence{
					sentence,
				},
				task.AllSentences[idx+1:]...,
			)...,
		)
	} else {
		task.AllSentences = append(
			task.AllSentences,
			sentence,
		)
	}
}

// GenerateWavName generates name for the result wav file
func (task *Task) GenerateWavFileName(
	targetDir string,
) string {
	tmpID := uuid.NewString()
	return fmt.Sprintf(
		"%s/%s_%s_%d_%s.wav",
		targetDir,
		task.TaskName,
		task.ID,
		task.EditTime.UnixNano(),
		tmpID,
	)
}

// Generate generates the tasks
func (task *Task) Generate(
	ctx context.Context,
	phontsDir string,
	targetDir string,
) error {
	// Convert sentences to task
	subtasks, err := task.ConvertToTasks()
	if err != nil {
		return err
	}
	task.Lock()
	defer task.Unlock()
	task.EditTime = time.Now()
	group, ctx := errgroup.WithContext(
		ctx,
	)
	group.SetLimit(
		runtime.NumCPU() * 2,
	)
	// Generate here
	for i := range subtasks {
		taskIdx := i
		group.Go(
			func() error {
				err := subtasks[taskIdx].Generate(
					ctx,
					phontsDir,
					task.wavDir,
				)
				return err
			},
		)
	}
	err = group.Wait()
	if err != nil {
		return err
	}
	// Combine audios
	resultPath := []string{}
	for i := range subtasks {
		resultFile := subtasks[i].GetResultFile()
		if resultFile == nil {
			return fmt.Errorf(
				"%d file not generated",
				i,
			)
		}
		resultPath = append(resultPath, *resultFile)
	}
	wavFileName := task.GenerateWavFileName(
		targetDir,
	)
	err = edit.SpliceAudios(
		wavFileName,
		resultPath,
	)
	return err
}

// GetTaskName gets task name
func (task *Task) GetTaskName() string {
	task.RLock()
	defer task.RUnlock()
	result := task.TaskName
	return result
}

// GetAllSentences gets all sentences
func (task *Task) GetAllSentences() []*Sentence {
	task.RLock()
	defer task.RUnlock()
	result := slices.Clone(
		task.AllSentences,
	)
	return result
}

// ToPreviewInfo converts sequence to info that can be previewed
func (task *Task) ToPreviewInfo() *SequenceInfo {
	task.RLock()
	defer task.RUnlock()
	var languageSet language.Language
	if task.taskConfig == nil {
		languageSet = language.Chinese
	} else {
		languageSet = task.taskConfig.TaskLanguage
	}
	result := &SequenceInfo{
		Language: languageSet,
		AllSentences: slices.Clone(
			task.AllSentences,
		),
	}
	return result
}
