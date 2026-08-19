package api

import (
	"fmt"

	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

// SpliceAudios splices the audios.
// Do not need to add wav after file name.
func SpliceAudios(
	audioPathes []string,
	targetDirectory string,
	targetFileName string,
) (string, error) {
	resultFileName := fmt.Sprintf(
		"%s/%s",
		targetDirectory,
		targetFileName,
	)
	err := edit.SpliceAudios(
		resultFileName,
		audioPathes,
	)
	if err != nil {
		return resultFileName, err
	}
	return resultFileName, nil
}
