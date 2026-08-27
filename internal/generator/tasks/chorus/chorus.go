package chorus

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yukumo-group/yukumo-script/internal/characters"
	"github.com/yukumo-group/yukumo-script/internal/generator/aquestalk2"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
	"golang.org/x/sync/errgroup"
)

// Task defines the chorus task for multiple characters
type Task struct {
	sync.RWMutex
	ID                string             `json:"id"`
	TaskName          string             `json:"taskName"`
	CharacterList     *[]string          `json:"characterList"`
	PhontList         *[]string          `json:"phontList"`
	ResultFile        *string            `json:"resultFile"`
	CreateTime        time.Time          `json:"createdTime"`
	EditTime          time.Time          `json:"editTime"`
	Text              string             `json:"text"`
	TaskLanguage      language.Language  `json:"taskLanguage"`
	Speed             int                `json:"speed"`
	MixingConfig      *edit.MixingConfig `json:"mixingConfig"`
	charactersManager *characters.Characters
	wavDir            string
}

// NewChorusTask creates new chorus task
func NewChorusTask(
	text string,
	phontList *[]string,
	characterList *[]string,
	speed int,
	taskName string,
	taskLanguage language.Language,
	mixingConfig *edit.MixingConfig,
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
		MixingConfig:      mixingConfig,
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

// GenerateWavName generates name for the result wav file
func (task *Task) GenerateWavFileName(
	targetDir string,
) string {
	return fmt.Sprintf(
		"%s/%s_%s_%d.wav",
		targetDir,
		task.TaskName,
		task.ID,
		task.EditTime.Unix(),
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
	task.Lock()
	task.EditTime = time.Now()
	speed := task.Speed
	tmpFilePath := task.GenerateTempFileName(
		generationMethod,
		idx,
	)
	task.Unlock()
	phontFilePath, err := tasks.PhontFile(
		phontsDir,
		phontName,
	)
	if err != nil {
		return nil, err
	}
	generator := aquestalk2.NewGenerator(
		speed,
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
	task.RLock()
	phontList := task.PhontList
	task.RUnlock()
	if phontList == nil {
		return nil, nil
	}
	resultFilePathChan := make(chan string, len(*phontList))
	group, _ := errgroup.WithContext(ctx)
	group.SetLimit(runtime.NumCPU() * 2)
	for i, phontName := range *phontList {
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

// GenerateAllAudiosPhonts generates all the files by characters.
// Returns the list of generated audios.
func (task *Task) GenerateAllAudiosCharacters(
	ctx context.Context,
	phontsDir string,
	convertedText string,
) ([]string, error) {
	task.RLock()
	characterList := task.CharacterList
	characterManager := task.charactersManager
	task.RUnlock()
	if characterList == nil {
		return nil, nil
	}
	resultFilePathChan := make(chan string, len(*characterList))
	group, _ := errgroup.WithContext(ctx)
	group.SetLimit(runtime.NumCPU() * 2)
	tmpCharacterList := characterManager.GetData()
	for i, characterID := range *characterList {
		tmpIdx := i
		group.Go(
			func() error {
				character, exists := tmpCharacterList[characterID]
				if !exists {
					return fmt.Errorf(
						"character %s does not exists",
						characterID,
					)
				}
				tmpPhontName := character.PhontName
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

// Generate generates mixed audio
func (task *Task) Generate(
	ctx context.Context,
	phontsDir string,
	targetDir string,
) error {
	task.RLock()
	task.EditTime = time.Now()
	fileName := task.GenerateWavFileName(
		targetDir,
	)
	taskLanguage := task.TaskLanguage
	taskText := task.Text
	taskMixingConfig := task.MixingConfig
	characterList := task.CharacterList
	phontList := task.PhontList
	task.RUnlock()
	length := 0
	if characterList != nil {
		length += len(*characterList)
	}
	if phontList != nil {
		length += len(*phontList)
	}
	if length == 0 {
		return errors.New(
			"the total length of character list and phont list cannot be 0",
		)
	}
	resultChan := make(chan string, length)
	// Start generation
	convertedText, err := language.ConvertText(
		taskText,
		taskLanguage,
	)
	if err != nil {
		return err
	}
	group, ctx := errgroup.WithContext(ctx)
	group.Go(
		func() error {
			results, err := task.GenerateAllAudiosCharacters(
				ctx,
				phontsDir,
				convertedText,
			)
			if err != nil {
				return err
			}
			for _, result := range results {
				resultChan <- result
			}
			return nil
		},
	)
	group.Go(
		func() error {
			results, err := task.GenerateAllAudiosPhonts(
				ctx,
				phontsDir,
				convertedText,
			)
			if err != nil {
				return err
			}
			for _, result := range results {
				resultChan <- result
			}
			return nil
		},
	)
	close(resultChan)
	resultList := []string{}
	for res := range resultChan {
		resultList = append(resultList, res)
	}
	err = edit.MixAudios(
		resultList,
		taskMixingConfig,
		fileName,
	)
	if err != nil {
		return err
	}
	task.Lock()
	task.ResultFile = &fileName
	task.Unlock()
	return nil
}
