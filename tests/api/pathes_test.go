package api

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/api"
)

func TestPathesSetting(t *testing.T) {
	t.Parallel()
	testPathes := &api.FilePathes{}
	testPathes.SetPathes(
		api.WithAssetsDir("114514"),
	)
	if testPathes.AssetsDir != "114514" {
		t.Errorf(
			"expected %s, got %s",
			"114514",
			testPathes.AssetsDir,
		)
	}
}
