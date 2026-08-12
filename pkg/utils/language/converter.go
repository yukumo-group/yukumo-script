package language

import (
	"errors"

	"github.com/yukumo-group/yukumo-script/pkg/utils/language/all2jap"
)

// ConvertText converts text to certain language
func ConvertText(
	text string,
	lang Language,
) (string, error) {
	switch lang {
	case English:
		return all2jap.EngToKana(text), nil
	case Japanese:
		return all2jap.JPToKana(text), nil
	default:
		return "", errors.New("Language not supported")
	}
}
