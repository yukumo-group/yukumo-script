package api

import (
	"fmt"

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

// GetAllPolyphonics gets all the polyphonics stored
func GetAllPolyphonics() map[string]string {
	return language.PolyphonicsManager.GetData()
}

// AddPolyphonic adss new polyphonic
func AddPolyphonic(
	chinese string,
	pinyin string,
) {
	language.PolyphonicsManager.AddPolyphonic(
		chinese,
		pinyin,
	)
}
