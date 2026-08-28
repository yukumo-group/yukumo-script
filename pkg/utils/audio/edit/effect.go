package edit

import (
	"errors"
	"fmt"
)

// AudioEffect defines the effect on audios
type AudioEffect struct {
	ProcessType Process `json:"processType"`
	Data        any     `json:"data"`
}

// NewAudioEffect creats new audio effect
func NewAudioEffect(
	processType Process,
	data any,
) *AudioEffect {
	return &AudioEffect{
		ProcessType: processType,
		Data:        data,
	}
}

// UseEffect uses effect on audio
func (effect *AudioEffect) UseEffect(
	originalFilePath string,
	targetFilePath string,
) error {
	switch effect.ProcessType {
	case Resample:
		resampleData, ok := effect.Data.(ResampleData)
		if !ok {
			return errors.New("data type not supported")
		}
		err := UpdateResampledFile(
			originalFilePath,
			resampleData.TargetSampleRate,
			targetFilePath,
		)
		return err
	case NotAvailable:
		return errors.New(
			"this process is not available",
		)
	default:
		return fmt.Errorf(
			"%s process does not supported",
			effect.ProcessType,
		)
	}
}
