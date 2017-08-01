package main

import (
	"fmt"
	"github.com/shadowapex/go-godot/godot"
)

// SetGodotGDNativeInit will set the given function to run on library initialization.
var gdNativeInit = godot.SetGodotGDNativeInit(func(options *godot.GodotGDNativeInitOptions) {
	fmt.Println("This is in example.go!")
})

// This never gets called.
func main() {
}
