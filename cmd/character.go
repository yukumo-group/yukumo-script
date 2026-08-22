package main

import (
	"yukumo-script-cmd/internal/cmdinterface"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

// NewCharacterName defines the name of the new character
var NewCharacterName string

// NewCharacterDiscription defines the description for the new character
var NewCharacterDescription string

// charactersCMD contains commands for managing commands
var charactersCMD = &cobra.Command{
	Use:   "characters",
	Short: "characters contains subcommands for managing characters",
	Long:  "characters contains subcommands for managing characters",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// addCharacterCMD adds new character through cli
var addCharacterCMD = &cobra.Command{
	Use:   "addCharacter",
	Short: "addCharacter adds new character",
	Long:  "addCharacter adds new character to the main character list.",
	Run: func(cmd *cobra.Command, args []string) {
		// Define the format of the texts
		title := color.New(color.FgGreen).Add(color.Bold)
		text := color.New(color.Italic)
		if api.IsCharacterExists(NewCharacterName) {
			ProcessErrorString(
				"%s already exists in the character list",
				NewCharacterName,
			)
			return
		}
		phontName, err := cmdinterface.ShowPhonts(
			title,
			text,
		)
		if err != nil {
			ProcessError(err)
			return
		}
		err = api.AddCharacter(
			NewCharacterName,
			phontName,
			NewCharacterDescription,
			nil,
		)
		if err != nil {
			ProcessError(err)
		}
	},
}
