package audio_test

import (
	"fmt"
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

func TestResampleEffect(t *testing.T) {
	t.Parallel()
	resampleData := edit.ResampleData{
		TargetSampleRate: edit.Hz16000,
	}
	newEffect := edit.NewAudioEffect(
		edit.Resample,
		resampleData,
	)
	tmpDir := t.TempDir()
	tmpFilePath := fmt.Sprintf(
		"%s/%s.wav",
		tmpDir,
		"test_resample_effect.wav",
	)
	err := newEffect.UseEffect(
		"testdata/example_aq_f1c.wav",
		tmpFilePath,
	)
	if err != nil {
		t.Error(err)
	}
}
