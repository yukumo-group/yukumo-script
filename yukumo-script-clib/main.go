package main

/*
 */
import "C"
import (
	"github.com/yukumo-group/yukumo-script/phontsmanager"
)

// InitializePhontNameToFileName initializes phont name to file name
// export InitializePhontNameToFileName
func InitializePhontNameToFileName(
	phontsDir string,
) C.int {
	phontsmanager.InitializePhontNameToFileName(
		phontsDir,
	)
	return 0
}
