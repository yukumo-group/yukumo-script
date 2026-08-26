package phontsmanager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yukumo-group/yukumo-script/pkg/utils/syncutils"
)

// PhontNameToFileName creates mapping of phont name and file
var PhontNameToFileName = syncutils.NewMap()

// GetAllPhonts gets all the phonts
func GetAllPhonts(phontsDir string) ([]os.DirEntry, error) {
	return os.ReadDir(phontsDir)
}

// InitializePhontNameToFileName initializes the key-value pair of phont name and phont file
func InitializePhontNameToFileName(phontsDir string) error {
	phonts, err := GetAllPhonts(phontsDir)
	if err != nil {
		return err
	}
	for _, phont := range phonts {
		name := phont.Name()
		suffix := filepath.Ext(name)
		if suffix == ".phont" {
			phontName := strings.TrimSuffix(
				name,
				suffix,
			)
			PhontNameToFileName.SetKV(phontName, name)
		}
	}
	return nil
}

// GetPhontFile gets the phont file and check if the file exists.
// Simply a reflection
func GetPhontFile(phontsDir string, phontName string) (string, error) {
	// Check if it exists in the PhontNameToFileName
	phontFile, exists := PhontNameToFileName.GetValue(
		phontName,
	)
	if !exists {
		return "", fmt.Errorf(
			"no phont file corresponds to %s",
			phontName,
		)
	}
	path := fmt.Sprintf(
		"%s/%s",
		phontsDir,
		phontFile,
	)
	// Check if it exists
	_, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	return path, nil
}
