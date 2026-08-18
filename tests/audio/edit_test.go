package audio_test

import (
	"fmt"
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/empty"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

func TestAudioSlicing(
	t *testing.T,
) {
	testAudioInfo1 := &audio.Info{
		ChannelNumber: 1,
		SampleRate:    8000,
		Length:        1.5,
		Precision:     2,
	}
	testAudioInfo2 := &audio.Info{
		ChannelNumber: 1,
		SampleRate:    8000,
		Length:        1.5,
		Precision:     2,
	}
	task1, err := empty.NewEmptyTask(
		1.5,
		testAudioInfo1,
	)
	if err != nil {
		t.Error(err)
	}
	task2, err := empty.NewEmptyTask(
		1.5,
		testAudioInfo2,
	)
	if err != nil {
		t.Error(err)
	}
	tmpDir := t.TempDir()
	err = task1.Generate("", tmpDir)
	err = task2.Generate("", tmpDir)
	if task1.ResultFile == nil || task2.ResultFile == nil {
		t.Error("the audio failed to generate")
	}
	newWAVDir := fmt.Sprintf(
		"%s/%s",
		tmpDir,
		"test.wav",
	)
	err = edit.SpliceAudios(
		newWAVDir,
		[]string{
			*task1.ResultFile,
			*task2.ResultFile,
		},
	)
	if err != nil {
		t.Error(err)
	}
	newAudioInfo, err := audio.GetAudioInfo(
		newWAVDir,
	)
	if err != nil {
		t.Error(err)
	}
	if newAudioInfo.Length != 3.0 {
		t.Errorf(
			"Expected to have length %f, got %f",
			3.0,
			newAudioInfo.Length,
		)
	}
	updatedWAVDir := fmt.Sprintf(
		"%s/%s",
		tmpDir,
		"updated_test.wav",
	)
	err = edit.UpdateResampledFile(
		*task1.ResultFile,
		16000,
		updatedWAVDir,
	)
	if err != nil {
		t.Error(err)
	}
	newWAVDir = fmt.Sprintf(
		"%s/%s",
		tmpDir,
		"test2.wav",
	)
	err = edit.SpliceAudios(
		newWAVDir,
		[]string{
			*task1.ResultFile,
			updatedWAVDir,
		},
	)
	if err == nil {
		t.Error("expected to return error when slicing files that do not have same format")
	}
}
