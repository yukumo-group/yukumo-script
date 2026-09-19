package main

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

// PolyphonicChinese defines the chinese phrase for adding
var PolyphonicChinese string

// PolyphonicPinyin defines the pinyin for the phrase
var PolyphonicPinyin string

// languageCMD defines the command for managing language
var languageCMD = &cobra.Command{
	Use:   "language",
	Short: "language command contains the commands that can manage the language related commands",
	Long:  "language command contains the commands that can manage the language related commands",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// showAllPolyphonicsCMD gets all stored polyphonics
var showAllPolyphonicsCMD = &cobra.Command{
	Use:   "showAllPolyphonics",
	Short: "showAllPolyphonics gets all added polyphonics",
	Long:  "showAllPolyphonics gets all added polyphonics",
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

// addPolyphonicCMD adds new polyphonics
var addPolyphonicCMD = &cobra.Command{
	Use:   "addPolyphonic",
	Short: "addPolyphonic adds new polyphonic pair",
	Long:  "addPolyphonic adds new polyphonic pair",
	Run: func(cmd *cobra.Command, args []string) {
		if PolyphonicChinese == "" || PolyphonicPinyin == "" {
			panic(
				"You cannot provide empty chinese and pinyin",
			)
		}
		err := api.AddPolyphonic(
			PolyphonicChinese,
			PolyphonicPinyin,
		)
		ProcessError(err)
	},
}
