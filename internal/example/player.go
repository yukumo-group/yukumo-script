package example

// GetAllExamplePhont gets the font name of all available phonts
func GetAllExamplePhont() []string {
	return examplesMap.GetAllKeys()
}
