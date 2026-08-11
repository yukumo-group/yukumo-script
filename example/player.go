package example

import (
	"fmt"

	"github.com/yukumo-group/yukumo-script/utils/audio"
)

// GetAllExampleFont gets the font name of all available phonts
func GetAllExampleFont() []string {
	return examplesMap.GetAllKeys()
}

// PlayExample plays example of a phont
func PlayExample(phontName string) (*string, error) {
	fileName, exists := examplesMap.GetValue(phontName)
	if !exists {
		return nil, fmt.Errorf(
			"Example for %s does not exists",
			phontName,
		)
	}
	return audio.PlayWAV(fileName)
}
