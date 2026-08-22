package singlesentence

import (
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
)

// PhontFile gets target phont file name
func PhontFile(
	phontsDir string,
	phontName string,
) (string, error) {
	return phontsmanager.GetPhontFile(
		phontsDir,
		phontName,
	)
}
