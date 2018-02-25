package main

import (
	"fmt"
	"github.com/shadowapex/godot-go/gdnative"
)

// SimpleClass is a structure that we can register with Godot.
type SimpleClass struct {
	base gdnative.Object
}

// Instances is a map of our created Godot classes. This will be populated when
// Godot calls the CreateFunc.
var Instances = map[string]*SimpleClass{}

// NativeScriptInit will run on NativeScript initialization. It is responsible
// for registering all our classes with Godot.
func nativeScriptInit() {
	gdnative.Log.Warning("Initializing nativescript from Go!")

	// Define an instance creation function. This will be called when Godot
	// creates a new instance of our class.
	createFunc := gdnative.InstanceCreateFunc{
		CreateFunc: simpleConstructor,
		MethodData: "SIMPLE",
		FreeFunc:   func(methodData string) {},
	}

	// Define an instance destroy function. This will be called when Godot
	// asks our library to destroy our class instance.
	destroyFunc := gdnative.InstanceDestroyFunc{
		DestroyFunc: simpleDestructor,
		MethodData:  "SIMPLE",
		FreeFunc:    func(methodData string) {},
	}

	// Register our class with Godot.
	gdnative.Log.Warning("Registering SIMPLE class...")
	gdnative.NativeScript.RegisterClass(
		"SIMPLE",
		"Reference",
		&createFunc,
		&destroyFunc,
	)

	// Register a method with Godot.
	gdnative.Log.Warning("Registering SIMPLE method...")
	gdnative.NativeScript.RegisterMethod(
		"SIMPLE",
		"get_data",
		&gdnative.MethodAttributes{
			RPCType: gdnative.MethodRpcModeDisabled,
		},
		&gdnative.InstanceMethod{
			Method:     simpleMethod,
			MethodData: "SIMPLE",
			FreeFunc:   func(methodData string) {},
		},
	)
}

func simpleConstructor(object gdnative.Object) string {
	gdnative.Log.Println("Creating new SimpleClass...")

	// Create a new instance of our struct.
	instance := &SimpleClass{
		base: object,
	}

	// Use the pointer address as the instance ID
	instanceID := fmt.Sprintf("%p", instance)
	Instances[instanceID] = instance

	// Return the instanceID
	return instanceID
}

func simpleDestructor(object gdnative.Object, methodData, userData string) {
	gdnative.Log.Println("Destroying SimpleClass with ID:", userData, "...")
	// Delete the instance from our map of instances
	delete(Instances, userData)
}

func simpleMethod(object gdnative.Object, methodData, userData string, numArgs int, args []gdnative.Variant) gdnative.Variant {
	gdnative.Log.Println("SIMPLE.get_data() called!")

	data := gdnative.NewStringWithWideString("World from godot-go!")
	ret := gdnative.NewVariantWithString(data)

	return ret
}

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// Set the initialization script that will run upon NativeScript initialization.
	// This function will handle using the NativeScript API to register all of our
	// classes.
	gdnative.SetNativeScriptInit(nativeScriptInit)
}

// This never gets called, but it necessary to export as a shared library.
func main() {
}
