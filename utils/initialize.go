package utils

import (
	"fmt"
	"os"

	"github.com/yukumo-group/yukumo-script/utils/logger"
)

var initializeLogger = logger.NewLogger(
	"Initialization",
	nil,
)

// InitializeDirectory creates a directory if it does not exists
func InitializeDirectory(directoryName string) {
	_, errStat := os.Stat(directoryName)
	if errStat == nil {
		initializeLogger.Info(
			fmt.Sprintf(
				"%s directory already exists \n",
				directoryName,
			),
		)
	} else if os.IsNotExist(errStat) {
		err := os.Mkdir(directoryName, 0644)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic(errStat.Error())
	}
}

// InitializeFile creates a file if it does not exists
func InitializeFile(fileName string) {
	_, errStat := os.Stat(fileName)
	if errStat == nil {
		initializeLogger.Info(
			fmt.Sprintf(
				"%s file already exists \n",
				fileName,
			),
		)
	} else if os.IsNotExist(errStat) {
		file, err := os.Create(fileName)
		if err != nil {
			defer file.Close()
			panic(err.Error())
		}
	} else {
		panic(errStat.Error())
	}
}
