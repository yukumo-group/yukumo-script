//go:build !clib

package api

import (
	"github.com/yukumo-group/yukumo-script/internal/example"
)

func PlayExample(
	phontName string,
) (*string, error) {
	return example.PlayExample(phontName)
}
