//go:build !clib

package example

// GetAllExampleFont gets the font name of all available phonts
func GetAllExampleFont() []string {
	return examplesMap.GetAllKeys()
}
