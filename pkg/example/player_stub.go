//go:build !clib && noaudio

package example

import "fmt"

// PlayExample is a stub when built with -tags noaudio (cross linux/macOS CLI).
func PlayExample(phontName string) (*string, error) {
	return nil, fmt.Errorf(
		"audio playback not included in this build (phont %s)",
		phontName,
	)
}
