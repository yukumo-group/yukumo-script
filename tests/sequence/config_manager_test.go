package sequence_test

import (
	"fmt"
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
)

func TestConfigManager(t *testing.T) {
	t.Parallel()
	tmpDir := t.TempDir()
	useDefault := true
	newManager := sequence.NewConfigManager()
	newManager.SetConfigFilePath(
		tmpDir,
		"test.json",
	)
	newConfig := sequence.NewRawConfig(
		"114",
		1,
		&useDefault,
		nil,
		nil,
	)
	err := newManager.AddRawConfig(
		newConfig,
		tmpDir,
	)
	if err != nil {
		t.Error(err)
	}
	err = newManager.AddFile(
		"testdata/test1.yaml",
		tmpDir,
	)
	if err != nil {
		t.Error(err)
	}
	err = newManager.SaveFile()
	if err != nil {
		t.Error(err)
	}
	_, err = sequence.NewConfigManagerFromFile(
		fmt.Sprintf(
			"%s/%s",
			tmpDir,
			"test.json",
		),
	)
	if err != nil {
		t.Error(err)
	}
}
