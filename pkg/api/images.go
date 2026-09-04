package api

import (
	"fmt"
	"os"

	"github.com/yukumo-group/yukumo-script/pkg/utils/osoperation"
)

// ImageInfo stores the data for
type ImageInfo struct {
	FileName string
	FileType string
}

// SaveFileAsImages saves file to images
func SaveFileAsImages(
	originalFilePath string,
	targetName string,
) error {
	suffix, supported := osoperation.IsSuffixSupported(
		originalFilePath,
		osoperation.SupportedImageSuffix,
	)
	if !supported {
		return fmt.Errorf(
			"the file %s's suffix does not supported",
			originalFilePath,
		)
	}
	errCopy := osoperation.CopyFile(
		originalFilePath,
		filePathForProg.ImagesDir,
		targetName,
		suffix,
	)
	return errCopy
}

// GetAllImages gets all the images.
// Show this before the user wants to set profile image for characters.
func GetAllImages() ([]ImageInfo, error) {
	allImages := []ImageInfo{}
	allPossibleImages, errReadDir := os.ReadDir(
		filePathForProg.ImagesDir,
	)
	if errReadDir != nil {
		return nil, errReadDir
	}
	for _, possibleImage := range allPossibleImages {
		fileName := possibleImage.Name()
		suffix, supported := osoperation.IsSuffixSupported(
			fileName,
			osoperation.SupportedImageSuffix,
		)
		if supported {
			allImages = append(
				allImages,
				ImageInfo{
					FileName: fileName,
					FileType: suffix,
				},
			)
		}
	}
	return allImages, nil
}

// CheckImageExists checks if this image exists
func CheckImageExists(
	path string,
) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
