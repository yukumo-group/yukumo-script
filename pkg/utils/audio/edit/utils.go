package edit

import (
	"math"
)

// GetBaseVolume gets the base of the volume.
// Return: Base, Volume
func GetBaseVolume(
	gain float64,
) (float64, float64) {
	if gain < 0 {
		return 2.0, math.Inf(-2)
	}
	volume := math.Log2(gain)
	return 2.0, volume
}
