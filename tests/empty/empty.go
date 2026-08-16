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
		ChannelNumbers: 1,
		SampleRate:     8000,
		Length:         1.5,
		Precision:      2,
	}
	newTask := empty.NewEmptyTask(
		testAudioInfo,
	)
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
