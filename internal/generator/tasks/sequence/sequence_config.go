package sequence

import (
	"os"

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

// RawConfig defines the raw config read from yaml
type RawConfig struct {
	UsePredefinedCharacters *bool          `yaml:"use_predefined_characters"`
	TaskName                *string        `yaml:"task_name"`
	DefaultSpeed            *int           `yaml:"default_speed"`
	Characters              []RawCharacter `yaml:"raw_characters"`
	Language                int            `yaml:"language"`
}

// TaskConfig defines the config for sequence tasks
type TaskConfig struct {
	TaskName     string
	Characters   *characters.Characters
	DefaultSpeed int
	TaskLanguage language.Language
}

// ReadRawConfig reads the raw config yaml file
func ReadRawConfig(
	filePath string,
) (*RawConfig, error) {
	newRawConfig := RawConfig{}
	newRawConfig.Characters = []RawCharacter{}
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
	return &newRawConfig, nil
}
