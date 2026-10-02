package api

import (
	"errors"
	"time"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

// filePathForProg stores the file path needed
var filePathForProg = &FilePathes{}

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

// ConfigLoadParam defines how to load config
type ConfigLoadParam struct {
	Method              LoadConfigMethod
	ConfigFilePath      *string
	ConfigNameInManager *string
}

// NewConfigLoadParam defines the param to load config
func NewConfigLoadParam(
	loadMethod LoadConfigMethod,
	configFilePath *string,
	configNameInManager *string,
) (*ConfigLoadParam, error) {
	if loadMethod == FromFile && configFilePath == nil {
		return nil, errors.New(
			"you cannot pass a nil path for config file if load method is from file",
		)
	}
	if loadMethod == FromManager && configNameInManager == nil {
		return nil, errors.New(
			"you cannot pass a nil config name if load method is from manager",
		)
	}
	return &ConfigLoadParam{
		Method:              loadMethod,
		ConfigFilePath:      configFilePath,
		ConfigNameInManager: configNameInManager,
	}, nil
}

// ToRawConfig converts load param to config
func (param *ConfigLoadParam) ToRawConfig() (*sequence.RawConfig, error) {
	switch param.Method {
	case FromFile:
		if param.ConfigFilePath == nil {
			return nil, errors.New(
				"you cannot load config from file when the file path pased is nil",
			)
		}
		return sequence.ReadRawConfig(
			*param.ConfigFilePath,
		)
	case FromManager:
		if param.ConfigNameInManager == nil {
			return nil, errors.New(
				"you cannot load config from file when the config name pased is nil",
			)
		}
		return sequence.ConfManager.GetConfig(
			*param.ConfigNameInManager,
		)
	case Default:
		return sequence.DefaultConfig, nil
	default:
		return nil, errors.New(
			"method not supported",
		)
	}
}

// GenerateSequenceParams defines the parameters for generating sequence task
type GenerateSequenceParams struct {
	TaskName string
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
	ConfigManagerFile       string
	PolyphonicsManagerFile  string
	SequenceTaskManagerFile string
	AssetsDir               string
	DictFilePath            string
}

// SetPathes sets the pathes
func (pathes *FilePathes) SetPathes(
	opts ...Option,
) {
	for _, opt := range opts {
		opt(pathes)
	}
}

// Option defines the option for setting the pathes
type Option func(*FilePathes)

// WithDictFilePath defines the setting of file path for dict file
func WithDictFilePath(
	dictFilePath string,
) Option {
	return func(pathes *FilePathes) {
		pathes.DictFilePath = dictFilePath
	}
}

// WithConfigDir defines the setting of config dir
func WithConfigDir(
	configDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.ConfigDir = configDir
	}
}

// WithSequenceDir defines the setting of sequence dir
func WithSequenceDir(
	sequenceDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.SequenceDir = sequenceDir
	}
}

// WithRuntimeDir defines the setting of runtime dir
func WithRuntimeDir(
	runtimeDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.RuntimeDir = runtimeDir
	}
}

// WithExampleDir defines the setting of example dir
func WithExampleDir(
	exampleDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.ExampleDir = exampleDir
	}
}

// WithPhontsDir defines the setting of example dir
func WithPhontsDir(
	phontsDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.PhontsDir = phontsDir
	}
}

// WithResultDir defines the setting of example dir
func WithResultDir(
	resultDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.ResultDir = resultDir
	}
}

// WithWavsDir defines the setting of wavs dir
func WithWavsDir(
	wavsDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.WavsDir = wavsDir
	}
}

// WithDataDir defines the setting of data dir
func WithDataDir(
	dataDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.DataDir = dataDir
	}
}

// WithImagesDir defines the setting of images dir
func WithImagesDir(
	imagesDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.ImagesDir = imagesDir
	}
}

// WithTaskDir defines the setting of task dir
func WithTaskDir(
	taskDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.TaskDir = taskDir
	}
}

// WithSingleSentenceDir defines the setting of single sentence dir
func WithSingleSentenceDir(
	singleSentenceDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.SingleSentenceDir = singleSentenceDir
	}
}

// WithSingleSentenceTasksFile defines the setting of task file for managing single sentence tasks
func WithSingleSentenceTasksFile(
	singleSentenceTasksFile string,
) Option {
	return func(pathes *FilePathes) {
		pathes.SingleSentenceTasksFile = singleSentenceTasksFile
	}
}

// WithConfPath defines the setting of directory for storing configuration
func WithConfPath(
	confPath string,
) Option {
	return func(pathes *FilePathes) {
		pathes.ConfPath = confPath
	}
}

// WithCharactersFile defines the setting of file for storing characters
func WithCharactersFile(
	charactersFile string,
) Option {
	return func(pathes *FilePathes) {
		pathes.CharactersFile = charactersFile
	}
}

// WithConfigManagerFile defines the file of config manager
func WithConfigManagerFile(
	configManagerFile string,
) Option {
	return func(pathes *FilePathes) {
		pathes.ConfigManagerFile = configManagerFile
	}
}

// WithPolyphonicsManagerFile defines the file of polyphonics manager
func WithPolyphonicsManagerFile(
	polyphonicsManagerFile string,
) Option {
	return func(pathes *FilePathes) {
		pathes.PolyphonicsManagerFile = polyphonicsManagerFile
	}
}

// WithSequenceTaskManagerFile defines the file for storing sequence tasks
func WithSequenceTaskManagerFile(
	sequenceTaskManagerFile string,
) Option {
	return func(pathes *FilePathes) {
		pathes.SequenceTaskManagerFile = sequenceTaskManagerFile
	}
}

// WithAssetsDir defines the directory for storing assets such as phont file
func WithAssetsDir(
	assetsDir string,
) Option {
	return func(pathes *FilePathes) {
		pathes.AssetsDir = assetsDir
	}
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
