package api

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
	ConfDir                 string
	CharactersFile          string
	EnglishTexts            string
}
