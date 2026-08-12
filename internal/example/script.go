package example

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/yukumo-group/yukumo-script/internal/generator/aquestalk2"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language/all2jap"
	"golang.org/x/sync/errgroup"
)

// GenerateExamples synthesizes example WAVs for each phont file.
func GenerateExamples(
	ctx context.Context,
	targetDir string,
	phontDir string,
	files []os.DirEntry,
) error {
	group, _ := errgroup.WithContext(ctx)
	group.SetLimit(runtime.NumCPU() * 2)
	for _, i := range files {
		if !i.IsDir() {
			name := i.Name()
			group.Go(
				func() error {
					phontName := strings.TrimSuffix(
						name,
						filepath.Ext(name),
					)
					phontFile := fmt.Sprintf(
						"%s/%s",
						phontDir,
						name,
					)
					targetFile := fmt.Sprintf(
						"%s/example_%s.wav",
						targetDir,
						phontName,
					)
					scriptLogger.Info(
						fmt.Sprintf(
							"Phont: %s, Target: %s",
							phontFile,
							targetFile,
						),
					)
					_, errStat := os.Stat(targetFile)
					if errStat == nil {
						examplesMap.SetKV(
							phontName,
							targetFile,
						)
						return nil
					} else if os.IsNotExist(errStat) {
						generator := aquestalk2.NewGenerator(
							100,
							phontFile,
							targetFile,
							all2jap.AllToKana("僕はGopherです。"),
						)
						err := generator.GenerateWav()
						if err == nil {
							examplesMap.SetKV(
								phontName,
								targetFile,
							)
							return nil
						}
						scriptLogger.Error(err.Error())
						return err
					} else {
						return errStat
					}
				},
			)
		}
	}
	return group.Wait()
}
