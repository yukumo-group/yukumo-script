package aquestalk2

import (
	"fmt"
	"runtime"
)

// libPath returns the third_party AquesTalk2 shared library path for the current GOOS/GOARCH.
func libPath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			return "third_party/aquestalk2/win/lib64/AquesTalk2.dll", nil
		case "386":
			return "third_party/aquestalk2/win/lib/AquesTalk2.dll", nil
		}
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return "third_party/aquestalk2/linux/lib64/libAquesTalk2Eva.so.2.3", nil
		case "386":
			return "third_party/aquestalk2/linux/lib/libAquesTalk2Eva.so.2.3", nil
		}
	case "darwin":
		return "third_party/aquestalk2/mac/lib/libAquesTalk2Eva.dylib", nil
	}
	return "", fmt.Errorf(
		"unsupported platform %s/%s for AquesTalk2",
		runtime.GOOS,
		runtime.GOARCH,
	)
}
