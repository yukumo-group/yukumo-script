package chorus

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

// GenerateWavName generates name for the result wav file
func (task *Task) GenerateWavFileName(
	targetDir string,
) string {
	tmpID := uuid.NewString()
	return fmt.Sprintf(
		"%s/%s_%s_%d_%s.wav",
		targetDir,
		task.TaskName,
		task.ID,
		task.EditTime.UnixNano(),
		tmpID,
	)
}

// GenerateTempWavFile genereats name for temporary wav file for Effects
func (task *Task) GenerateTempWavFile(
	tempFileDir string,
	processMethod edit.Process,
) string {
	tmpID := uuid.NewString()
	return fmt.Sprintf(
		"%s/%s_%s_%d_%s_%s.wav",
		tempFileDir,
		task.TaskName,
		task.ID,
		task.EditTime.UnixNano(),
		processMethod,
		tmpID,
	)
}
