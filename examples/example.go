package main

import (
	"fmt"
	"github.com/shadowapex/go-godot/godot"
)

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// SetGodotGDNativeInit will set the given function to run on library initialization.
	godot.SetGodotGDNativeInit(func(options *godot.GodotGDNativeInitOptions) {
		fmt.Println("GO: This is being called from example.go!")
	})
}

// This never gets called.
func main() {
}
