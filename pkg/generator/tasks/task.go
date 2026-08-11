package tasks

// Task defines the interface for task audio generation.
type Task interface {
	// Generate generates an audio file for the task.
	Generate(
		phontPath string,
		targetDir string,
	) error
}
