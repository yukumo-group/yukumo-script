package audio

import (
	"bytes"
	"fmt"
	"os"

	"github.com/1Vewton/yukumo-script/utils/osoperation"
	"github.com/braheezy/shine-mp3/pkg/mp3"
	"github.com/go-audio/wav"
)

// WAV2MP3 converts .wav file to .mp3 file.
// This function contains code from the shine-mp3 project (https://github.com/braheezy/shine-mp3)
// Original copyright (c) 2023 braheezy, licensed under the GNU General Public License v2.
// As a derivative work, this file is also licensed under the same license.
func WAV2MP3(
	wavFileName string,
	targetDirectory string,
	targetFileName string,
) error {
	// Read wav data
	wavBytes, errReadWavFile := os.ReadFile(wavFileName)
	if errReadWavFile != nil {
		return errReadWavFile
	}
	wavReader := bytes.NewReader(wavBytes)
	wavDecoder := wav.NewDecoder(wavReader)
	wavBuffer, errDecode := wavDecoder.FullPCMBuffer()
	if errDecode != nil {
		return errDecode
	}
	// Convert audio data to int16
	decodedData := make([]int16, len(wavBuffer.Data))
	for i, data := range wavBuffer.Data {
		decodedData[i] = int16(data)
	}
	decodedData = UpdateChannelNumberTo2(decodedData)
	resampledDecodedData, errResample := ResampleWAV(
		decodedData,
		2,
		wavBuffer.Format.SampleRate,
		16000,
	)
	if errResample != nil {
		return errResample
	}
	// Create mp3 encoder
	mp3Encoder := mp3.NewEncoder(
		wavBuffer.Format.SampleRate,
		2,
	)
	mp3FilePath, errGetFilePath := osoperation.GetNewFilePath(
		targetDirectory,
		targetFileName,
		"mp3",
	)
	if errGetFilePath != nil {
		return errGetFilePath
	}
	out, errOpen := os.OpenFile(
		mp3FilePath,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if errOpen != nil {
		return errOpen
	}
	defer out.Close()
	errWrite := mp3Encoder.Write(out, resampledDecodedData)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

// ConvertAll converts wav to different formats
func ConvertAll(
	originalFilePath string,
	targetDirectory string,
	targetFileName string,
	format Format,
) error {
	switch format {
	case WAV:
		return osoperation.CopyFile(
			originalFilePath,
			targetDirectory,
			targetFileName,
			WAV.ToString(),
		)
	case MP3:
		return WAV2MP3(
			originalFilePath,
			targetDirectory,
			targetFileName,
		)
	default:
		return fmt.Errorf(
			"%s format does not supported",
			format.ToString(),
		)
	}
}
