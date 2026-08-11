package api

import (
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/yukumo-group/yukumo-script/pkg/characters"
	"github.com/yukumo-group/yukumo-script/pkg/example"
	"github.com/yukumo-group/yukumo-script/pkg/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/pkg/language"
	"github.com/yukumo-group/yukumo-script/pkg/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
)

// Init initializes runtime dirs, examples, phont map, characters, and tasks.
func Init() error {
	InitRuntimeDirs()

	dir, err := phontsmanager.GetAllPhonts(utils.PhontsDir)
	if err != nil {
		return err
	}
	if err := example.GenerateExamples(
		context.Background(),
		utils.ExampleDir,
		utils.PhontsDir,
		dir,
	); err != nil {
		return err
	}
	if err := InitPhontMap(); err != nil {
		return err
	}

	characters.CharacterList.SetTargetFile(
		utils.DataDir,
		utils.CharactersFile,
	)
	if err := characters.CharacterList.ReadData(); err != nil {
		return err
	}
	return InitTaskManager()
}

// GenerateByPhont converts text, generates a wav via AquesTalk2, and registers the task.
func GenerateByPhont(params GenerateByPhontParams) (*GenerateByPhontResult, error) {
	task, err := PrepareGenerateByPhont(params)
	if err != nil {
		return nil, err
	}
	phontPath, err := PhontPath(*task.PhontName)
	if err != nil {
		return nil, err
	}
	if err := task.Generate(phontPath, utils.ResultDir); err != nil {
		return nil, err
	}
	if task.ResultFile == nil {
		return nil, fmt.Errorf("result file path is nil after generation")
	}
	taskFile, err := RegisterGeneratedTask(task)
	if err != nil {
		return nil, err
	}
	return &GenerateByPhontResult{
		ResultFile: *task.ResultFile,
		TaskFile:   taskFile,
	}, nil
}

// InitRuntimeDirs creates the runtime directories used by CLI and clib.
func InitRuntimeDirs() {
	utils.InitializeDirectory(utils.PhontsDir)
	utils.InitializeDirectory(utils.ResultDir)
	utils.InitializeDirectory(utils.WavsDir)
	utils.InitializeDirectory(utils.DataDir)
	utils.InitializeDirectory(utils.ExampleDir)
	utils.InitializeDirectory(utils.ImagesDir)
	utils.InitializeDirectory(utils.TaskDir)
	utils.InitializeDirectory(utils.SingleSentenceDir)
}

// InitPhontMap loads phont name → file mappings from PhontsDir.
func InitPhontMap() error {
	return phontsmanager.InitializePhontNameToFileName(utils.PhontsDir)
}

// InitTaskManager configures and loads the single-sentence task manager.
func InitTaskManager() error {
	singlesentence.Manager.SetTargetFile(
		utils.TaskDir,
		utils.SingleSentenceTasksFile,
	)
	return singlesentence.Manager.ReadData()
}

// ListPhonts returns available phont names.
func ListPhonts() []string {
	return phontsmanager.PhontNameToFileName.GetAllKeys()
}

// ListTasks returns registered single-sentence task names.
func ListTasks() []string {
	return slices.Collect(maps.Keys(singlesentence.Manager.GetAllTasks()))
}

// GenerateByPhontParams holds inputs for GenerateByPhont.
type GenerateByPhontParams struct {
	TaskName  string
	Text      string
	Language  int
	Speed     int
	PhontName string
}

// GenerateByPhontResult holds outputs from a successful generation.
type GenerateByPhontResult struct {
	ResultFile string
	TaskFile   string
}

// PrepareGenerateByPhont validates inputs and creates a task without generating audio.
func PrepareGenerateByPhont(params GenerateByPhontParams) (*singlesentence.Task, error) {
	if singlesentence.Manager.HasTask(params.TaskName) {
		return nil, fmt.Errorf("task %s already exists", params.TaskName)
	}
	_, exists := phontsmanager.PhontNameToFileName.GetValue(params.PhontName)
	if !exists {
		return nil, fmt.Errorf("no such phont %s", params.PhontName)
	}
	processedText, err := language.ConvertText(
		params.Text,
		language.ToLanguage(params.Language),
	)
	if err != nil {
		return nil, err
	}
	phontName := params.PhontName
	return singlesentence.NewSingleSentenceTask(
		processedText,
		nil,
		&phontName,
		params.Speed,
		params.TaskName,
	)
}

// RegisterGeneratedTask saves task metadata and registers it in the manager.
func RegisterGeneratedTask(task *singlesentence.Task) (string, error) {
	taskFile, err := task.SaveFile(utils.SingleSentenceDir)
	if err != nil {
		return "", err
	}
	if err := singlesentence.Manager.NewTask(task.TaskName, taskFile); err != nil {
		return "", err
	}
	return taskFile, nil
}

// PhontPath resolves a phont name to its file path under PhontsDir.
func PhontPath(phontName string) (string, error) {
	return phontsmanager.GetPhontFile(utils.PhontsDir, phontName)
}
