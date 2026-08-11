package example

import (
	"github.com/yukumo-group/yukumo-script/utils/logger"
	"github.com/yukumo-group/yukumo-script/utils/syncutils"
)

var scriptLogger = logger.NewLogger("Example", nil)
var examplesMap = syncutils.NewMap()
