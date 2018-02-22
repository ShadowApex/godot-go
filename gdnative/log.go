package gdnative

/*
#include <gdnative/string.h>
#include <gdnative/gdnative.h>
#include "gdnative.h"
*/
import "C"

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"github.com/vitaminwater/cgo.wchar"
)

// Log is used to log messages to Godot, and makes them viewable inside the
// Godot debugger.
var Log = &logger{}

// logger is a native Go structure for logging in Godot.
type logger struct{}

// Print will print the given message to the Godot debugger and console.
func (l *logger) Println(message ...interface{}) {
	// Convert our message into a Go string.
	goString := fmt.Sprint(message...)

	// Convert the go string into a godot string
	gdString := stringAsGodotString(goString)

	// Print our string.
	C.go_godot_print(GDNative.api, gdString)

	// Free the string from memory.
	C.go_godot_string_destroy(GDNative.api, gdString)
}

// Warning will print a warning message to the Godot debugger and console.
func (l *logger) Warning(message ...interface{}) {
	l.log(false, message...)
}

// Error will print an error message to the Godot debugger and console.
func (l *logger) Error(message ...interface{}) {
	l.log(true, message...)
}

// Write will call logger.Println from the given bytes, to implement the io.Writer
// interface.
func (l *logger) Write(data []byte) (int, error) {
	buffer := bytes.NewBuffer(data)
	line := strings.TrimRight(buffer.String(), "\n")
	l.Println(line)

	return len(data), nil
}

// log is a helper function that will log either Warnings or Errors in Godot.
func (l *logger) log(isError bool, message ...interface{}) {
	// Convert the messages into a Go string.
	goDescription := fmt.Sprint(message...)

	// Get the caller filename and line number.
	pc, file, no, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Unable to get caller!")
		return
	}

	// Get the function name.
	funcName := "unknown"
	details := runtime.FuncForPC(pc)
	if details != nil {
		funcName = details.Name()
	}

	// Convert the go string into a C string
	cDescription := C.CString(goDescription)
	cFuncName := C.CString(funcName)
	cFile := C.CString(file)
	cLine := C.int(no)

	if isError {
		C.go_godot_print_error(GDNative.api, cDescription, cFuncName, cFile, cLine)
		return
	}
	C.go_godot_print_warning(GDNative.api, cDescription, cFuncName, cFile, cLine)
}

func stringAsGodotString(value string) *C.godot_string {
	// Convert the Go string into a wchar
	wcharString, err := wchar.FromGoString(value)
	if err != nil {
		fmt.Println("Error decoding string:", err)
	}

	// Build the Godot string with the wchar
	var godotString C.godot_string
	C.go_godot_string_new_with_wide_string(
		GDNative.api,
		&godotString,
		(*C.wchar_t)(wcharString.Pointer()),
		C.int(len(value)),
	)

	return &godotString
}
