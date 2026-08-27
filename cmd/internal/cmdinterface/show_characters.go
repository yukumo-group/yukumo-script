package cmdinterface

import (
	"fmt"
	"yukumo-script-cmd/internal/info"

	"github.com/fatih/color"
	"github.com/yukumo-group/yukumo-script/pkg/api"
)

// GetCharacter asks the user which character to use (Return in form of character ID or name)
func GetCharacter(
	title *color.Color,
	text *color.Color,
) (string, error) {
	_, _ = title.Println("Here are the available characters: ")
	info.ShowAllCharacters(title, text)
	_, _ = title.Println("Input the name of the character you want to use:")
	var targetCharacterID string
	_, err := fmt.Scan(&targetCharacterID)
	if err != nil {
		return "", err
	}
	characters := api.GetAllCharacters()
	for _, character := range characters {
		if character.Name == targetCharacterID {
			return targetCharacterID, nil
		}
	}
	return "", fmt.Errorf(
		"%s does not exists in characters",
		targetCharacterID,
	)
}
