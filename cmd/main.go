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
	if err := generateByFileCMD.MarkFlagRequired("text"); err != nil {
		panic(err)
	}
	if err := generateByFileCMD.MarkFlagRequired("language"); err != nil {
		panic(err)
	}
	// Add Flag for addCharacterCMD
	addCharacterCMD.Flags().StringVarP(
		&NewCharacterName,
		"name",
		"n",
		"Remilia Scarlet",
		"The name for the new character",
	)
	if err := addCharacterCMD.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
	addCharacterCMD.Flags().StringVarP(
		&NewCharacterDescription,
		"description",
		"d",
		"Remilia Scarlet (レミリア・スカーレット Remiria Sukāretto) is a vampire who is the head of the Scarlet Devil Mansion. She is the sister of Flandre Scarlet, the mistress of Hong Meiling and Sakuya Izayoi as well as the other fairy maids and the friend of Patchouli Knowledge. She first appears as the Stage 6 boss and main antagonist of the Embodiment of Scarlet Devil and has been a playable character in multiple games since Imperishable Night, in a team with Sakuya, and Immaterial and Missing Power as well.",
		"The description for the new character",
	)
	if err := addCharacterCMD.MarkFlagRequired("description"); err != nil {
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
