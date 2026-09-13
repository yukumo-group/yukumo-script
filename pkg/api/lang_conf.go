package api

import (
	"fmt"
	"os"

	"github.com/yukumo-group/Chinese2KanaConverter/pkg/polyphonic"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

// GetPolyphonicsFilePath gets the path for polyphonics file
func GetPolyphonicsFilePath() string {
	return fmt.Sprintf(
		"%s/%s",
		filePathForProg.DataDir,
		filePathForProg.PolyphonicsManagerFile,
	)
}

// InitializeLanguageConfig initializes the polyphonic manager
func InitializeLanguageConfig() error {
	filePath := GetPolyphonicsFilePath()
	// Check if file exists
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			language.PolyphonicsManager.SetTargetFile(
				filePath,
			)
			return nil
		} else {
			return err
		}
	}
	newManager, err := polyphonic.NewManagerFromFile(
		filePath,
	)
	if err != nil {
		return err
	}
	language.PolyphonicsManager = newManager
	return nil
}

// GetAllPolyphonics gets all the polyphonics stored
func GetAllPolyphonics() map[string]string {
	return language.PolyphonicsManager.GetData()
}

// AddPolyphonic adss new polyphonic
func AddPolyphonic(
	chinese string,
	pinyin string,
) error {
	language.PolyphonicsManager.AddPolyphonic(
		chinese,
		pinyin,
	)
	return language.PolyphonicsManager.Save()
}
