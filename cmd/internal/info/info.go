package info

import (
	"github.com/fatih/color"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

// ShowAllCharacters shows all the characters
func ShowAllCharacters(
	title *color.Color,
	text *color.Color,
) {
	characters := api.GetAllCharacters()
	for _, character := range characters {
		_, _ = title.Println(character.Name)
		_, _ = text.Println(character.Name)
	}
}
