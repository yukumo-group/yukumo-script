package edit

import (
	"errors"
	"fmt"
	"os"

	"github.com/unitoftime/beep"
	"github.com/unitoftime/beep/wav"
)

// CheckCanCombine checks if two audios have same precision and other data
func CheckCanCombine(
	formats []beep.Format,
) error {
	if len(formats) < 1 {
		return errors.New(
			"length of formats list cannot be 0",
		)
	}
	firstSampleRate := formats[0].SampleRate
	for i := 1; i < len(formats); i++ {
		if formats[i].SampleRate != firstSampleRate {
			return fmt.Errorf(
				"the format of wav file number %d got sample rate %d, expected %d",
				i+1,
				formats[i].SampleRate,
				firstSampleRate,
			)
		}
	}
	firstPrecision := formats[0].Precision
	for i := 1; i < len(formats); i++ {
		if formats[i].Precision != firstPrecision {
			return fmt.Errorf(
				"the format of wav file number %d got precision %d, expected %d",
				i+1,
				formats[i].Precision,
				firstPrecision,
			)
		}
	}
	firstNumChannels := formats[0].NumChannels
	for i := 1; i < len(formats); i++ {
		if formats[i].NumChannels != firstNumChannels {
			return fmt.Errorf(
				"the format of wav file number %d got channels number %d, expected %d",
				i+1,
				formats[i].NumChannels,
				firstNumChannels,
			)
		}
	}
	return nil
}

// SpliceAudios splices the audio.
func SpliceAudios(
	targetFileName string,
	audioPathes []string,
) error {
	streamers := []beep.Streamer{}
	formats := []beep.Format{}
	for _, audioPath := range audioPathes {
		file, errReadFile := os.Open(audioPath)
		if errReadFile != nil {
			return errReadFile
		}
		closeFile := func() {
			_ = file.Close()
		}
		defer closeFile()
		streamer, format, errDecode := wav.Decode(
			file,
		)
		if errDecode != nil {
			return errDecode
		}
		closeStream := func() {
			_ = streamer.Close()
		}
		defer closeStream()
		streamers = append(streamers, streamer)
		formats = append(formats, format)
	}
	errCanCombine := CheckCanCombine(formats)
	if errCanCombine != nil {
		return errCanCombine
	}
	combinedStream := beep.Seq(streamers...)
	outFile, errCreateOutFile := os.Create(targetFileName)
	if errCreateOutFile != nil {
		return errCreateOutFile
	}
	closeOutFile := func() {
		_ = outFile.Close()
	}
	defer closeOutFile()
	errEncode := wav.Encode(
		outFile,
		combinedStream,
		formats[0],
	)
	if errEncode != nil {
		return errEncode
	}
	return nil
}
