package api_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/api"
)

func TestLoadMethods(t *testing.T) {
	t.Parallel()
	result := api.ToLoadConfigMethod(
		0,
	)
	if result != api.FromFile {
		t.Errorf(
			"expected %s, got %s",
			api.FromFile.ToString(),
			result.ToString(),
		)
	}
	result = api.ToLoadConfigMethod(
		114514,
	)
	if result != api.Default {
		t.Errorf(
			"expected %s, got %s",
			api.Default.ToString(),
			result.ToString(),
		)
	}
}
