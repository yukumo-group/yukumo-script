package main

import (
	"github.com/yukumo-group/yukumo-script/pkg/api"
	"github.com/yukumo-group/yukumo-script/pkg/utils/logger"
)

var cliLogger = logger.NewLogger(
	"CLI",
	nil,
)

// Initialize directories and shared runtime state
func init() {
	// Initialize
	if err := api.Init(); err != nil {
		cliLogger.Error(err.Error())
		panic(err)
	}
	// Add flags for generateByFile
	generateByFileCMD.Flags().StringVarP(
		&SingleSentenceTaskNameByFile,
		"name",
		"n",
		"task",
		"task name of the new generation task",
	)
	generateByFileCMD.Flags().StringVarP(
		&SingleSentenceTaskTextByFile,
		"text",
		"t",
		"",
		"text to generate",
	)
	generateByFileCMD.Flags().IntVarP(
		&SingleSentenceTaskLanguageByFile,
		"language",
		"l",
		0,
		"0: Japanese, 1: English, 2: Chinese",
	)
	generateByFileCMD.Flags().IntVarP(
		&SingleSentenceTaskSpeedByFile,
		"speed",
		"s",
		100,
		"Speed of the audio(default: 100)",
	)
	if err := generateByFileCMD.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
	if err := generateByFileCMD.MarkFlagRequired("text"); err != nil {
		panic(err)
	}
	if err := generateByFileCMD.MarkFlagRequired("language"); err != nil {
		panic(err)
	}
	// Add subcommands
	charactersCMD.AddCommand(
		addCharacterCMD,
	)
	phontsCMD.AddCommand(
		showAvailablePhontsCMD,
		playExampleCMD,
	)
	generationCMD.AddCommand(
		generateByFileCMD,
	)
	tasksCMD.AddCommand(
		showAllSingleSentenceTasksCMD,
	)
	rootCMD.AddCommand(
		phontsCMD,
		generationCMD,
		tasksCMD,
		charactersCMD,
	)
}

// Main process
func main() {
	Execute()
}
