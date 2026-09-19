package sequence

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

func TestInsertSentence(t *testing.T) {
	t.Parallel()
	testSentence := sequence.NewSingleSentence(
		"abc",
		"Remilia Scarlet",
		nil,
		[]*edit.AudioEffect{},
	)
	tmpDir := t.TempDir()
	config, err := sequence.ReadRawConfig(
		"testdata/test1.yaml",
	)
	if err != nil {
		t.Error(err)
	}
	task, err := sequence.NewSequenceTask(
		"test",
		config,
		tmpDir,
	)
	if err != nil {
		t.Error(err)
	}
	task.AddSentence(
		sequence.NewEmptySentence(
			114.514,
		),
	)
	task.AddSentence(
		sequence.NewEmptySentence(
			114.514,
		),
	)
	task.InsertSentence(
		testSentence,
		0,
	)
	expectedResult := []sequence.SentenceType{
		sequence.Empty,
		sequence.SingleSentence,
		sequence.Empty,
	}
	info := task.ToPreviewInfo()
	if len(info.AllSentences) != 3 {
		t.Errorf(
			"expected length to be %d, got %d",
			3,
			len(info.AllSentences),
		)
	}
	for i, sentence := range info.AllSentences {
		if i > len(expectedResult)-1 {
			t.Errorf(
				"%d idx not exists in expected result",
				i,
			)
			if sentence.TypeSentence != expectedResult[i] {
				t.Errorf(
					"expected %s, got %s for %d",
					expectedResult[i].ToString(),
					sentence.TypeSentence.ToString(),
					i,
				)
			}
		}
	}
	task.InsertSentence(
		testSentence,
		2,
	)
}
