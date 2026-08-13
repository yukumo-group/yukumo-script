package api

import (
	"context"

	"github.com/yukumo-group/yukumo-script/internal/example"
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
)

// GenerateExamples generates all the examples
func GenerateExamples() error {
	dir, err := phontsmanager.GetAllPhonts(filePathForProg.PhontsDir)
	if err != nil {
		return err
	}
	if err := example.GenerateExamples(
		context.Background(),
		filePathForProg.ExampleDir,
		filePathForProg.PhontsDir,
		dir,
	); err != nil {
		return err
	}
	return nil
}

// GetAllExamplePhont gets all the example phonts
func GetAllExamplePhont() []string {
	return example.GetAllExamplePhont()
}
