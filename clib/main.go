package main

/*
#include <stdlib.h>
#include <string.h>

// ErrorMessage manages the result of a golang function for C callers.
typedef struct ErrorMessage {
	int code;
	char* information;
} ErrorMessage;

static ErrorMessage new_error_message(int code, char* information) {
	ErrorMessage msg;
	msg.code = code;
	msg.information = information;
	return msg;
}

// StringList is a NULL-terminated list of C strings allocated by this library.
typedef struct StringList {
	char** items;
	int count;
} StringList;
*/
import "C"
import (
	"context"
	"unsafe"

	"github.com/yukumo-group/yukumo-script/pkg/api"
)

func errMessage(code int, err error) C.ErrorMessage {
	if err == nil {
		return C.new_error_message(C.int(code), nil)
	}
	return C.new_error_message(C.int(code), C.CString(err.Error()))
}

func okMessage() C.ErrorMessage {
	return C.new_error_message(0, nil)
}

func toStringList(values []string) C.StringList {
	n := len(values)
	if n == 0 {
		return C.StringList{items: nil, count: 0}
	}
	// Allocate array of char* plus a trailing NULL for C convenience.
	size := C.size_t(n+1) * C.size_t(unsafe.Sizeof(uintptr(0)))
	raw := C.malloc(size)
	if raw == nil {
		return C.StringList{items: nil, count: 0}
	}
	items := unsafe.Slice((**C.char)(raw), n+1)
	for i, v := range values {
		items[i] = C.CString(v)
	}
	items[n] = nil
	return C.StringList{
		items: (**C.char)(raw),
		count: C.int(n),
	}
}

//export YukumoInit
func YukumoInit() C.ErrorMessage {
	if err := api.Init(); err != nil {
		return errMessage(1, err)
	}
	return okMessage()
}

//export YukumoListPhonts
func YukumoListPhonts() C.StringList {
	return toStringList(api.ListPhonts())
}

//export YukumoListTasks
func YukumoListTasks() C.StringList {
	return toStringList(api.ListTasks())
}

//export YukumoGenerateByPhont
func YukumoGenerateByPhont(
	taskName *C.char,
	text *C.char,
	language C.int,
	speed C.int,
	phontName *C.char,
	outResultFile **C.char,
) C.ErrorMessage {
	ctx := context.Background()
	if outResultFile != nil {
		*outResultFile = nil
	}
	result, err := api.GenerateByPhont(
		ctx,
		api.NewGenerateByPhontParams(
			C.GoString(taskName),
			C.GoString(text),
			int(language),
			int(speed),
			C.GoString(phontName),
		),
	)
	if err != nil {
		return errMessage(1, err)
	}
	if outResultFile != nil {
		*outResultFile = C.CString(result.ResultFile)
	}
	return okMessage()
}

//export YukumoFreeString
func YukumoFreeString(s *C.char) {
	if s != nil {
		C.free(unsafe.Pointer(s))
	}
}

//export YukumoFreeStringList
func YukumoFreeStringList(list C.StringList) {
	if list.items == nil {
		return
	}
	n := int(list.count)
	items := unsafe.Slice(list.items, n)
	for i := 0; i < n; i++ {
		if items[i] != nil {
			C.free(unsafe.Pointer(items[i]))
		}
	}
	C.free(unsafe.Pointer(list.items))
}

//export YukumoFreeErrorMessage
func YukumoFreeErrorMessage(msg C.ErrorMessage) {
	if msg.information != nil {
		C.free(unsafe.Pointer(msg.information))
	}
}

func main() {}
