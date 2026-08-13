package characters

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

// Characters stores the list of characters
type Characters struct {
	sync.RWMutex
	Data     map[string]*Character `json:"data"`
	fileName string
}

// NewCharacters creates new Characters
func NewCharacters() *Characters {
	return &Characters{
		Data: make(map[string]*Character),
	}
}

// SetTargetFile sets the target file to store the Characters
func (characters *Characters) SetTargetFile(
	folder string,
	file string,
) {
	characters.Lock()
	defer characters.Unlock()
	characters.fileName = fmt.Sprintf(
		"%s/%s",
		folder,
		file,
	)
}

// AddCharacter adds new character to the slice
func (characters *Characters) AddCharacter(
	character *Character,
) error {
	if character != nil {
		characters.Lock()
		defer characters.Unlock()
		_, exists := characters.Data[character.Name]
		if exists {
			return fmt.Errorf(
				"character with character ID %s already exists",
				character.Name,
			)
		}
		characters.Data[character.Name] = character
		return nil
	}
	return errors.New(
		"character cannot be nil",
	)
}

// UpdateCharacter change the info inside a character
func (characters *Characters) ChangeCharacter(
	newCharacter *Character,
) error {
	if newCharacter != nil {
		characters.Lock()
		defer characters.Unlock()
		_, exists := characters.Data[newCharacter.Name]
		if !exists {
			return fmt.Errorf(
				"character with character ID %s does not exists",
				newCharacter.Name,
			)
		}
		characters.Data[newCharacter.Name] = newCharacter
		return nil
	}
	return errors.New(
		"character cannot be nil",
	)
}

// saveTo saves to the target file
func (characters *Characters) saveTo(target string) error {
	characters.RLock()
	defer characters.RUnlock()
	jsonData, errJSON := json.Marshal(characters)
	if errJSON != nil {
		return errJSON
	}
	errWrite := os.WriteFile(
		target,
		jsonData,
		0644,
	)
	return errWrite
}

// ReadData reads the data inside the file stored
func (characters *Characters) ReadData() error {
	characters.Lock()
	_, errExist := os.Stat(characters.fileName)
	if errExist != nil {
		if os.IsNotExist(errExist) {
			file, errCreate := os.Create(characters.fileName)
			if errCreate != nil {
				characters.Unlock()
				return errCreate
			}
			if err := file.Close(); err != nil {
				characters.Unlock()
				return err
			}
			characters.Unlock()
			return characters.saveTo(characters.fileName)
		}
		characters.Unlock()
		return errExist
	}
	data, errRead := os.ReadFile(characters.fileName)
	if errRead != nil {
		characters.Unlock()
		return errRead
	}
	characters.Unlock()
	return json.Unmarshal(data, characters)
}

// SaveData saves the data to the file
func (characters *Characters) SaveData() error {
	characters.RLock()
	_, errExist := os.Stat(characters.fileName)
	if errExist != nil {
		if os.IsNotExist(errExist) {
			file, errCreate := os.Create(characters.fileName)
			if errCreate != nil {
				characters.RUnlock()
				return errCreate
			}
			if err := file.Close(); err != nil {
				characters.RUnlock()
				return err
			}
			characters.RUnlock()
			return characters.saveTo(characters.fileName)
		}
		characters.RUnlock()
		return errExist
	}
	characters.RUnlock()
	return characters.saveTo(characters.fileName)
}

// GetData gets the slice of characters
func (characters *Characters) GetData() map[string]*Character {
	characters.RLock()
	defer characters.RUnlock()
	return characters.Data
}
