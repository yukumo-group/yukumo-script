package cmdinterface

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

// SaveFileAs saves the audio file to other directory
func SaveFileAs(
	title *color.Color,
	originalFilePath string,
) error {
	var targetDirectory string
	_, _ = title.Println("Input the target directory you want to save")
	_, err := fmt.Scan(&targetDirectory)
	if err != nil {
		return err
	}
	var targetFileName string
	_, _ = title.Println("Input the file name of the exported file")
	_, err = fmt.Scan(&targetFileName)
	if err != nil {
		return err
	}
	err = api.SaveAudioTo(
		originalFilePath,
		targetDirectory,
		targetFileName,
		"wav",
	)
	if err != nil {
		return err
	}
	return nil
}
