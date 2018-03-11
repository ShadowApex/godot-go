package gdnative

/*
#include <gdnative/string.h>
#include <gdnative/gdnative.h>
#include "gdnative.gen.h"
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
var Log = &Logger{StackNum: 2}

// Logger is a native Go structure for logging in Godot.
type Logger struct {
	// StackNum is how far up the stack logs should show up as.
	StackNum int
}

// Print will print the given message to the Godot debugger and console.
func (l *Logger) Println(message ...interface{}) {
	GDNative.checkInit()

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
func (l *Logger) Warning(message ...interface{}) {
	GDNative.checkInit()
	l.log(false, message...)
}

// Error will print an error message to the Godot debugger and console.
func (l *Logger) Error(message ...interface{}) {
	GDNative.checkInit()
	l.log(true, message...)
}

// Write will call Logger.Println from the given bytes, to implement the io.Writer
// interface.
func (l *Logger) Write(data []byte) (int, error) {
	GDNative.checkInit()
	buffer := bytes.NewBuffer(data)
	line := strings.TrimRight(buffer.String(), "\n")
	l.Println(line)

	return len(data), nil
}

// log is a helper function that will log either Warnings or Errors in Godot.
func (l *Logger) log(isError bool, message ...interface{}) {
	// Convert the messages into a Go string.
	goDescription := fmt.Sprint(message...)

	// Get the caller filename and line number.
	pc, file, no, ok := runtime.Caller(l.StackNum)
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
	var godotString C.godot_string
	if value == "" {
		C.go_godot_string_new(GDNative.api, &godotString)
		return &godotString
	}

	// Convert the Go string into a wchar
	wcharString, err := wchar.FromGoString(value)
	if err != nil {
		fmt.Println("Error decoding string '"+value+"': ", err)
		C.go_godot_string_new(GDNative.api, &godotString)
		return &godotString
	}

	// Build the Godot string with the wchar
	C.go_godot_string_new_with_wide_string(
		GDNative.api,
		&godotString,
		(*C.wchar_t)(wcharString.Pointer()),
		C.int(len(value)),
	)

	return &godotString
}
