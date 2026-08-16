package audio_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
)

func TestGetAudioInfo(t *testing.T) {
	data, errGetAudioInfo := audio.GetAudioInfo(
		"testdata/example_aq_f1c.wav",
	)
	if errGetAudioInfo != nil {
		t.Error(errGetAudioInfo)
	}
	if data.Precision != 2 {
		t.Errorf(
			"Expected precision to be %d, got %d",
			2,
			data.Precision,
		)
	}
	if data.Length != 1.087000 {
		t.Errorf(
			"Expected audio length to be %f, got %f",
			1.087000,
			data.Length,
		)
	}
	if data.ChannelNumber != 1 {
		t.Errorf(
			"Expected channel number to be %d, got %d",
			1,
			data.ChannelNumber,
		)
	}
	if data.SampleRate != 8000 {
		t.Errorf(
			"Expected sample rate to be %d, got %d",
			8000,
			data.SampleRate,
		)
	}
}
