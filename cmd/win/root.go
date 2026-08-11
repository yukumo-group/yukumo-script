package win

import (
	"github.com/1Vewton/yukumo-script/utils"
	"github.com/1Vewton/yukumo-script/utils/logger"
	"github.com/spf13/cobra"
)

var cmdLogger = logger.NewLogger("CMD", nil)

// rootCMD defines the root command
var rootCMD = &cobra.Command{
	Use:   "yukumo",
	Short: "yukumo is a program that can generate yukumo audio",
	Long: `
Yukumo is a simple and flexible program that can generate yukumo audio without the need for network connection. 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CLIWelcome()
	},
}

func init() {
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
	generateByFileCMD.MarkFlagRequired("name")
	generateByFileCMD.MarkFlagRequired("text")
	generateByFileCMD.MarkFlagRequired("language")
	// Add subcommands
	phontsCMD.AddCommand(
		showAvailablePhontsCMD,
		playExampleCMD,
	)
	generationCMD.AddCommand(
		generateByFileCMD,
	)
	rootCMD.AddCommand(
		phontsCMD,
		generationCMD,
	)
}

// Execute executes the command
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		panic(err.Error())
	}
}
