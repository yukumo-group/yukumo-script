package language

// Language defines the language to use
type Language int

const (
	// Japanese : 0
	Japanese Language = iota
	// English : 1
	English
	// Chinese : 2
	Chinese
)

// ToInt converts language to integer
func (lang Language) ToInt() int {
	switch lang {
	case Japanese:
		return 0
	case English:
		return 1
	case Chinese:
		return 2
	default:
		return 0
	}
}

// ToLanguage converts integer to Language
func ToLanguage(langIdx int) Language {
	switch langIdx {
	case 0:
		return Japanese
	case 1:
		return English
	case 2:
		return Chinese
	default:
		return Japanese
	}
}
