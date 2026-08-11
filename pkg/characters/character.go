package characters

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
	if character.ProfileImagePath == nil {
		return false
	}
	return true
}

// ShowInfo shows the info of the character
func (character *Character) ShowInfo() {

}
