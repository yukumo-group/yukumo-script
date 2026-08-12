package characters

import (
	"github.com/fatih/color"
)

// Character defines the info of a character
type Character struct {
	Name             string  `json:"name"`
	PhontName        string  `json:"phontName"`
	Description      string  `json:"description"`
	ProfileImagePath *string `json:"profileImagePath"`
}

// NewCharacter creates new character
func NewCharacter(
	name string,
	phontName string,
	description string,
	profileImagePath *string,
) *Character {
	return &Character{
		Name:             name,
		PhontName:        phontName,
		Description:      description,
		ProfileImagePath: profileImagePath,
	}
}

// HasProfileImage checks if there is profile image
func (character *Character) HasProfileImage() bool {
	return character.ProfileImagePath != nil
}

// ShowInfo shows the info of the character
func (character *Character) ShowInfo(
	title *color.Color,
	text *color.Color,
) {
	_, _ = title.Println(character.Name)
	_, _ = text.Printf(
		"PhontName: %s\n",
		character.PhontName,
	)
	_, _ = text.Println(character.Description)
}
