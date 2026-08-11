//go:build windows
// +build windows

package singlesentence

import (
	"fmt"

	"github.com/yukumo-group/yukumo-script/pkg/generator/generatorwin"
)

// GenerateWin generates the wav file in windows64 system
func (task *Task) GenerateWin(
	phontPath string,
	targetDir string,
) error {
	fileName := fmt.Sprintf(
		"%s/%s_%s_%d.wav",
		targetDir,
		task.TaskName,
		task.ID,
		task.CreateTime.Unix(),
	)
	generator := generatorwin.NewGeneratorWin(
		task.Speed,
		phontPath,
		fileName,
		task.Text,
	)
	err := generator.GenerateWav()
	task.ResultFile = &fileName
	return err
}
