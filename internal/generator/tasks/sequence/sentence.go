package sequence

import (
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

// SentenceType defines the type of sentence
type SentenceType int

const (
	// SingleSetence refers to sentence that is purely talked by one character
	SingleSentence SentenceType = iota
	// Chorus refers to a sentence told by multiple character
	Chorus
	// Rest refers to a period of rest between sentences
	Rest
)

// Sentence defines a single sentence component for the sequence task
type Sentence struct {
	TypeSentence           SentenceType        `json:"sentenceTtpe"`
	Text                   string              `json:"text"`
	RestTime               *int                `json:"restTime"`
	Speed                  *int                `json:"speed"`
	CharacterNamesIncluded []string            `json:"characterNameIncluded"`
	EffectList             []*edit.AudioEffect `json:"effectList"`
}
