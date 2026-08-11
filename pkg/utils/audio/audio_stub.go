//go:build noaudio

package audio

import "fmt"

// PlayWAV is omitted in cross builds (-tags noaudio).
func PlayWAV(fileName string) (*string, error) {
	return &fileName, fmt.Errorf("audio playback not included in this build")
}
