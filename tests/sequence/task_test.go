package sequence

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
)

func TestInsertSentence(t *testing.T) {
	t.Parallel()
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
}
