package edit

// MixingMethod defines the way of mixing audios
type MixingMethod int

const (
	// ByDefault (0) will not change the volume of each subaudios.
	// This might cause wave cutting in real circumstances
	ByDefault MixingMethod = iota
	// ByAverage (1) will average the volume of each subaudios
	ByAverage
	// ByCustom (2) allows the user to define their own proportion of volume
	ByCustom
)

// ToInt converts mixing method to int
func (mixingMethod MixingMethod) ToInt() int {
	switch mixingMethod {
	case ByDefault:
		return 0
	case ByAverage:
		return 1
	case ByCustom:
		return 2
	default:
		return 1
	}
}

// ToMixingMethod converts integer to mixing method
func ToMixingMethod(
	data int,
) MixingMethod {
	switch data {
	case 0:
		return ByDefault
	case 1:
		return ByAverage
	case 2:
		return ByCustom
	default:
		return ByAverage
	}
}

// MixingConfig defines the config of mixing
type MixingConfig struct {
	Method     MixingMethod
	AudioGains *[]float64
}

// NewMixingMethod creates new mixing method
func NewMixingMethod(
	Method MixingMethod,
	gains *[]float64,
) *MixingConfig {
	method := Method
	if gains == nil && method == ByCustom {
		method = ByAverage
	}
	return &MixingConfig{
		Method:     method,
		AudioGains: gains,
	}
}
