package audio

import (
	"os"

	"github.com/faiface/beep/wav"
)

// Info reveals the information about the audio
type Info struct {
	ChannelNumbers int
	SampleRate     int
	Length         float64 // length in seconds
	BitDepth       int
}

// GetAudioInfo gets the information of an audio
func GetAudioInfo(
	fileName string,
) (*Info, error) {
	file, errFile := os.Open(fileName)
	if errFile != nil {
		return nil, errFile
	}
	streamer, format, errDecode := wav.Decode(file)
	if errDecode != nil {
		return nil, errDecode
	}
	closeStreamer := func() {
		_ = streamer.Close()
	}
	defer closeStreamer()
	totalSamples := streamer.Len()
	sampleRate := format.SampleRate
	length := float64(totalSamples) / float64(sampleRate)
	return &Info{
		ChannelNumbers: format.NumChannels,
		SampleRate:     int(sampleRate),
		Length:         length,
		BitDepth:       format.Precision,
	}, nil
}
