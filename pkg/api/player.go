//go:build !clib

package api

import (
	"github.com/yukumo-group/yukumo-script/internal/example"
)

// PlayExample plays the example file
func PlayExample(
	phontName string,
) (*string, error) {
	return example.PlayExample(phontName)
}
