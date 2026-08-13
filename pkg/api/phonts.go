package api

import (
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
)

// InitPhontMap loads phont name → file mappings from PhontsDir.
func InitPhontMap() error {
	return phontsmanager.InitializePhontNameToFileName(filePathForProg.PhontsDir)
}

// ListPhonts returns available phont names.
func ListPhonts() []string {
	return phontsmanager.PhontNameToFileName.GetAllKeys()
}

// PhontPath resolves a phont name to its file path under PhontsDir.
func PhontPath(phontName string) (string, error) {
	return phontsmanager.GetPhontFile(filePathForProg.PhontsDir, phontName)
}

// IsPhontExists checks if phont exists
func IsPhontExists(phontName string) bool {
	_, exists := phontsmanager.PhontNameToFileName.GetValue(phontName)
	return exists
}
