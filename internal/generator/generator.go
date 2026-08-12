package generator

// Generator interface defines methods of  generator
type Generator interface {
	// GenerateWav generates yukumo .wav file
	GenerateWav() error
}
