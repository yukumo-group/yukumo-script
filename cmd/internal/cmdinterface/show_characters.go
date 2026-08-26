package cmdinterface

import (
	"fmt"
	"yukumo-script-cmd/internal/info"

	"github.com/fatih/color"
)

// GetCharacter asks the user which character to use (Return in form of character ID or name)
func GetCharacter(
	title *color.Color,
	text *color.Color,
) (string, error) {
	info.ShowAllCharacters(title, text)
	var targetCharacterID string
	fmt.Scan(&targetCharacterID)
	return "", nil
}
