// package name: libgodot
package main

/*
#cgo CXXFLAGS: -I./godot -std=c11
#include <stddef.h>
#include "godot.h"
*/
import "C"

import "fmt"

// Godot entry point

//export godot_native_init
func godot_native_init(options *C.godot_native_init_options) {
	fmt.Println("Hey from go!")
	fmt.Println(options)
}

// Godot termination point

//export godot_native_terminate
func godot_native_terminate() {
	fmt.Println("Exiting go!")
}

func InitOptions() {
	var options C.godot_native_init_options
	fmt.Println(options)
}

// main is needed to be exported as a shared library
func main() {}
