package aquestalk2

/*
#cgo windows LDFLAGS: -lkernel32
#cgo linux LDFLAGS: -ldl
#include "YukumoGenerator.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

// Generator synthesizes yukumo audio via AquesTalk2 loaded at runtime.
type Generator struct {
	speed      int
	phontPath  string
	resultPath string
	text       string
}

// NewGenerator creates a Generator for the given synthesis parameters.
func NewGenerator(
	speed int,
	phontPath string,
	resultPath string,
	text string,
) *Generator {
	return &Generator{
		speed:      speed,
		phontPath:  phontPath,
		resultPath: resultPath,
		text:       text,
	}
}

// GenerateWav generates a yukumo .wav file using AquesTalk2_Synthe_Utf8.
func (g *Generator) GenerateWav() error {
	path, err := libPath()
	if err != nil {
		return err
	}
	cLibPath := C.CString(path)
	cPhontPath := C.CString(g.phontPath)
	cText := C.CString(g.text)
	cResultPath := C.CString(g.resultPath)
	defer C.free(unsafe.Pointer(cLibPath))
	defer C.free(unsafe.Pointer(cPhontPath))
	defer C.free(unsafe.Pointer(cText))
	defer C.free(unsafe.Pointer(cResultPath))

	result := C.generate_wav(
		cLibPath,
		cPhontPath,
		cText,
		cResultPath,
		C.int(g.speed),
	)
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
		return fmt.Errorf("unexpected return %d", result)
	}
}
