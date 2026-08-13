package api

import (
	"fmt"

	"github.com/yukumo-group/yukumo-script/internal/characters"
)

// GetAllCharacters gets all the data
func GetAllCharacters() map[string]*characters.Character {
	return characters.CharacterList.GetData()
}

// IsCharacterExists checks if characters exists
func IsCharacterExists(characterName string) bool {
	allCharacters := GetAllCharacters()
	_, exists := allCharacters[characterName]
	return exists
}

// AddCharacter adds new character
func AddCharacter(
	characterName string,
	phontName string,
	description string,
	profileImagePath *string,
) error {
	// Check if phont exists
	phontExists := IsPhontExists(
		phontName,
	)
	if phontExists {
		return fmt.Errorf(
			"phont %s does not exists",
			phontName,
		)
	}
	// Add characrer
	newCharacter := characters.NewCharacter(
		characterName,
		phontName,
		description,
		profileImagePath,
	)
	errAddCharacter := characters.CharacterList.AddCharacter(
		newCharacter,
	)
	return errAddCharacter
}

// ChangeCharacter changes certain character
func ChangeCharacter(
	characterName string,
	phontName string,
	description string,
	profileImagePath *string,
) error {
	// Check if phont exists
	phontExists := IsPhontExists(
		phontName,
	)
	if phontExists {
		return fmt.Errorf(
			"phont %s does not exists",
			phontName,
		)
	}
	// Add characrer
	newCharacter := characters.NewCharacter(
		characterName,
		phontName,
		description,
		profileImagePath,
	)
	errAddCharacter := characters.CharacterList.ChangeCharacter(
		newCharacter,
	)
	return errAddCharacter
}

// GetPhontNameByCharacterName gets phont name by character name
func GetPhontNameByCharacterName(
	characterName string,
) (string, error) {
	allCharacters := GetAllCharacters()
	thisCharacter, exists := allCharacters[characterName]
	if !exists {
		return "", fmt.Errorf(
			"character with name %s does not exists",
			characterName,
		)
	}
	phontName := thisCharacter.PhontName
	if !IsPhontExists(phontName) {
		return "", fmt.Errorf(
			"phont with name %s does not exists",
			phontName,
		)
	}
	return phontName, nil
}
