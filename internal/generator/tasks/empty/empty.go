package empty

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/unitoftime/beep"
	"github.com/unitoftime/beep/wav"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
)

// Task creates new empty task
type Task struct {
	ID          string
	SampleRate  int
	Length      float64
	CreatedTime time.Time
	EditTime    time.Time
	ResultFile  *string
	NumChannels int
	Precision   int
}

// NewEmptyTask creates new empty task.
// Length is in seconds.
func NewEmptyTask(
	length float64,
	data *audio.Info,
) (*Task, error) {
	if data == nil {
		return nil, errors.New(
			"the input data cannot be nil",
		)
	}
	id := uuid.NewString()
	return &Task{
		ID:          id,
		Length:      length,
		SampleRate:  data.SampleRate,
		CreatedTime: time.Now(),
		EditTime:    time.Now(),
		NumChannels: data.ChannelNumber,
		Precision:   data.Precision,
	}, nil
}

// Generate generates empty audio.
// phont path should let empty "".
// As a file for combining with other files, it is recommended to save it at wav directory
func (task *Task) Generate(
	ctx context.Context,
	phontsDir string,
	targetDir string,
) error {
	task.EditTime = time.Now()
	// Create file
	targetFileDir := fmt.Sprintf(
		"%s/empty_%s_%d.wav",
		targetDir,
		task.ID,
		task.EditTime.Unix(),
	)
	file, errCreateFile := os.Create(targetFileDir)
	if errCreateFile != nil {
		return errCreateFile
	}
	closeFile := func() {
		_ = file.Close()
	}
	defer closeFile()
	// Generate empty audio
	totalSamples := task.Length *
		float64(task.SampleRate) *
		float64(task.NumChannels)
	silenceStream := beep.Silence(int(totalSamples))
	errEncode := wav.Encode(
		file,
		silenceStream,
		beep.Format{
			SampleRate:  beep.SampleRate(task.SampleRate),
			NumChannels: task.NumChannels,
			Precision:   task.Precision,
		},
	)
	if errEncode != nil {
		return errEncode
	}
	task.ResultFile = &targetFileDir
	return nil
}
