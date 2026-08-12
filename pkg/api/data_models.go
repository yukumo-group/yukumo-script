package api

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
