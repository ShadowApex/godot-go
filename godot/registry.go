package godot

import (
	"reflect"
)

// registeredClass is a structure for holding on to the reflected details of a Go
// struct that has been registered as a Godot class. It has the struct's name,
// methods, the arguments that method takes, and the value types it returns.
type registeredClass struct {
	name       string
	structType reflect.Type
	methods    map[string]*registeredMethod
}

// newRegisteredClass will return a structure for
func newRegisteredClass(classType reflect.Type) *registeredClass {
	class := &registeredClass{
		name:       classType.String(),
		structType: classType,
		methods:    map[string]*registeredMethod{},
	}

	return class
}

// addMethod will add the given registered method to the class.
func (r *registeredClass) addMethod(name string, method *registeredMethod) {
	r.methods[name] = method
}

// registeredMethod is a structure for holding on to the reflected details of a Go
// method that has been registered as a Godot method. It contains the method's
// argument types and return types.
type registeredMethod struct {
	method    reflect.Method
	arguments []reflect.Type
	returns   []reflect.Type
}

// newRegisteredMethod takes in a struct type and uses reflection to discover all
// of the methods it has attached to it, and returns a registered method structure.
func newRegisteredMethod(classMethod reflect.Method) *registeredMethod {
	method := &registeredMethod{}
	method.method = classMethod

	// Get the type of the method.
	methodType := method.method.Type

	// Check the number of arguments the method takes and returns.
	numArgs := methodType.NumIn()
	numReturns := methodType.NumOut()
	method.arguments = make([]reflect.Type, numArgs)
	method.returns = make([]reflect.Type, numReturns)

	// Loop through the arguments and check what types of arguments they
	// take.
	for j := 0; j < numArgs; j++ {
		argType := methodType.In(j)
		method.arguments[j] = argType
	}

	// Loop through the returns and check what types it returns.
	for j := 0; j < numReturns; j++ {
		retType := methodType.Out(j)
		method.returns[j] = retType
	}

	return method
}

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

// classRegistry is a mapping of all classes that have been registered in Godot.
var classRegistry = map[string]*registeredClass{}

// instanceRegistry is a mapping of all instances that have currently been created. This map
// allows instance methods to find which instance they belong to.
var instanceRegistry = map[string]interface{}{}
