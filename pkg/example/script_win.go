//go:build windows
// +build windows

package example

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/yukumo-group/yukumo-script/pkg/generator/generatorwin"
	"github.com/yukumo-group/yukumo-script/pkg/language/all2jap"
	"golang.org/x/sync/errgroup"
)

// GenerateExampleWin generates examples for phont file in win64
func GenerateExampleWin(
	ctx context.Context,
	targetDir string,
	phontDir string,
	files []os.DirEntry,
) error {
	group, ctx := errgroup.WithContext(ctx)
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
						// Add Example if generation is successful
						examplesMap.SetKV(
							phontName,
							targetFile,
						)
						return nil
					} else if os.IsNotExist(errStat) {
						generatorW := generatorwin.NewGeneratorWin(
							100,
							phontFile,
							targetFile,
							all2jap.AllToKana("僕はGopherです。"),
						)
						err := generatorW.GenerateWav()
						if err == nil {
							// Add Example if generation is successful
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
	err := group.Wait()
	return err
}
