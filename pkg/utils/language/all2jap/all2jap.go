package all2jap

import (
	"regexp"

	kanatrans "github.com/Luigi-Pizzolito/English2KanaTransliteration"
	"github.com/yukumo-group/Chinese2KanaConverter/pkg/converter"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language/convertnums"
)

// AllToKana converts English and Japanese characters to Kana
func AllToKana(text string) string {
	allToKana := kanatrans.NewAllToKana(true)
	convertResult := allToKana.Convert(text)
	re := regexp.MustCompile(`[^\p{Katakana}]+`)
	result := re.ReplaceAllString(convertResult, "")
	return result
}

// EngToKana converts English to Kana
func EngToKana(text string) string {
	numResult := convertnums.ConvertNumToEnglish(text)
	return AllToKana(numResult)
}

// JPToKana converts japanese to Kana
func JPToKana(text string) string {
	numResult := convertnums.ConverNumToJP(text)
	return AllToKana(numResult)
}

// CnToKana converts chinese to kana with only words
func CnToKana(
	text string,
) (string, error) {
	replacedText := convertnums.ConverNumToCN(
		text,
	)
	convertResult, err := converter.SingleChinesePieceToKana(
		replacedText,
		true,
	)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`[^\p{Katakana}]+`)
	result := re.ReplaceAllString(convertResult, "")
	return result, nil
}
