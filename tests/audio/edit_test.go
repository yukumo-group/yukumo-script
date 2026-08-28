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
	t.Parallel()
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
	err = task1.Generate(t.Context(), "", tmpDir)
	if err != nil {
		t.Error(err)
	}
	err = task2.Generate(t.Context(), "", tmpDir)
	if err != nil {
		t.Error(err)
	}
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

func TestMixingMethodToInt(
	t *testing.T,
) {
	t.Parallel()
	if edit.ByAverage.ToInt() != 1 {
		t.Errorf(
			"Expected %d, got %d",
			1,
			edit.ByAverage.ToInt(),
		)
	}
	if edit.ToMixingMethod(1) != edit.ByAverage {
		t.Errorf(
			"Expected %d, got %d",
			edit.ByAverage,
			edit.ToMixingMethod(1),
		)
	}
	if edit.ToMixingMethod(4) != edit.ByAverage {
		t.Errorf(
			"Expected %d, got %d",
			edit.ByAverage,
			edit.ToMixingMethod(4),
		)
	}
}

func TestGetBaseVolume(
	t *testing.T,
) {
	t.Parallel()
	base, volume := edit.GetBaseVolume(
		0.5,
	)
	if base != 2.0 {
		t.Errorf(
			"Expected %f, got %f",
			2.0,
			base,
		)
	}
	if volume != -1.0 {
		t.Errorf(
			"Expected %f, got %f",
			-1.0,
			volume,
		)
	}
	_, volume = edit.GetBaseVolume(
		-0.5,
	)
	t.Log(volume)
	if volume >= 0.0 {
		t.Error(
			"Expected volum to be negative",
		)
	}
}

func TestMixAudio(
	t *testing.T,
) {
	t.Parallel()
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
	err = task1.Generate(t.Context(), "", tmpDir)
	if err != nil {
		t.Error(err)
	}
	err = task2.Generate(t.Context(), "", tmpDir)
	if err != nil {
		t.Error(err)
	}
	if task1.ResultFile == nil || task2.ResultFile == nil {
		t.Error("the audio failed to generate")
	}
	newWAVDir := fmt.Sprintf(
		"%s/%s",
		tmpDir,
		"testMixAudio.wav",
	)
	testMixingConfig := edit.NewMixingMethod(
		edit.ByDefault,
		nil,
	)
	err = edit.MixAudios(
		[]string{
			*task1.ResultFile,
			*task2.ResultFile,
		},
		testMixingConfig,
		newWAVDir,
	)
	if err != nil {
		t.Error(err)
	}
	testMixingConfig = edit.NewMixingMethod(
		edit.ByAverage,
		nil,
	)
	err = edit.MixAudios(
		[]string{
			*task1.ResultFile,
			*task2.ResultFile,
		},
		testMixingConfig,
		newWAVDir,
	)
	if err != nil {
		t.Error(err)
	}
	testMixingConfig = edit.NewMixingMethod(
		edit.ByCustom,
		nil,
	)
	err = edit.MixAudios(
		[]string{
			*task1.ResultFile,
			*task2.ResultFile,
		},
		testMixingConfig,
		newWAVDir,
	)
	if err != nil {
		t.Error(err)
	}
	testMixingConfig = edit.NewMixingMethod(
		edit.ByCustom,
		&[]float64{0.5, 0.5},
	)
	err = edit.MixAudios(
		[]string{
			*task1.ResultFile,
			*task2.ResultFile,
		},
		testMixingConfig,
		newWAVDir,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestIntegratedEffectMethod(t *testing.T) {
	t.Parallel()
	effect := edit.NewAudioEffect(
		edit.Resample,
		edit.ResampleData{
			TargetSampleRate: 16000,
		},
	)
	targetFilePath := fmt.Sprintf(
		"%s/%s",
		t.TempDir(),
		"test.wav",
	)
	err := effect.UseEffect(
		"testdata/example_aq_f1c.wav",
		targetFilePath,
	)
	if err != nil {
		t.Error(err)
	}
	info, err := audio.GetAudioInfo(targetFilePath)
	if err != nil {
		t.Error(err)
	}
	t.Log(info.Length)
}
