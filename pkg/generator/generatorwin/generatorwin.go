package generatorwin

/*
#cgo LDFLAGS: -lkernel32
#include<YukumoGenerator.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// GeneratorWin generates yukumo audio for windows
type GeneratorWin struct {
	speed      int
	phontPath  string
	resultPath string
	text       string
}

// NewGeneratorWin creates new GeneratorWin struct
func NewGeneratorWin(
	speed int,
	phontPath string,
	resultPath string,
	text string,
) *GeneratorWin {
	return &GeneratorWin{
		speed:      speed,
		phontPath:  phontPath,
		resultPath: resultPath,
		text:       text,
	}
}

// GenerateWav generates yukumo .wav file on Windows
func (winGenerator *GeneratorWin) GenerateWav() error {
	encoder := japanese.ShiftJIS.NewEncoder()
	encode, _, errTrans := transform.String(encoder, winGenerator.text)
	if errTrans != nil {
		return errTrans
	}
	tmpPhontPath := C.CString(winGenerator.phontPath)
	tmpText := C.CString(encode)
	tmpResultPath := C.CString(winGenerator.resultPath)
	result := C.generate_wav(
		tmpPhontPath,
		tmpText,
		tmpResultPath,
		C.int(winGenerator.speed),
	)
	C.free(unsafe.Pointer(tmpPhontPath))
	C.free(unsafe.Pointer(tmpText))
	C.free(unsafe.Pointer(tmpResultPath))
	switch result {
	case 0:
		return nil
	case -1:
		return errors.New("file loading error")
	case -2:
		return errors.New("wav generating error")
	case -3:
		return errors.New("file opening error")
	case -4:
		return errors.New("write incomplete")
	default:
		return fmt.Errorf(
			"unexpected return %d",
			result,
		)
	}
}
