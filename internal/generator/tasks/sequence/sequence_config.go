package sequence

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
	"go.yaml.in/yaml/v4"
)

// RawCharacter defines the raw config of character
type RawCharacter struct {
	PhontName     string  `yaml:"phont_name"`
	CharacterName string  `yaml:"character_name"`
	Description   *string `yaml:"description"`
}

// ToCharacter converts RawCharacter to Character
func (rawCharacter RawCharacter) ToCharacter() *characters.Character {
	return characters.NewCharacter(
		rawCharacter.CharacterName,
		rawCharacter.PhontName,
		*rawCharacter.Description,
		nil,
	)
}

// RawConfig defines the raw config read from yaml
type RawConfig struct {
	UsePredefinedCharacters *bool           `yaml:"use_predefined_characters"`
	ConfigName              *string         `yaml:"config_name"`
	DefaultSpeed            *int            `yaml:"default_speed"`
	Characters              *[]RawCharacter `yaml:"raw_characters"`
	Language                int             `yaml:"language"`
}

// TaskConfig defines the config for sequence tasks
type TaskConfig struct {
	ConfigName   string
	Characters   *characters.Characters
	DefaultSpeed int
	TaskLanguage language.Language
}

// ReadRawConfig reads the raw config yaml file
func ReadRawConfig(
	filePath string,
) (*RawConfig, error) {
	newRawConfig := RawConfig{}
	newRawConfig.Characters = &[]RawCharacter{}
	fileData, err := os.ReadFile(
		filePath,
	)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(
		fileData,
		&newRawConfig,
	)
	if err != nil {
		return nil, err
	}
	if newRawConfig.ConfigName == nil {
		fullFileName := filepath.Base(filePath)
		suffix := filepath.Ext(fullFileName)
		resultFileName := strings.TrimSuffix(
			fullFileName,
			suffix,
		)
		newRawConfig.ConfigName = &resultFileName
	}
	return &newRawConfig, nil
}

// ToTaskConfig converts RawConfig to TaskConfig
func (config *RawConfig) ToTaskConfig() (*TaskConfig, error) {
	thisCharacterList := characters.NewCharacters()
	// Set Characters
	var useDefaultPredefinedCharacters bool
	if config.UsePredefinedCharacters == nil {
		useDefaultPredefinedCharacters = false
	} else {
		useDefaultPredefinedCharacters = *config.UsePredefinedCharacters
	}
	if useDefaultPredefinedCharacters {
		thisCharacterList = characters.CharacterList
	} else {
		if config.Characters == nil {
			return nil, errors.New(
				"the character list cannot be nil when you do not want to use predefined characters",
			)
		}
		for _, rawCharacter := range *config.Characters {
			newProcessedCharacter := rawCharacter.ToCharacter()
			thisCharacterList.AddCharacter(newProcessedCharacter)
		}
	}
	// Set Default Speed
	var speed int
	if config.DefaultSpeed == nil {
		speed = 100
	} else {
		speed = *config.DefaultSpeed
	}
	// Config name
	if config.ConfigName == nil {
		return nil, errors.New(
			"the config name cannot be nil when converting to task config",
		)
	}
	return &TaskConfig{
		Characters:   thisCharacterList,
		TaskLanguage: language.ToLanguage(config.Language),
		DefaultSpeed: speed,
		ConfigName:   *config.ConfigName,
	}, nil
}
