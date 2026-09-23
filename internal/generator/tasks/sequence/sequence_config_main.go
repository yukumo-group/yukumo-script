package sequence

var defaultUsePredefinedCharacters = true

// DefaultConfig defines the default configuration
var DefaultConfig *RawConfig = NewRawConfig(
	"default",
	1,
	&defaultUsePredefinedCharacters,
	nil,
	nil,
)
