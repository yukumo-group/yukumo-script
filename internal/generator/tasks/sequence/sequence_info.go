package sequence

import (
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

// SequenceInfo defines the info shows to the user when showing info
type SequenceInfo struct {
	Language     language.Language
	AllSentences []*Sentence
}
