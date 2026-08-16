package empty_test

import (
	"os"
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/empty"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
)

func TestCreateEmptyTask(t *testing.T) {
	t.Parallel()
	testAudioInfo := &audio.Info{
		ChannelNumber: 1,
		SampleRate:    8000,
		Length:        1.5,
		Precision:     2,
	}
	newTask, errNewEmptyTask := empty.NewEmptyTask(
		1.5,
		testAudioInfo,
	)
	if errNewEmptyTask != nil {
		t.Error(errNewEmptyTask)
	}
	tmpDir := os.TempDir()
	errGenerate := newTask.Generate(
		"",
		tmpDir,
	)
	if errGenerate != nil {
		t.Error(errGenerate)
	}
	if newTask.ResultFile == nil {
		t.Error("Result file should not be nil after generation")
	}
}
