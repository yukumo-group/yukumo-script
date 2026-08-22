package main

import (
	"fmt"

	"github.com/fatih/color"
)

// ProcessError records the error
func ProcessError(
	err error,
) {
	errMessage := color.New(color.FgRed).Add(color.Bold)
	cmdLogger.Error(err.Error())
	_, _ = errMessage.Println(err.Error())
}

// ProcessErrorString processes the error using string
func ProcessErrorString(
	formatString string,
	contents ...any,
) {
	errMessage := color.New(color.FgRed).Add(color.Bold)
	resultString := fmt.Sprintf(
		formatString,
		contents...,
	)
	cmdLogger.Error(resultString)
	_, _ = errMessage.Println(resultString)
}
