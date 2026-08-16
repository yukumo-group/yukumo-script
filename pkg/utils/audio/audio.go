//go:build !noaudio

package audio

import (
	"os"
	"time"

	"github.com/unitoftime/beep"
	"github.com/unitoftime/beep/speaker"
	"github.com/unitoftime/beep/wav"
)

// PlayWAV plays the wav file
func PlayWAV(fileName string) (*string, error) {
	resultChan := make(chan bool)
	file, errFile := os.Open(fileName)
	if errFile != nil {
		return &fileName, errFile
	}
	streamer, format, err := wav.Decode(file)
	if err != nil {
		return &fileName, err
	}
	defer func() { _ = streamer.Close() }()
	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return &fileName, err
	}
	speaker.Play(beep.Seq(
		streamer,
		beep.Callback(
			func() {
				resultChan <- true
			},
		),
	))
	<-resultChan
	time.Sleep(500 * time.Millisecond)
	return &fileName, nil
}
