package audio

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
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
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
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
