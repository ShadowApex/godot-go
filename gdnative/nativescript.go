// Package nativescript provides a wrapper around Godot's nativescript
// extension. It exists to provide a way to use Go as an alternative scripting
// language from GDScript.
package gdnative

/*
#include <nativescript/godot_nativescript.h>
*/
import "C"

import (
	"log"
	"unsafe"
)

// Handle is a pointer to the nativescript handler. It must be passed to any
// Godot nativescript functions. This will be populated when 'godot_nativescript_init'
// is called by Godot upon script initialization.
var handle unsafe.Pointer

/** Script entry (Registering all the classes and stuff) **/
// godot_nativescript_init is the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes,
// etc. The `unsafe.Pointer` type is used to represent a void C pointer.
//export godot_nativescript_init
func godot_nativescript_init(hdl unsafe.Pointer) {
	log.Println("Initializing NativeScript")
	handle = hdl
}

// RegisterClass will register the given class with Godot. This will make it
// available to be attached to a Node in Godot. The name of the class that you
// provide here will be the name that you specify when you attach a NativeScript
// script to a Node. The base parameter that you specify will be what the class
// should be inheriting from (e.g. Node2D, Node, etc.).
//
// The create and destroy function parameters are C structs that contain
// function pointers to C methods that will be called when the object is created
// or destroyed. Because of the pointer passing rules, Go code can not pass a
// function value directly to C, so a gateway function should be used. More
// information on using a gateway function can be found here:
// https://github.com/golang/go/wiki/cgo#function-variables
//func registerClass(name, base string, createFunc C.godot_instance_create_func, destroyFunc C.godot_instance_destroy_func) {
//	C.godot_nativescript_register_class(
//		handle,
//		C.CString(name),
//		C.CString(base),
//		createFunc,
//		destroyFunc,
//	)
//}
//
//// RegisterToolClass will register the given class with Godot as a tool. Refer to
//// the 'RegisterClass' method for more information on how to use this.
//func registerToolClass(name, base string, createFunc C.godot_instance_create_func, destroyFunc C.godot_instance_destroy_func) {
//	C.godot_nativescript_register_tool_class(
//		handle,
//		C.CString(name),
//		C.CString(base),
//		createFunc,
//		destroyFunc,
//	)
//}
//
//// RegisterMethod will register the given method with Godot and associate it with
//// the given class name. The name parameter is the name of the class you wish to
//// attach this method to. The funcName parameter is the name of the function you
//// want to register. The attributes and method are what will actually be called
//// when Godot calls the method on the object.
//func registerMethod(name, funcName string, attributes C.godot_method_attributes, method C.godot_instance_method) {
//	C.godot_nativescript_register_method(
//		handle,
//		C.CString(name),
//		C.CString(funcName),
//		attributes,
//		method,
//	)
//}
//
//// RegisterProperty will register the given property with Godot and associate it
//// with the given class name. The name parameter is the name of the class you wish
//// to attach this property to. The path is the name of the property itself. The
//// attributes and setter/getter methods are what will be called when Godot gets
//// or sets this property.
//func registerProperty(name, path string, attributes *C.godot_property_attributes, setFunc C.godot_property_set_func, getFunc C.godot_property_get_func) {
//	C.godot_nativescript_register_property(
//		handle,
//		C.CString(name),
//		C.CString(path),
//		attributes,
//		setFunc,
//		getFunc,
//	)
//}
//
//// RegisterSignal will register the given signal with Godot.
//func registerSignal(name string, signal *C.godot_signal) {
//	C.godot_nativescript_register_signal(
//		handle,
//		C.CString(name),
//		signal,
//	)
//}
