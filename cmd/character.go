package main

import (
	"github.com/spf13/cobra"
)

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
}
