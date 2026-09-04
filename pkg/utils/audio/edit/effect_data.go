package edit

import (
	"fmt"
)

// SupportedSampleRate defines all the
type SupportedSampleRate int

const (
	// Hz8000 8000Hz
	Hz8000 SupportedSampleRate = 8000
	// Hz11025 11025Hz
	Hz11025 SupportedSampleRate = 11025
	// Hz16000 16000Hz
	Hz16000 SupportedSampleRate = 16000
	// Hz22050 22050Hz
	Hz22050 SupportedSampleRate = 22050
	// Hz32000 32000Hz
	Hz32000 SupportedSampleRate = 32000
	// Hz44100 44100Hz
	Hz44100 SupportedSampleRate = 44100
	// Hz48000 48000Hz
	Hz48000 SupportedSampleRate = 48000
)

// ToInt converts sample rate to integer
func (sampleRate SupportedSampleRate) ToInt() (int, error) {
	switch sampleRate {
	case Hz8000:
		return 8000, nil
	case Hz11025:
		return 11025, nil
	case Hz16000:
		return 16000, nil
	case Hz32000:
		return 32000, nil
	case Hz44100:
		return 44100, nil
	case Hz48000:
		return 48000, nil
	default:
		return 0, fmt.Errorf(
			"%d sample rate is not supported",
			sampleRate,
		)
	}
}

// GetAllSupportedSampleRates gets all the data supported
func GetAllSupportedSampleRates() []SupportedSampleRate {
	return []SupportedSampleRate{
		Hz8000,
		Hz11025,
		Hz16000,
		Hz22050,
		Hz32000,
		Hz44100,
		Hz48000,
	}
}

// ResampleData contains the target sample rate to resample a file
type ResampleData struct {
	TargetSampleRate SupportedSampleRate
}
