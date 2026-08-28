package tasks

import (
	"context"
)

// Task defines the interface for task audio generation.
type Task interface {
	// Generate generates an audio file for the task.
	Generate(
		ctx context.Context,
		phontsDir string,
		targetDir string,
	) error
	// UseEffect uses effects on the generated audio file
	UseEffect(
		ctx context.Context,
		targetDir string,
		tempDir string,
	) error
}
