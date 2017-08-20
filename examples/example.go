package main

import (
	"github.com/shadowapex/go-godot/godot"
	"log"
)

// NewSimpleClass is a constructor that we can pass to godot.
func NewSimpleClass() godot.Class {
	simpleClass := &SimpleClass{
		Name: "MySimpleClass",
	}

	return simpleClass
}

// SimpleClass is a simple Godot class that can be registered.
type SimpleClass struct {
	godot.Node
	Name string
}

func (s *SimpleClass) Xready() {
	godot.Log.Warning("Simple class is ready!")
}

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// SetGodotGDNativeInit will set the given function to run on library initialization.
	godot.SetGodotGDNativeInit(func(options *godot.GodotGDNativeInitOptions) {
		log.Println("This is being called from example.go!")
	})

	// Register will register the given class constructor with Godot.
	godot.Register(NewSimpleClass)
}

// This never gets called, but it necessary to export as a shared library.
func main() {
}
