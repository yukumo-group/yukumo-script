package api_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/api"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

func TestLanguageConfig(
	t *testing.T,
) {
	t.Parallel()
	err := api.Init()
	if err != nil {
		t.Error(err)
	}
	err = api.AddPolyphonic(
		"都市",
		"du shi",
	)
	if err != nil {
		t.Error(err)
	}
	err = api.AddPolyphonic(
		"银行",
		"yin hang",
	)
	if err != nil {
		t.Error(err)
	}
	result, err := api.ConvertText(
		"都市银行",
		language.Chinese,
	)
	if err != nil {
		t.Error(err)
	}
	const expecteResult string = "トゥーシーインハン"
	if result != expecteResult {
		t.Errorf(
			"expected %s, got %s",
			expecteResult,
			result,
		)
	}
}
