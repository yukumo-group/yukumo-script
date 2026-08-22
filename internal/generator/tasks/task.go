package tasks

import (
	"context"
)

// Task defines the interface for task audio generation.
type Task interface {
	// Generate generates an audio file for the task.
	Generate(
		ctx context.Context,
		phontPath string,
		targetDir string,
	) error
}
