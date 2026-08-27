package cmdinterface

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

// SaveFileAs saves the audio file to other directory
func SaveFileAs(
	title *color.Color,
	text *color.Color,
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
	formatList := api.GetAllPossibleAudioFormats()
	_, _ = title.Println("Here are the available formats:")
	for _, format := range formatList {
		_, _ = text.Println(format)
	}
	_, _ = title.Println("Input the format you want for the exported file:")
	var targetFormat string
	fmt.Scan(&targetFormat)
	find := false
	for _, format := range formatList {
		if format == targetFormat {
			find = true
		}
	}
	if !find {
		return fmt.Errorf(
			"format %s not supported",
			targetFormat,
		)
	}
	err = api.SaveAudioTo(
		originalFilePath,
		targetDirectory,
		targetFileName,
		api.ConvertStringToFormat(targetFormat),
	)
	if err != nil {
		return err
	}
	return nil
}
