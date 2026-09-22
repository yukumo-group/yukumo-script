package sequence

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

func TestWorkSpace(
	t *testing.T,
) {
	// Test load task
	t.Parallel()
	tmpDir := t.TempDir()
	testWorkSpace := sequence.NewWorkSpace()
	config, err := sequence.ReadRawConfig(
		"testdata/test1.yaml",
	)
	if err != nil {
		t.Fatal(err)
	}
	task, err := sequence.NewSequenceTask(
		"test",
		config,
		tmpDir,
	)
	if err != nil {
		t.Fatal(err)
	}
	id, err := testWorkSpace.LoadTaskToWorkSpace(
		task,
	)
	if err != nil {
		t.Error(err)
	}
	result := testWorkSpace.ShowAllTasks()
	_, exists := result[id]
	if !exists {
		t.Errorf(
			"%s id task that expected to exists does not exists",
			id,
		)
	}
	// Test add sentence
	sentenceEmpty := sequence.NewEmptySentence(
		10.0,
	)
	info := testWorkSpace.AddSentence(
		id,
		sentenceEmpty,
	)
	infoSentenceLength := len(info.AllSentences)
	if infoSentenceLength != 1 {
		t.Errorf(
			"expected length of all sentences to be %d, got %d",
			1,
			infoSentenceLength,
		)
	}
	// Test insert sentence
	sentenceEmpty2 := sequence.NewEmptySentence(
		11.0,
	)
	info = testWorkSpace.InsertSentence(
		id,
		sentenceEmpty2,
		0,
	)
	infoSentenceLength = len(info.AllSentences)
	if infoSentenceLength != 2 {
		t.Errorf(
			"expected length of all sentences to be %d, got %d",
			2,
			infoSentenceLength,
		)
	}
	sentenceSingleSentence := sequence.NewSingleSentence(
		"abc",
		"Remilia Scarlet",
		nil,
		[]*edit.AudioEffect{},
	)
	info = testWorkSpace.InsertSentence(
		id,
		sentenceSingleSentence,
		0,
	)
	infoSentenceLength = len(info.AllSentences)
	if infoSentenceLength != 3 {
		t.Errorf(
			"expected length of all sentences to be %d, got %d",
			3,
			infoSentenceLength,
		)
	}
	if info.AllSentences[1].TypeSentence != sequence.SingleSentence {
		t.Errorf(
			"expected type of sentence 1 to be %s, got %s",
			info.AllSentences[1].TypeSentence.ToString(),
			sequence.SingleSentence.ToString(),
		)
	}
}
