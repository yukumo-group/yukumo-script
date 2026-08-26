package chorus

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/aquestalk2"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
	"golang.org/x/sync/errgroup"
)

// Task defines the chorus task for multiple characters
type Task struct {
	ID                string            `json:"id"`
	TaskName          string            `json:"taskName"`
	CharacterList     *[]string         `json:"characterList"`
	PhontList         *[]string         `json:"phontList"`
	ResultFile        *string           `json:"resultFile"`
	CreateTime        time.Time         `json:"createdTime"`
	EditTime          time.Time         `json:"editTime"`
	Text              string            `json:"text"`
	TaskLanguage      language.Language `json:"taskLanguage"`
	Speed             int               `json:"speed"`
	charactersManager *characters.Characters
	wavDir            string
}

// NewChorustTask creates new chorus task
func NewChorustTask(
	text string,
	phontList *[]string,
	characterList *[]string,
	speed int,
	taskName string,
	taskLanguage language.Language,
	charactersManager *characters.Characters,
	wavDir string,
) (*Task, error) {
	id := uuid.NewString()
	if characterList == nil && phontList == nil {
		return nil, errors.New(
			"you have to choose at least one way to generate the audio",
		)
	}
	return &Task{
		ID:                id,
		Text:              text,
		TaskLanguage:      taskLanguage,
		TaskName:          taskName,
		CharacterList:     characterList,
		CreateTime:        time.Now(),
		EditTime:          time.Now(),
		charactersManager: charactersManager,
		wavDir:            wavDir,
	}, nil
}

// GenerateTempFileName generates filename of metadata for this task
func (task *Task) GenerateTempFileName(
	generationMethod string,
	idx int,
) string {
	return fmt.Sprintf(
		"%s/%s_%s_%d_%s_%d.wav",
		task.wavDir,
		task.TaskName,
		task.ID,
		task.EditTime.Unix(),
		generationMethod,
		idx,
	)
}

// GenerateSingleFile generates a single file for certain phont name
func (task *Task) GenerateSingleFile(
	phontsDir string,
	phontName string,
	convertedText string,
	generationMethod string,
	idx int,
) (*string, error) {
	task.EditTime = time.Now()
	tmpFilePath := task.GenerateTempFileName(
		generationMethod,
		idx,
	)
	phontFilePath, err := tasks.PhontFile(
		phontsDir,
		phontName,
	)
	if err != nil {
		return nil, err
	}
	generator := aquestalk2.NewGenerator(
		task.Speed,
		phontFilePath,
		tmpFilePath,
		convertedText,
	)
	err = generator.GenerateWav()
	if err != nil {
		return nil, err
	}
	return &tmpFilePath, nil
}

// GenerateAllAudiosPhonts generates all the files by phonts.
// Returns the list of generated audios.
func (task *Task) GenerateAllAudiosPhonts(
	ctx context.Context,
	phontsDir string,
	convertedText string,
) ([]string, error) {
	if task.PhontList == nil {
		return nil, nil
	}
	resultFilePathChan := make(chan string, len(*task.PhontList))
	group, _ := errgroup.WithContext(ctx)
	group.SetLimit(runtime.NumCPU() * 2)
	for i, phontName := range *task.PhontList {
		tmpIdx := i
		tmpPhontName := phontName
		group.Go(
			func() error {
				resultFilePath, err := task.GenerateSingleFile(
					phontsDir,
					tmpPhontName,
					convertedText,
					"Phont",
					tmpIdx,
				)
				if err != nil {
					return err
				}
				if resultFilePath == nil {
					return fmt.Errorf(
						"result for phont number %d is nil",
						tmpIdx,
					)
				}
				resultFilePathChan <- *resultFilePath
				return nil
			},
		)
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}
	close(resultFilePathChan)
	result := []string{}
	for resultPath := range resultFilePathChan {
		result = append(result, resultPath)
	}
	return result, nil
}
