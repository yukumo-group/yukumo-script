package api

import (
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils"
)

// ListPhonts returns available phont names.
func ListPhonts() []string {
	return phontsmanager.PhontNameToFileName.GetAllKeys()
}

// PhontPath resolves a phont name to its file path under PhontsDir.
func PhontPath(phontName string) (string, error) {
	return phontsmanager.GetPhontFile(utils.PhontsDir, phontName)
}
