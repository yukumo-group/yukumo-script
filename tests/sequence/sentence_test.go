package sequence

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

func TestToTask(t *testing.T) {
	t.Parallel()
	tmpCharacters := characters.NewCharacters()
	err := tmpCharacters.AddCharacter(
		characters.NewCharacter(
			"Remilia Scarlet",
			"aq_f1c",
			"",
			nil,
		),
	)
	if err != nil {
		t.Error(err)
	}
	tmpDir := t.TempDir()
	sentenceEmpty := sequence.NewEmptySentence(
		10.0,
	)
	_, err = sentenceEmpty.ToTask(
		tmpDir,
		"testEmpty",
		audio.DefaultAudioInfo,
		100,
		tmpCharacters,
		language.English,
	)
	if err != nil {
		t.Error(err)
	}
	sentenceSingleSentence := sequence.NewSingleSentence(
		"abc",
		"Remilia Scarlet",
		nil,
		[]*edit.AudioEffect{},
	)
	_, err = sentenceSingleSentence.ToTask(
		tmpDir,
		"testEmpty",
		audio.DefaultAudioInfo,
		100,
		tmpCharacters,
		language.English,
	)
	if err != nil {
		t.Error(err)
	}
	newMixingConfig := edit.NewMixingMethod(
		edit.ByAverage,
		nil,
	)
	sentenceChorus := sequence.NewChorus(
		"11",
		[]string{"Remilia Scarlet"},
		nil,
		[]*edit.AudioEffect{},
		newMixingConfig,
	)
	_, err = sentenceChorus.ToTask(
		tmpDir,
		"testEmpty",
		audio.DefaultAudioInfo,
		100,
		tmpCharacters,
		language.English,
	)
	if err != nil {
		t.Error(err)
	}
}
