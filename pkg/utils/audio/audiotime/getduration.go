package audiotime

import (
	"time"
)

// GetDuration gets time duration in float64
func GetDuration(
	duration float64,
) time.Duration {
	durationInNano := duration * 1e9
	return time.Duration(durationInNano)
}
