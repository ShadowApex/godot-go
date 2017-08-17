package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
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
	gdString := godotString(goString)

	// Print our string.
	C.godot_print(&gdString)

	// Free the string from memory.
	C.godot_string_destroy(&gdString)
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
		C.godot_print_error(cDescription, cFuncName, cFile, cLine)
		return
	}
	C.godot_print_warning(cDescription, cFuncName, cFile, cLine)
}

// godotString will convert the given string to a Godot string.
func godotString(str string) C.godot_string {
	// Create a Godot String and convert the Go string into a C string.
	var gdString C.godot_string
	cString := C.CString(str)
	C.godot_string_new_data(&gdString, cString, C.int(len(str)))

	return gdString
}
