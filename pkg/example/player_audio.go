//go:build !clib && !noaudio

package example

import (
	"fmt"

	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
)

// PlayExample plays example of a phont
func PlayExample(phontName string) (*string, error) {
	fileName, exists := examplesMap.GetValue(phontName)
	if !exists {
		return nil, fmt.Errorf(
			"example for %s does not exist",
			phontName,
		)
	}
	return audio.PlayWAV(fileName)
}
