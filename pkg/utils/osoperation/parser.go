package osoperation

import (
	"path/filepath"
	"strings"
)

// ParseWindowsPath parses the path for windows
func ParseWindowsPath(
	originalPath string,
) string {
	return strings.ReplaceAll(
		originalPath,
		`\`,
		"/",
	)
}

// IsSuffixSupported checks if the suffix is supported and returns the suffix and a boolean to show if it exists
func IsSuffixSupported(
	fileName string,
	supportedSuffix []string,
) (string, bool) {
	suffix := filepath.Ext(fileName)
	for _, supported := range supportedSuffix {
		if supported == suffix {
			return suffix, true
		}
	}
	return suffix, false
}
