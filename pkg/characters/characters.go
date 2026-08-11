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
	characters.fileName = fmt.Sprintf(
		"%s/%s",
		folder,
		file,
	)
}

// AddCharacter adds new character to the slice
func (characters *Characters) AddCharacter(
	characterID string,
	character *Character,
) error {
	if character != nil {
		characters.Lock()
		defer characters.Unlock()
		_, exists := characters.Data[characterID]
		if exists {
			return fmt.Errorf(
				"Character with Character ID %s already exists",
				characterID,
			)
		}
		characters.Data[characterID] = character
		return nil
	}
	return errors.New(
		"This characterID already exists",
	)
}

// saveTo saves to the target file
func (characters *Characters) saveTo(target string) error {
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
	_, errExist := os.Stat(characters.fileName)
	if errExist != nil {
		if os.IsNotExist(errExist) {
			file, errCreate := os.Create(characters.fileName)
			if errCreate != nil {
				return errCreate
			}
			defer file.Close()
			return characters.saveTo(characters.fileName)
		}
		return errExist
	}
	data, errRead := os.ReadFile(characters.fileName)
	if errRead != nil {
		return errRead
	}
	return json.Unmarshal(data, characters)
}

// SaveData saves the data to the file
func (characters *Characters) SaveData() error {
	_, errExist := os.Stat(characters.fileName)
	if errExist != nil {
		if os.IsNotExist(errExist) {
			file, errCreate := os.Create(characters.fileName)
			if errCreate != nil {
				return errCreate
			}
			defer file.Close()
			return characters.saveTo(characters.fileName)
		}
		return errExist
	}
	return characters.saveTo(characters.fileName)
}

// GetData gets the slice of characters
func (characters *Characters) GetData() map[string]*Character {
	characters.RLock()
	defer characters.RUnlock()
	return characters.Data
}
