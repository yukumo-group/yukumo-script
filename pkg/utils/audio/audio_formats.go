package audio

// Format defines the type of the audio
type Format string

const (
	// WAV : .wav
	WAV Format = "wav"
	// MP3 : mp3
	MP3 Format = "mp3"
	// AAC : m4a
	AAC Format = "m4a"
	// FLAC : flac
	FLAC Format = "flac"
)

// GetAllFormats gets all the formats of audios to transform
func GetAllFormats() []Format {
	return []Format{
		WAV,
		MP3,
		AAC,
		FLAC,
	}
}

// ToString converts Format to string
func (audioType Format) ToString() string {
	switch audioType {
	case WAV:
		return "wav"
	case MP3:
		return "mp3"
	case AAC:
		return "m4a"
	case FLAC:
		return "flac"
	default:
		return "wav"
	}
}

// ToFormat converts string to AudioFormat
func ToFormat(
	audioType string,
) Format {
	switch audioType {
	case "wav":
		return WAV
	case "mp3":
		return MP3
	case "m4a":
		return AAC
	case "flac":
		return FLAC
	default:
		return WAV
	}
}
