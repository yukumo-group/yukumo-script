package sequence

import (
	"errors"
	"fmt"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/chorus"
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
	MixingConfig           *edit.MixingConfig  `json:"mixing_config"`
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
	text string,
	characterName string,
	speed *int,
	effectList []*edit.AudioEffect,
) *Sentence {
	return &Sentence{
		TypeSentence:           SingleSentence,
		EffectList:             effectList,
		CharacterNamesIncluded: []string{characterName},
		Speed:                  speed,
		Text:                   text,
	}
}

// NewChorus creates new chorus task
func NewChorus(
	text string,
	charactersNames []string,
	speed *int,
	effectList []*edit.AudioEffect,
	mixingConfig *edit.MixingConfig,
) *Sentence {
	return &Sentence{
		TypeSentence:           Chorus,
		CharacterNamesIncluded: charactersNames,
		Text:                   text,
		Speed:                  speed,
		MixingConfig:           mixingConfig,
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
	tmpDir string,
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
				"%s is not character or phont",
				id,
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
	case Chorus:
		if len(sentence.CharacterNamesIncluded) < 1 {
			return nil, errors.New(
				"the CharacterNamesIncluded cannot be empty ",
			)
		}
		phontList := []string{}
		characterNamesList := []string{}
		for _, characterName := range sentence.CharacterNamesIncluded {
			isCharacter, err := IsCharacterOrPhont(
				characterName,
				characterList,
			)
			if err != nil {
				return nil, err
			}
			switch isCharacter {
			case 1:
				phontList = append(phontList, characterName)
			case -1:
				characterNamesList = append(characterNamesList, characterName)
			default:
				return nil, fmt.Errorf(
					"%s is not character or phont",
					characterName,
				)
			}
		}
		return chorus.NewChorusTask(
			sentence.Text,
			&phontList,
			&characterNamesList,
			speedUsed,
			taskName,
			taskLanguage,
			sentence.MixingConfig,
			characterList,
			tmpDir,
		)
	default:
		return nil, fmt.Errorf(
			"%d sentence type not supported",
			sentence.TypeSentence,
		)
	}
}
