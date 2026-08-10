package osoperation

import (
	"fmt"
	"os"
)

// SaveDataTo saves certain data to certain file
func SaveDataTo(
	targetDirectory string,
	targetFileName string,
	suffix string,
	data []byte,
) error {
	targetFilePath := fmt.Sprintf(
		"%s/%s.%s",
		targetDirectory,
		targetFileName,
		suffix,
	)
	_, errTargetFileExists := os.Stat(targetFilePath)
	if errTargetFileExists != nil {
		if os.IsNotExist(errTargetFileExists) {
			file, errCreateFile := os.Create(targetFilePath)
			if errCreateFile != nil {
				return errCreateFile
			}
			file.Close()
		} else {
			return errTargetFileExists
		}
	}
	errWriteToTargetFile := os.WriteFile(
		targetFilePath,
		data,
		0644,
	)
	if errWriteToTargetFile != nil {
		return errWriteToTargetFile
	}
	return nil
}

// CopyFile copies the file to a new file or overwrite the original file
func CopyFile(
	originalFilePath string,
	targetDirectory string,
	targetFileName string,
	suffix string,
) error {
	_, errOriginalFileExists := os.Stat(originalFilePath)
	if errOriginalFileExists != nil {
		if os.IsNotExist(errOriginalFileExists) {
			return fmt.Errorf(
				"%s file does not exists, you cannot copy file that does not exists",
				originalFilePath,
			)
		}
		return errOriginalFileExists
	}
	originalContent, errReadOriginalFile := os.ReadFile(
		originalFilePath,
	)
	if errReadOriginalFile != nil {
		return errReadOriginalFile
	}
	errWriteToTargetFile := SaveDataTo(
		targetDirectory,
		targetFileName,
		suffix,
		originalContent,
	)
	if errWriteToTargetFile != nil {
		return errWriteToTargetFile
	}
	return nil
}
