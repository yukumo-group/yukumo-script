package api

import (
	"time"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

// filePathForProg stores the file path needed
var filePathForProg = FilePathes{}

// GenerateByPhontParams holds inputs for GenerateByPhont.
type GenerateByPhontParams struct {
	TaskName  string
	Text      string
	Language  int
	Speed     int
	PhontName string
}

// GenerateByCharacterParams holds inputs for GenerateByPhont.
type GenerateByCharacterParams struct {
	TaskName      string
	Text          string
	Language      int
	Speed         int
	CharacterName string
}

// GenerateResult holds outputs from a successful generation.
type GenerateResult struct {
	ResultFile string
	TaskFile   string
}

// NewGenerateByPhontParams creates new GenerateByPhontParams
func NewGenerateByPhontParams(
	taskName string,
	text string,
	language int,
	speed int,
	phontName string,
) *GenerateByPhontParams {
	return &GenerateByPhontParams{
		TaskName:  taskName,
		Text:      text,
		Language:  language,
		Speed:     speed,
		PhontName: phontName,
	}
}

// NewGenerateByCharacterParams creates new GenerateByPhontParams
func NewGenerateByCharacterParams(
	taskName string,
	text string,
	language int,
	speed int,
	characterName string,
) *GenerateByCharacterParams {
	return &GenerateByCharacterParams{
		TaskName:      taskName,
		Text:          text,
		Language:      language,
		Speed:         speed,
		CharacterName: characterName,
	}
}

// GenerateEmptyParams defines the parameters for generating empty audio
type GenerateEmptyParams struct {
	AudioInfo *audio.Info
}

// FilePathes stores the pathes needed by the program
type FilePathes struct {
	RuntimeDir              string
	ExampleDir              string
	PhontsDir               string
	ResultDir               string
	WavsDir                 string
	DataDir                 string
	ImagesDir               string
	TaskDir                 string
	SingleSentenceDir       string
	SingleSentenceTasksFile string
	ConfPath                string
	CharactersFile          string
	EnglishTexts            string
	SequenceDir             string
	ConfigDir               string
}

// TaskInfo defines the information for task
type TaskInfo struct {
	TaskName    string
	CreateTime  time.Time
	EditTime    time.Time
	Generated   bool
	EffectsUsed bool
	Text        string
	Effects     []*edit.AudioEffect
}

// SingleSentenceTaskToTaskInfo converts single sentence task to task info
func SingleSentenceTaskToTaskInfo(
	task *singlesentence.Task,
) *TaskInfo {
	generated := task.IsGenerated()
	effectUsed := task.IsEffectUsed()
	task.RLock()
	defer task.RUnlock()
	return &TaskInfo{
		TaskName:    task.TaskName,
		CreateTime:  task.CreateTime,
		EditTime:    task.EditTime,
		Generated:   generated,
		Text:        task.Text,
		EffectsUsed: effectUsed,
		Effects:     task.EffectList,
	}
}
