package sequence_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/sequence"
)

func TestReadConfig(t *testing.T) {
	t.Parallel()
	testConf, err := sequence.ReadRawConfig(
		"test_data/test1.yaml",
	)
	if err != nil {
		t.Error(err)
	}
	if testConf.Characters == nil {
		t.Error(
			"failed to read the configuration of characters",
		)
	}
	if len(*testConf.Characters) != 1 {
		t.Errorf(
			"expected length of characters in conf is %d, got %d",
			1,
			len(*testConf.Characters),
		)
	}
}

func TestTaskNameParsing(t *testing.T) {
	t.Parallel()
	testConf, err := sequence.ReadRawConfig(
		"test_data/test2.yaml",
	)
	if err != nil {
		t.Error(err)
	}
	if testConf.ConfigName == nil {
		t.Error("the config name is not successfully set")
	}
	if *testConf.ConfigName != "test2" {
		t.Errorf(
			"expected config name to be %s got %s",
			"test2",
			*testConf.ConfigName,
		)
	}
}

func TestToTaskConfig(t *testing.T) {
	t.Parallel()
	testConf, err := sequence.ReadRawConfig(
		"test_data/test3.yaml",
	)
	if err != nil {
		t.Error(err)
	}
	if testConf.Characters == nil {
		t.Error(
			"failed to read the configuration of characters",
		)
	}
	if len(*testConf.Characters) != 1 {
		t.Errorf(
			"expected length of characters in conf is %d, got %d",
			1,
			len(*testConf.Characters),
		)
	}
	taskConfig, err := testConf.ToTaskConfig()
	if err != nil {
		t.Error(err)
	}
	data := taskConfig.Characters.GetData()
	length := 0
	for range data {
		length++
	}
	if length != len(*testConf.Characters) {
		t.Errorf(
			"expected character data length to be %d, got %d",
			len(*testConf.Characters),
			len(data),
		)
	}
}
