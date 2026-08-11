package audio

import (
	"github.com/zeozeozeo/gomplerate"
)

// ResampleWAV resamples the wav data
func ResampleWAV(
	data []int16,
	channelAmount int,
	originalSampleRate int,
	targetSampleRate int,
) ([]int16, error) {
	resampler, errCreateResampler := gomplerate.NewResampler(
		channelAmount,
		originalSampleRate,
		targetSampleRate,
	)
	if errCreateResampler != nil {
		return nil, errCreateResampler
	}
	resampledData := resampler.ResampleInt16(data)
	return resampledData, nil
}
