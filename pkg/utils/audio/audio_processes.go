package audio

// Processes stores the processes that can be done on audio files
type Process string

const (
	// Resample defines the process of resample the wav
	Resample Process = "Resample"
	// NotAvailable defines undefined process
	NotAvailable Process = "NotAvailable"
)

// ToString converts processes to string
func (process Process) ToString() string {
	switch process {
	case Resample:
		return "Resample"
	case NotAvailable:
		return "NotAvailable"
	default:
		return "NotAvailable"
	}
}

// ToProcess converts string to process
func ToProcess(
	content string,
) Process {
	switch content {
	case "Resample":
		return Resample
	case "NotAvailable":
		return NotAvailable
	default:
		return NotAvailable
	}
}
