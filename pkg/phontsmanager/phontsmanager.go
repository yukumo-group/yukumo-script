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
		extension := filepath.Ext(name)
		if extension == ".phont" {
			phontName := strings.TrimSuffix(
				name,
				extension,
			)
			PhontNameToFileName.SetKV(phontName, name)
		}
	}
	return nil
}

// GetPhontFile gets the phont file and check if the file exists
func GetPhontFile(phontsDir string, phontName string) (string, error) {
	// Check if it exists in the PhontNameToFileName
	phontFile, exists := PhontNameToFileName.GetValue(
		phontName,
	)
	if !exists {
		return "", fmt.Errorf(
			"No phont file correspond to %s",
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
