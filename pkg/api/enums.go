package api

// LoadConfigMethod defines the way of loading config
type LoadConfigMethod int

const (
	// FromFile loads config from file.
	// 0
	FromFile LoadConfigMethod = iota
	// FromManager loads config from the manager.
	// 1
	FromManager
	// Default loads default manager.
	// 2
	Default
)

// ToLoadConfigMethod converts integer to LoadConfigManager
func ToLoadConfigMethod(
	data int,
) LoadConfigMethod {
	switch data {
	case 0:
		return FromFile
	case 1:
		return FromManager
	case 2:
		return Default
	default:
		return Default
	}
}
