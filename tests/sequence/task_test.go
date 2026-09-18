package sequence

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
)

func TestInsertSentence(t *testing.T) {
	t.Parallel()
	tmpDir := t.TempDir()
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
	task, err := sequence.NewSequenceTask(
		"test",
		nil,
		tmpDir,
	)
	task.AddSentence(
		*sequence.NewEmptySentence(
			114.514,
		),
	)
}
