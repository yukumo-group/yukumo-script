package osoperation

import (
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
