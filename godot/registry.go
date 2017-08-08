package godot

import ()

// godotClassesToRegister is a slice of objects that will be registered as a Godot class
// upon library initialization.
var godotConstructorsToRegister = []ClassConstructor{}

// Expose will register the given object as a Godot class. It will be available
// inside Godot.
func Expose(constructor ClassConstructor) {
	godotConstructorsToRegister = append(godotConstructorsToRegister, constructor)
}

// typeRegistry is a mapping of all Godot class constructors that have been registered.
var typeRegistry = map[string]ClassConstructor{}
