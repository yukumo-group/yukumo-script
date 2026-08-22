package cmdinterface

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

func ShowPhonts(
	title *color.Color,
	text *color.Color,
) (string, error) {
	// Print info
	_, _ = title.Println("Here are the available phonts:")
	for _, phontName := range api.ListPhonts() {
		_, _ = text.Println(phontName)
	}
	_, _ = title.Println("Input the name of the phont you want to use to generate audio:")
	// Input
	var phontName string
	_, errInput := fmt.Scan(&phontName)
	if errInput != nil {
		return "", errInput
	}
	for _, iPhontName := range api.ListPhonts() {
		if iPhontName != phontName {
			return "", fmt.Errorf(
				"%s phont name does not exists",
				phontName,
			)
		}
	}
	return phontName, nil
}
