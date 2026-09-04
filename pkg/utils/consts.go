package utils

const (
	// RuntimeDir defines the directory to store the file generated when running
	RuntimeDir string = "runtime"
	// ExampleDir defines the directory to store the example audios
	ExampleDir string = "runtime/examples"
	// PhontsDir defines the directory to store the phont files
	PhontsDir string = "runtime/phonts"
	// ResultDir defines the directory to store the generated result files
	ResultDir string = "runtime/result"
	// WavsDir defines the directory to store the generated temporary wav files
	WavsDir string = "runtime/wav"
	// DataDir defines the directory to store data such as characters
	DataDir string = "runtime/data"
	// ConfigDir defines the directory to store configurations
	ConfigDir string = "runtime/data/config"
	// ImagesDir define the directory to store images such as profiles
	ImagesDir string = "runtime/data/images"
	// TaskDir define the directory to store task data
	TaskDir string = "runtime/data/tasks"
	// SingleSentenceDir defines the directory to store task files of single sentences
	SingleSentenceDir string = "runtime/data/tasks/single_sentence"
	// SequenceDir defines the directory to store task files of sequence task
	SequenceDir string = "runtime/data/tasks/sequence"
	// SingleSentenceTasksFile defines the file to store the info of tasks to make task management easier
	SingleSentenceTasksFile string = "single_sentence_tasks.json"
	// ConfPath defines the path of the conf file
	ConfPath string = "conf.ini"
	// CharactersFile defines the name of the file for storing characters
	CharactersFile string = "characters.json"
	// EnglishTexts defines the path of the english texts file
	EnglishTexts string = "text_en.properties"
)
