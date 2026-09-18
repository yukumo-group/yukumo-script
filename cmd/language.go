package main

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

// languageCMD defines the command for managing language
var languageCMD = &cobra.Command{
	Use:   "language",
	Short: "language command contains the commands that can manage the language related commands",
	Long:  "language command contains the commands that can manage the language related commands",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// getAllPolyphonicsCMD gets all stored polyphonics
var getAllPolyphonicsCMD = &cobra.Command{
	Use:   "getAllPolyphonics",
	Short: "getAllPolyphonics gets all added polyphonics",
	Long:  "getAllPolyphonics gets all added polyphonics",
	Run: func(cmd *cobra.Command, args []string) {
		polyphonics := api.GetAllPolyphonics()
		title := color.New(color.FgGreen).Add(color.Bold)
		text := color.New(color.Italic)
		for chinese, pinyin := range polyphonics {
			title.Printf("%s ", chinese)
			text.Printf("%s\n", pinyin)
		}
	},
}
