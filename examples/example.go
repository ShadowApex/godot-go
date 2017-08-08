package main

import (
	"github.com/shadowapex/go-godot/godot"
	"log"
)

// NewSimpleClass is a constructor that we can pass to godot.
func NewSimpleClass() interface{} {
	simpleClass := &SimpleClass{
		Name: "MySimpleClass",
		Base: "Node2D",
	}
	simpleClass.Ready = func() {
		log.Println("SimpleClass is ready!")
	}

	return simpleClass
}

// SimpleClass is a simple Godot class that can be registered.
type SimpleClass struct {
	Name  string
	Base  string `godot:"_inherits"`
	Ready func() `godot:"_ready"`
}

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// SetGodotGDNativeInit will set the given function to run on library initialization.
	godot.SetGodotGDNativeInit(func(options *godot.GodotGDNativeInitOptions) {
		log.Println("This is being called from example.go!")
	})

	// Expose will register the given class constructor with Godot.
	godot.Expose(NewSimpleClass)
}

// This never gets called, but it necessary to export as a shared library.
func main() {
}
