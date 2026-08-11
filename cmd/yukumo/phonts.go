package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yukumo-group/yukumo-script/pkg/example"
)

// phontsCMD contains subcommands for managing phonts
var phontsCMD = &cobra.Command{
	Use:   "phonts",
	Short: "phonts contains subcommands for managing phonts",
	Long: `
You can manage phonts and play examples here.
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// showAvailablePhontsCMD shows all the available phonts
var showAvailablePhontsCMD = &cobra.Command{
	Use:   "showAvailablePhonts",
	Short: "showAvailablePhonts shows the name of all available phonts",
	Long: `
ShowAvailablePhonts shows all the available phonts that can be used to generate audio
	`,
	Run: func(cmd *cobra.Command, args []string) {
		title := color.New(color.FgGreen).Add(color.Bold)
		title.Println("Here are the available phonts:")
		text := color.New(color.Italic)
		for _, phontName := range example.GetAllExampleFont() {
			text.Println(phontName)
		}
	},
}

// playExampleCMD plays the example for the phont
var playExampleCMD = &cobra.Command{
	Use:   "playExample",
	Short: "playExample allows you to play example generated for certain phonts",
	Long:  "playExample allows you to play example generated for certain phonts",
	Run: func(cmd *cobra.Command, args []string) {
		// Define the format of the texts
		title := color.New(color.FgGreen).Add(color.Bold)
		errMessage := color.New(color.FgRed).Add(color.Bold)
		text := color.New(color.Italic)
		// Print info
		title.Println("Here are the available phonts:")
		for _, phontName := range example.GetAllExampleFont() {
			text.Println(phontName)
		}
		title.Println("Input the name of the phont you want to play:")
		// Input
		var phontName string
		_, errInput := fmt.Scan(&phontName)
		if errInput != nil {
			cmdLogger.Error(errInput.Error())
			errMessage.Println(errInput.Error())
			return
		}
		// Play
		file, err := example.PlayExample(phontName)
		if err != nil {
			cmdLogger.Error(err.Error())
			if file != nil {
				cmdLogger.Error(
					fmt.Sprintf(
						"Error occurs when reading %s",
						*file,
					),
				)
			}
			errMessage.Println(err.Error())
		}
	},
}
