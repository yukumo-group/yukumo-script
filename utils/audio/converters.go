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
) ([]byte, error) {
	// Read wav data
	wavBytes, errReadWavFile := os.ReadFile(wavFileName)
	if errReadWavFile != nil {
		return nil, errReadWavFile
	}
	wavReader := bytes.NewReader(wavBytes)
	wavDecoder := wav.NewDecoder(wavReader)
	wavBuffer, errDecode := wavDecoder.FullPCMBuffer()
	if errDecode != nil {
		return nil, errDecode
	}
	// Convert audio data to int16
	decodedData := make([]int16, len(wavBuffer.Data))
	for i, data := range wavBuffer.Data {
		decodedData[i] = int16(data)
	}
	// Create mp3 encoder
	mp3Encoder := mp3.NewEncoder(
		wavBuffer.Format.SampleRate,
		wavBuffer.Format.NumChannels,
	)
	var outBuffer bytes.Buffer
	// Encode to mp3
	errEncode := mp3Encoder.Write(
		&outBuffer,
		decodedData,
	)
	if errEncode != nil {
		return nil, errEncode
	}
	mp3Data := outBuffer.Bytes()
	return mp3Data, nil
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
		resultByte, errConvert := WAV2MP3(
			originalFilePath,
		)
		if errConvert != nil {
			return errConvert
		}
		return osoperation.SaveDataTo(
			targetDirectory,
			targetFileName,
			MP3.ToString(),
			resultByte,
		)
	default:
		return fmt.Errorf(
			"%s format does not supported!",
			format.ToString(),
		)
	}
}
