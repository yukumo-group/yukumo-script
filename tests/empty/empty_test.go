package empty_test

import (
	"os"
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
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
		t.Context(),
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

func TestInterface(t *testing.T) {
	t.Parallel()
	var testTask interface{} = &empty.Task{}
	_, ok := testTask.(tasks.Task)
	if !ok {
		t.Error("Empty task cannot suit to the task interface")
	}
}
