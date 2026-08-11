//go:build !windows

package api

import "errors"

// Init initializes runtime dirs, phont map, and tasks (non-Windows; no AquesTalk2 examples).
func Init() error {
	InitRuntimeDirs()
	if err := InitPhontMap(); err != nil {
		return err
	}
	return InitTaskManager()
}

// GenerateByPhont is only supported on Windows (AquesTalk2).
func GenerateByPhont(params GenerateByPhontParams) (*GenerateByPhontResult, error) {
	return nil, errors.New("GenerateByPhont is only supported on Windows")
}
