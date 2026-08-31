package sequence_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
)

func TestReadConfig(t *testing.T) {
	testConf, err := sequence.ReadRawConfig(
		"test_data/test1.yaml",
	)
	if err != nil {
		t.Error(err)
	}
	if len(testConf.Characters) != 1 {
		t.Errorf(
			"expected length of characters in conf is %d, got %d",
			1,
			len(testConf.Characters),
		)
	}
}
