package convertnums

import (
	"strings"
)

// enList defines the transition from numbers to english
var enList = map[string]string{
	"1": " one ",
	"2": " two ",
	"3": " three ",
	"4": " four ",
	"5": " five ",
	"6": " six ",
	"7": " seven ",
	"8": " eight ",
	"9": " nine ",
	"0": " zero ",
}

// jpList defines the transition from numbers to Kana
var jpList = map[string]string{
	"1": " イチ ",
	"2": " ニ ",
	"3": " サン ",
	"4": " シ ",
	"5": " ゴ ",
	"6": " ロク ",
	"7": " シチ ",
	"8": " ハチ ",
	"9": " ク ",
	"0": " ゼロ ",
}

// cnList defines the transition from numbers to Chinese
var cnList = map[string]string{
	"1": " 一 ",
	"2": " 二 ",
	"3": " 三 ",
	"4": " 四 ",
	"5": " 五 ",
	"6": " 六 ",
	"7": " 七 ",
	"8": " 八 ",
	"9": " 九 ",
	"0": " 零 ",
}

// ConvertNumToEnglish converts numbers to English vocabularies
func ConvertNumToEnglish(text string) string {
	result := text
	for num, vocab := range enList {
		result = strings.ReplaceAll(result, num, vocab)
	}
	return result
}

// ConverNumToJP converts numbers to Kana
func ConverNumToJP(text string) string {
	result := text
	for num, vocab := range jpList {
		result = strings.ReplaceAll(result, num, vocab)
	}
	return result
}

// ConverNumToCN converts numbers to Chinese
func ConverNumToCN(text string) string {
	result := text
	for num, vocab := range cnList {
		result = strings.ReplaceAll(result, num, vocab)
	}
	return result
}
