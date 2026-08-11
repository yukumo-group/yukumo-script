//go:build windows
// +build windows

package tasks

// Task defines the interface for task generation in windows
type Task interface {
	// GenerateWin generates audio file in windows
	GenerateWin(
		phontPath string,
		targetDir string,
	) error
}
