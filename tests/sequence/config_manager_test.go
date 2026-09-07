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
	reloadedManager, err := sequence.NewConfigManagerFromFile(
		fmt.Sprintf(
			"%s/%s",
			tmpDir,
			"test.json",
		),
	)
	if err != nil {
		t.Error(err)
	}
	allConfigs := reloadedManager.GetAllConfigNames()
	finded114 := false
	findedtest1 := false
	for _, configName := range allConfigs {
		if configName == "114" {
			finded114 = true
		}
		if configName == "test1" {
			findedtest1 = true
		}
	}
	if !finded114 {
		t.Error("114 config not found")
	}
	if !findedtest1 {
		t.Error("test1 config not found")
	}
	configtest1, err := reloadedManager.GetConfig(
		"test1",
	)
	if err != nil {
		t.Error(err)
	}
	if configtest1.Characters == nil {
		t.Error("failed to read character list")
	}
	length := len(*configtest1.Characters)
	if length != 2 {
		t.Errorf(
			"expected length of character list to be %d, got %d",
			2,
			length,
		)
	}
}
