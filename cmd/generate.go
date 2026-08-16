package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

// SingleSentenceTaskNameByFile defines the name of the task
var SingleSentenceTaskNameByFile string

// SingleSentenceTaskTextByFile defines the text to generate for the task
var SingleSentenceTaskTextByFile string

// SingleSentenceTaskSpeedByFile defines the speed of the audio
var SingleSentenceTaskSpeedByFile int

// SingleSentenceTaskLanguageByFile defines the language of the audio
var SingleSentenceTaskLanguageByFile int

// generationCommand contains commands for generation
var generationCMD = &cobra.Command{
	Use:   "generation",
	Short: "generation command contains subcommands that can generate audio through pre-defined characters or phont files",
	Long: `
subcommands of generation command allow you to generate your audio
	`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// generateByFile generates wav by file
var generateByFileCMD = &cobra.Command{
	Use:   "generateByFile",
	Short: "Generate yukumo audio through the file",
	Long: `
generateByFile allows you to generate yukumo audio through phont file directly
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Define the format of the texts
		title := color.New(color.FgGreen).Add(color.Bold)
		text := color.New(color.Italic)
		// Print info
		_, _ = title.Println("Here are the available phonts:")
		for _, phontName := range api.ListPhonts() {
			_, _ = text.Println(phontName)
		}
		_, _ = title.Println("Input the name of the phont you want to use to generate audio:")
		// Input
		var phontName string
		_, errInput := fmt.Scan(&phontName)
		if errInput != nil {
			ProcessError(errInput)
			return
		}
		var taskName string
		if SingleSentenceTaskNameByFile != "" {
			taskName = api.RandomTaskName("SingleSentence")
		} else {
			taskName = SingleSentenceTaskNameByFile
		}
		newGenerationParam := api.NewGenerateByPhontParams(
			taskName,
			SingleSentenceTaskTextByFile,
			SingleSentenceTaskLanguageByFile,
			SingleSentenceTaskSpeedByFile,
			phontName,
		)
		result, errGenerate := api.GenerateByPhont(newGenerationParam)
		if errGenerate != nil {
			ProcessError(errGenerate)
			return
		}
		_, _ = title.Printf(
			"File saved at %s",
			result.ResultFile,
		)
	},
}
