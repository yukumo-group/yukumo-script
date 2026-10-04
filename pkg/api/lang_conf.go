package api

import (
	"fmt"

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

// GetDictFilePath gets the path for directory
func GetDictFilePath() string {
	return fmt.Sprintf(
		"%s/%s",
		filePathForProg.AssetsDir,
		filePathForProg.DictFilePath,
	)
}

// ConvertText converts text to kana
func ConvertText(
	originalText string,
	lang language.Language,
) (string, error) {
	return language.ConvertText(
		originalText,
		lang,
	)
}

// InitializeLanguageConfig initializes the polyphonic manager
func InitializeLanguageConfig() error {
	filePath := GetPolyphonicsFilePath()
	dictPath := GetDictFilePath()
	newManager, err := polyphonic.NewManagerFromFile(
		filePath,
		dictPath,
	)
	if err != nil {
		return err
	}
	language.PolyphonicsManager = newManager
	language.PolyphonicsManager.Initialize()
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
	data := language.PolyphonicsManager.GetData()
	_, exists := data[chinese]
	if exists {
		return fmt.Errorf(
			"pinyin for %s is already recorded",
			chinese,
		)
	}
	language.PolyphonicsManager.AddPolyphonic(
		chinese,
		pinyin,
	)
	err := language.PolyphonicsManager.SaveGSEDict()
	if err != nil {
		return err
	}
	return language.PolyphonicsManager.Save()
}
