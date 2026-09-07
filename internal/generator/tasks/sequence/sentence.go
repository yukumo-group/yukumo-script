package sequence

import (
	"errors"
	"fmt"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/empty"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

// SentenceType defines the type of sentence
type SentenceType int

const (
	// SingleSetence refers to sentence that is purely talked by one character
	SingleSentence SentenceType = iota
	// Chorus refers to a sentence told by multiple character
	Chorus
	// Empty refers to a period of empty between sentences
	Empty
)

// Sentence defines a single sentence component for the sequence task
type Sentence struct {
	TypeSentence           SentenceType        `json:"sentenceTtpe"`
	Text                   string              `json:"text"`
	RestTime               *float64            `json:"restTime"`
	Speed                  *int                `json:"speed"`
	CharacterNamesIncluded []string            `json:"characterNameIncluded"`
	EffectList             []*edit.AudioEffect `json:"effectList"`
}

// NewEmptySentence creates new sentence with type empty
func NewEmptySentence(
	restTime float64,
) *Sentence {
	return &Sentence{
		TypeSentence: Empty,
		RestTime:     &restTime,
	}
}

// NewSingleSentence creates sentence with a single speaker
func NewSingleSentence(
	characterName string,
	effectList []*edit.AudioEffect,
) *Sentence {
	return &Sentence{
		TypeSentence:           SingleSentence,
		EffectList:             effectList,
		CharacterNamesIncluded: []string{characterName},
	}
}

// IsCharacterOrPhont decides whether it is a character or phont.
// -1: Character
// 0: Not character and not phont
// 1: Phont
func IsCharacterOrPhont(
	id string,
	characterList *characters.Characters,
) (int, error) {
	characterMap := characterList.GetData()
	_, exists := characterMap[id]
	if exists {
		return -1, nil
	}
	_, exists = phontsmanager.PhontNameToFileName.GetValue(
		id,
	)
	if exists {
		return 1, nil
	}
	return 0, fmt.Errorf(
		"the program failed to find %s in both character list and phont list",
		id,
	)
}

// ToTask converts Sentence to task
func (sentence *Sentence) ToTask(
	taskName string,
	audioInfo *audio.Info,
	speed int,
	characterList *characters.Characters,
	taskLanguage language.Language,
) (tasks.Task, error) {
	var speedUsed int
	if sentence.Speed == nil {
		speedUsed = speed
	} else {
		speedUsed = *sentence.Speed
	}
	switch sentence.TypeSentence {
	case SingleSentence:
		// Get character name or phont name
		if len(sentence.CharacterNamesIncluded) < 1 {
			return nil, errors.New(
				"the CharacterNamesIncluded cannot be empty ",
			)
		}
		id := sentence.CharacterNamesIncluded[0]
		isCharacter, err := IsCharacterOrPhont(
			id,
			characterList,
		)
		if err != nil {
			return nil, err
		}
		var characterName *string = nil
		var phontName *string = nil
		// Check if it is character name or phont name
		switch isCharacter {
		case -1:
			characterName = &id
			phontName = nil
		case 1:
			characterName = nil
			phontName = &id
		default:
			return nil, fmt.Errorf(
				"%d is not character or phont",
				isCharacter,
			)
		}
		return singlesentence.NewSingleSentenceTask(
			sentence.Text,
			characterName,
			phontName,
			speedUsed,
			taskName,
			taskLanguage,
			characterList,
		)
	case Empty:
		return empty.NewEmptyTask(
			*sentence.RestTime,
			audioInfo,
		)
	default:
		return nil, fmt.Errorf(
			"%d sentence type not supported",
			sentence.TypeSentence,
		)
	}
}
