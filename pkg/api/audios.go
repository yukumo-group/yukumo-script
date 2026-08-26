package api

import (
	"fmt"

	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
	"github.com/yukumo-group/yukumo-script/pkg/utils/osoperation"
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

// SaveAudioTo copies file to target directory with certain name
func SaveAudioTo(
	originalFilePath string,
	targetDirectory string,
	targetFileName string,
	format audio.Format,
) error {
	return osoperation.CopyFile(
		originalFilePath,
		targetDirectory,
		targetFileName,
		format.ToString(),
	)
}
