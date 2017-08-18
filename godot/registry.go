package godot

import (
	"reflect"
)

// godotClassesToRegister is a slice of objects that will be registered as a Godot class
// upon library initialization.
var godotConstructorsToRegister = []ClassConstructor{}

// Expose will register the given object as a Godot class. It will be available
// inside Godot.
func Expose(constructor ClassConstructor) {
	godotConstructorsToRegister = append(godotConstructorsToRegister, constructor)
}

// constructorRegistry is a mapping of all Godot class constructors that have been registered.
var constructorRegistry = map[string]ClassConstructor{}

// typeRegistry is a mapping of all Godot class types that have been registered.
var typeRegistry = map[string]reflect.Type{}

// instanceRegistry is a mapping of all instances that have currently been created. This map
// allows instance methods to find which instance they belong to.
var instanceRegistry = map[string]interface{}{}
