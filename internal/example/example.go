package example

import (
	"github.com/yukumo-group/yukumo-script/pkg/utils/logger"
	"github.com/yukumo-group/yukumo-script/pkg/utils/syncutils"
)

var scriptLogger = logger.NewLogger("Example", nil)
var examplesMap = syncutils.NewMap()
