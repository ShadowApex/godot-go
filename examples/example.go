package main

import (
	"fmt"
	"github.com/shadowapex/go-godot/godot"
)

type SimpleClass struct {
	Name string
}

func (s *SimpleClass) Ready() {
	fmt.Println("GO: SimpleClass is ready!")
}

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// SetGodotGDNativeInit will set the given function to run on library initialization.
	godot.SetGodotGDNativeInit(func(options *godot.GodotGDNativeInitOptions) {
		fmt.Println("GO: This is being called from example.go!")
	})

	// RegisterClass will register the given class with Godot.
	simpleClass := &SimpleClass{Name: "MySimpleClass"}
	godot.RegisterClass(simpleClass)
}

// This never gets called.
func main() {
}
