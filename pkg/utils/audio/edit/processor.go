package edit

import (
	"bytes"
	"os"

	"github.com/go-audio/wav"
	"github.com/jonchammer/audio-io/wave"
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

// UpdateResampledFile updates a file
func UpdateResampledFile(
	wavFileName string,
	targetSampleRate int,
) error {
	// Read wav data
	wavBytes, errReadWavFile := os.ReadFile(wavFileName)
	if errReadWavFile != nil {
		return errReadWavFile
	}
	wavReader := bytes.NewReader(wavBytes)
	wavDecoder := wav.NewDecoder(wavReader)
	wavBuffer, errDecode := wavDecoder.FullPCMBuffer()
	if errDecode != nil {
		return errDecode
	}
	// Convert audio data to int16
	decodedData := make([]int16, len(wavBuffer.Data))
	for i, data := range wavBuffer.Data {
		decodedData[i] = int16(data)
	}
	resampledDecodedData, errResample := ResampleWAV(
		decodedData,
		wavBuffer.Format.NumChannels,
		wavBuffer.Format.SampleRate,
		targetSampleRate,
	)
	if errResample != nil {
		return errResample
	}
	// Open file
	file, errOpenFile := os.Open(wavFileName)
	if errOpenFile != nil {
		return errOpenFile
	}
	// Write
	writer, errCreateWriter := wave.NewWriter(
		file,
		wave.SampleTypeInt16,
		uint32(targetSampleRate),
		wave.WithChannelCount(
			uint16(wavBuffer.Format.NumChannels),
		),
	)
	if errCreateWriter != nil {
		return errCreateWriter
	}
	flushWriter := func() {
		_ = writer.Flush()
	}
	defer flushWriter()
	errWrite := writer.WriteInt16(resampledDecodedData)
	return errWrite
}

// UpdateChannelNumberTo2 updates the data to two channels.
// This is just a temporary solution for the problem mentioned in https://github.com/braheezy/shine-mp3/issues/11.
func UpdateChannelNumberTo2(
	data []int16,
) []int16 {
	newData := make([]int16, len(data)*2)
	for i, sample := range data {
		newData[i*2] = sample
		newData[i*2+1] = sample
	}
	return newData
}
