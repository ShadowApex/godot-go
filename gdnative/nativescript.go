// Package gdnative provides a wrapper around Godot's nativescript
// extension. It exists to provide a way to use Go as an alternative scripting
// language from GDScript.
package gdnative

/*
#include <nativescript/godot_nativescript.h>
#include "nativescript.gen.h"
#include "nativescript.h"
#include "variant.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

type PropertyUsageFlags int

func (p PropertyUsageFlags) getBase() C.godot_property_usage_flags {
	return C.godot_property_usage_flags(p)
}

const (
	PropertyUsageStorage             PropertyUsageFlags = 1
	PropertyUsageEditor              PropertyUsageFlags = 2
	PropertyUsageNetwork             PropertyUsageFlags = 4
	PropertyUsageEditorHelper        PropertyUsageFlags = 8
	PropertyUsageCheckable           PropertyUsageFlags = 16  //used for editing global variables
	PropertyUsageChecked             PropertyUsageFlags = 32  //used for editing global variables
	PropertyUsageInternationalized   PropertyUsageFlags = 64  //hint for internationalized strings
	PropertyUsageGroup               PropertyUsageFlags = 128 //used for grouping props in the editor
	PropertyUsageCategory            PropertyUsageFlags = 256
	PropertyUsageStoreIfNonZero      PropertyUsageFlags = 512  //only store if nonzero
	PropertyUsageStoreIfNonOne       PropertyUsageFlags = 1024 //only store if false
	PropertyUsageNoInstanceState     PropertyUsageFlags = 2048
	PropertyUsageRestartIfChanged    PropertyUsageFlags = 4096
	PropertyUsageScriptVariable      PropertyUsageFlags = 8192
	PropertyUsageStoreIfNull         PropertyUsageFlags = 16384
	PropertyUsageAnimateAsTrigger    PropertyUsageFlags = 32768
	PropertyUsageUpdateAllIfModified PropertyUsageFlags = 65536

	PropertyUsageDefault     PropertyUsageFlags = PropertyUsageStorage | PropertyUsageEditor | PropertyUsageNetwork
	PropertyUsageDefaultIntl PropertyUsageFlags = PropertyUsageStorage | PropertyUsageEditor | PropertyUsageNetwork | PropertyUsageInternationalized
	PropertyUsageNoEditor    PropertyUsageFlags = PropertyUsageStorage | PropertyUsageNetwork
)

// PropertyUsageFlags is a string-based lookup table of constants for PropertyUsageFlags.
var PropertyUsageFlagsLookupMap = map[string]PropertyUsageFlags{
	"PropertyUsageStorage":             PropertyUsageStorage,
	"PropertyUsageEditor":              PropertyUsageEditor,
	"PropertyUsageNetwork":             PropertyUsageNetwork,
	"PropertyUsageEditorHelper":        PropertyUsageEditorHelper,
	"PropertyUsageCheckable":           PropertyUsageCheckable,
	"PropertyUsageChecked":             PropertyUsageChecked,
	"PropertyUsageInternationalized":   PropertyUsageInternationalized,
	"PropertyUsageGroup":               PropertyUsageGroup,
	"PropertyUsageCategory":            PropertyUsageCategory,
	"PropertyUsageStoreIfNonZero":      PropertyUsageStoreIfNonZero,
	"PropertyUsageStoreIfNonOne":       PropertyUsageStoreIfNonOne,
	"PropertyUsageNoInstanceState":     PropertyUsageNoInstanceState,
	"PropertyUsageRestartIfChanged":    PropertyUsageRestartIfChanged,
	"PropertyUsageScriptVariable":      PropertyUsageScriptVariable,
	"PropertyUsageStoreIfNull":         PropertyUsageStoreIfNull,
	"PropertyUsageAnimateAsTrigger":    PropertyUsageAnimateAsTrigger,
	"PropertyUsageUpdateAllIfModified": PropertyUsageUpdateAllIfModified,
	"PropertyUsageDefault":             PropertyUsageDefault,
	"PropertyUsageDefaultIntl":         PropertyUsageDefaultIntl,
	"PropertyUsageNoEditor":            PropertyUsageNoEditor,
}

// CreateFunc will be called when we need to create a new instance of a class.
// When it is called, the Godot object will passed as an argument, as well as the
// methodData string, which is usually the name of the class to be created.
type CreateFunc func(Object, string) string

// CreateFuncRegistry is a mapping of instance creation functions. This map is
// used whenever a CreateFunc is registered. It is also used to look up a
// Creation function when Godot asks Go to create a new class instance.
var CreateFuncRegistry = map[string]CreateFunc{}

// DestroyFunc will be called when the object is destroyed. Takes the instance
// object, method data, user data. The method data is generally the class name,
// and the user data is generally the class instance id to destroy.
type DestroyFunc func(Object, string, string)

// DestroyFuncRegistry is a mapping of instance destroy functions. This map is
// used whenever a DestroyFunc is registered. It is also used to look up a
// Destroy function when Godot asks Go to destroy a class instance.
var DestroyFuncRegistry = map[string]DestroyFunc{}

// FreeFunc will be called when we should free the instance from memory. Takes
// in method data. The method data is generally the class name to be freed.
type FreeFunc func(string)

// SetPropertyFunc will be called when Godot requests to set a property on a given
// class. When it is called, the Godot object instance will be passed as an argument,
// as well as the methodData (which is usually the name of the class), the
// userData (which is generally the instance ID of the object), and the value
// to set.
type SetPropertyFunc func(Object, string, string, Variant)

// SetPropertyFuncRegistry is a mapping of instance property setters. This map is
// used whenever a SetPropertyFunc is registered. It is also used to look up
// a property setter function when Godot asks Go to set a property on an object.
var SetPropertyFuncRegistry = map[string]SetPropertyFunc{}

// GetPropertyFunc will be called when Godot requests a property on a given class
// instance. When it is called, Godot will pass the Godot object instance as an
// argument, as well as the methodData (which is usually the name of the class),
// and the userData (which is generally the instance ID of the object. You should
// return the property as a Variant.
type GetPropertyFunc func(Object, string, string) Variant

// GetPropertyFuncRegistry is a mapping of instance property getters. This map is
// used whenever a GetPropertyFunc is registered. It is also used to look up a
// property getter function when Godot asks Go to get a property on an object.
var GetPropertyFuncRegistry = map[string]GetPropertyFunc{}

// FreeFuncRegistry is a mapping of instance free functions. This map is used
// whenever a FreeFunc is registered. It is also used to look up a Free
// function when Godot asks Go to free a class instance.
var FreeFuncRegistry = map[string]FreeFunc{}

// MethodFunc will be called when a method attached to an instance is called.
// When it is called, it will be passed the Godot object the method is attached to,
// the methodData string (which is usually the class and method name that is being called),
// the userData string (which is usually the class instance ID),  the number of
// arguments being passed to the function, and a list of Variant arguments to pass
// to the function.
type MethodFunc func(Object, string, string, int, []Variant) Variant

// MethodFuncRegistry is a mapping of instance method functions. This map is
// used whenever a MethofFunc is registered. It is also used to look up a Method
// function when Godot asks Go to call a class method.
var MethodFuncRegistry = map[string]MethodFunc{}

// InstanceCreateFunc is a structure that contains the instance creation function
// that will be called when Godot asks Go to create a new instance of a class.
type InstanceCreateFunc struct {
	base       C.godot_instance_create_func
	CreateFunc CreateFunc
	MethodData string
	FreeFunc   FreeFunc
}

func (i *InstanceCreateFunc) getBase() C.godot_instance_create_func {
	return i.base
}

// InstanceDestroyFunc is a structure that contains the instance destruction function
// that will be called when Godot asks Go to destroy a class instance.
type InstanceDestroyFunc struct {
	base        C.godot_instance_destroy_func
	DestroyFunc DestroyFunc
	MethodData  string
	FreeFunc    FreeFunc
}

func (i *InstanceDestroyFunc) getBase() C.godot_instance_destroy_func {
	return i.base
}

// InstanceMethod is a structure that contains an instance method function that
// will be called when Godot asks Go to call a method on a class.
type InstanceMethod struct {
	base       C.godot_instance_method
	Method     MethodFunc
	MethodData string
	FreeFunc   FreeFunc
}

func (i *InstanceMethod) getBase() C.godot_instance_method {
	return i.base
}

type MethodAttributes struct {
	base    C.godot_method_attributes
	RPCType MethodRpcMode
}

func (m *MethodAttributes) getBase() C.godot_method_attributes {
	return m.base
}

// InstancePropertySet is a structure that contains an instance property setter
// function that will be called when Godot asks Go to set a property on a class.
type InstancePropertySet struct {
	base       C.godot_property_set_func
	SetFunc    SetPropertyFunc
	MethodData string
	FreeFunc   FreeFunc
}

func (i *InstancePropertySet) getBase() C.godot_property_set_func {
	return i.base
}

// InstancePropertyGet is a structure that contains an instance property getter
// function that will be called when Godot asks Go to get a property from a class
// instance.
type InstancePropertyGet struct {
	base       C.godot_property_get_func
	GetFunc    GetPropertyFunc
	MethodData string
	FreeFunc   FreeFunc
}

func (i *InstancePropertyGet) getBase() C.godot_property_get_func {
	return i.base
}

// NativeScript is a wrapper for the NativeScriptAPI.
var NativeScript = &nativeScript{}

// nativeScript is a structure that wraps the NativeScriptAPI methods.
type nativeScript struct {
	api *C.godot_gdnative_ext_nativescript_api_struct

	// Handle is a pointer to the gdnative handler. It must be passed to any
	// Godot nativescript functions. This will be populated when 'godot_nativescript_init'
	// is called by Godot upon script initialization.
	handle unsafe.Pointer
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
func (n *nativeScript) RegisterClass(name, base string, createFunc *InstanceCreateFunc, destroyFunc *InstanceDestroyFunc) {
	// Construct the C struct based on the Go struct wrappers
	createFunc.base.create_func = (C.create_func)(unsafe.Pointer(C.cgo_gateway_create_func))
	createFunc.base.method_data = unsafe.Pointer(C.CString(createFunc.MethodData))
	createFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))
	destroyFunc.base.destroy_func = (C.destroy_func)(unsafe.Pointer(C.cgo_gateway_destroy_func))
	destroyFunc.base.method_data = unsafe.Pointer(C.CString(destroyFunc.MethodData))
	destroyFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))

	// Register our Create and Destroy functions in a Go map, so the correct
	// function can be called when cgo_gateway_<type>_func is called.
	CreateFuncRegistry[createFunc.MethodData] = createFunc.CreateFunc
	FreeFuncRegistry[createFunc.MethodData] = createFunc.FreeFunc
	DestroyFuncRegistry[destroyFunc.MethodData] = destroyFunc.DestroyFunc
	FreeFuncRegistry[destroyFunc.MethodData] = destroyFunc.FreeFunc

	// Register the class with Godot.
	C.go_godot_nativescript_register_class(
		n.api,
		n.handle,
		C.CString(name),
		C.CString(base),
		createFunc.getBase(),
		destroyFunc.getBase(),
	)
}

// RegisterToolClass will register the given class with Godot as a tool. Refer to
// the 'RegisterClass' method for more information on how to use this.
func (n *nativeScript) RegisterToolClass(name, base string, createFunc *InstanceCreateFunc, destroyFunc *InstanceDestroyFunc) {
	// Construct the C struct based on the Go struct wrappers
	createFunc.base.create_func = (C.create_func)(unsafe.Pointer(C.cgo_gateway_create_func))
	createFunc.base.method_data = unsafe.Pointer(C.CString(createFunc.MethodData))
	createFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))
	destroyFunc.base.destroy_func = (C.destroy_func)(unsafe.Pointer(C.cgo_gateway_destroy_func))
	destroyFunc.base.method_data = unsafe.Pointer(C.CString(destroyFunc.MethodData))
	destroyFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))

	// Register our Create and Destroy functions in a Go map, so the correct
	// function can be called when cgo_gateway_<type>_func is called.
	CreateFuncRegistry[createFunc.MethodData] = createFunc.CreateFunc
	FreeFuncRegistry[createFunc.MethodData] = createFunc.FreeFunc
	DestroyFuncRegistry[destroyFunc.MethodData] = destroyFunc.DestroyFunc
	FreeFuncRegistry[destroyFunc.MethodData] = destroyFunc.FreeFunc

	// Register the class with Godot.
	C.go_godot_nativescript_register_tool_class(
		n.api,
		n.handle,
		C.CString(name),
		C.CString(base),
		createFunc.getBase(),
		destroyFunc.getBase(),
	)
}

// RegisterMethod will register the given method with Godot and associate it with
// the given class name. The name parameter is the name of the class you wish to
// attach this method to. The funcName parameter is the name of the function you
// want to register. The attributes and method are what will actually be called
// when Godot calls the method on the object.
func (n *nativeScript) RegisterMethod(name, funcName string, attributes *MethodAttributes, method *InstanceMethod) {
	// Construct the C struct based on the Go struct wrappers
	attributes.base.rpc_type = attributes.RPCType.getBase()
	method.base.method = (C.method)(unsafe.Pointer(C.cgo_gateway_method_func))
	method.base.method_data = unsafe.Pointer(C.CString(method.MethodData))
	method.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))

	// Register the Method function in a Go map, so the correct function can
	// be called when cgo_gateway_<type>_func is called.
	MethodFuncRegistry[method.MethodData] = method.Method
	FreeFuncRegistry[method.MethodData] = method.FreeFunc

	// Register the method with Godot.
	C.go_godot_nativescript_register_method(
		n.api,
		n.handle,
		C.CString(name),
		C.CString(funcName),
		attributes.getBase(),
		method.getBase(),
	)
}

// RegisterProperty will register the given property with Godot and associate it
// with the given class name. The name parameter is the name of the class you wish
// to attach this property to. The path is the name of the property itself. The
// attributes and setter/getter methods are what will be called when Godot gets
// or sets this property.
//func (n *nativeScript) RegisterProperty(name, path string, attributes *C.godot_property_attributes, setFunc C.godot_property_set_func, getFunc C.godot_property_get_func) {
func (n *nativeScript) RegisterProperty(name, path string, attributes *PropertyAttributes, setFunc *InstancePropertySet, getFunc *InstancePropertyGet) {
	// Construct the C struct based on the attributes Go wrapper
	var attr C.godot_property_attributes
	attributes.base = &attr
	attributes.base.rset_type = attributes.RsetType.getBase()
	attributes.base._type = attributes.Type.getBase()
	attributes.base.hint = attributes.Hint.getBase()
	attributes.base.hint_string = *(attributes.HintString.getBase())
	attributes.base.usage = attributes.Usage.getBase()
	attributes.base.default_value = *(attributes.DefaultValue.getBase())

	// Construct the C struct based on the setFunc Go wrapper
	setFunc.base.set_func = (C.set_property_func)(unsafe.Pointer(C.cgo_gateway_property_set_func))
	setFunc.base.method_data = unsafe.Pointer(C.CString(setFunc.MethodData))
	setFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))

	// Construct the C struct based on the getFunc Go wrapper
	getFunc.base.get_func = (C.get_property_func)(unsafe.Pointer(C.cgo_gateway_property_get_func))
	getFunc.base.method_data = unsafe.Pointer(C.CString(getFunc.MethodData))
	getFunc.base.free_func = (C.free_func)(unsafe.Pointer(C.cgo_gateway_free_func))

	// Register the set/get property functions in a Go map, so the correct function can
	// be called when cgo_gateway_<type>_func is called.
	SetPropertyFuncRegistry[setFunc.MethodData] = setFunc.SetFunc
	GetPropertyFuncRegistry[getFunc.MethodData] = getFunc.GetFunc
	FreeFuncRegistry[setFunc.MethodData] = setFunc.FreeFunc
	FreeFuncRegistry[getFunc.MethodData] = getFunc.FreeFunc

	// Register the property with Godot.
	C.go_godot_nativescript_register_property(
		n.api,
		n.handle,
		C.CString(name),
		C.CString(path),
		attributes.getBase(),
		setFunc.getBase(),
		getFunc.getBase(),
	)
}

// RegisterSignal will register the given signal with Godot.
func (n *nativeScript) RegisterSignal(name string, signal *Signal) {
	// Construct the C struct based on the signal Go wrapper
	var base C.godot_signal
	signal.base = &base
	signal.base.name = *(signal.Name.getBase())
	signal.base.num_args = C.int(signal.NumArgs.getBase())
	signal.base.num_default_args = C.int(signal.NumDefaultArgs.getBase())

	// Build the arguments
	argsArray := C.go_godot_signal_argument_build_array(C.int(signal.NumArgs))
	for i, arg := range signal.Args {
		var cArg C.godot_signal_argument
		cArg.name = *(arg.Name.getBase())
		cArg._type = arg.Type.getBase()
		cArg.default_value = *(arg.DefaultValue.getBase())
		cArg.hint = arg.Hint.getBase()
		cArg.hint_string = *(arg.HintString.getBase())
		cArg.usage = arg.Usage.getBase()

		C.go_godot_signal_argument_add_element(argsArray, &cArg, C.int(i))
	}
	signal.base.args = *argsArray

	// Build the default arguments
	variantArray := C.go_godot_variant_build_array(C.int(signal.NumDefaultArgs))
	for i, variant := range signal.DefaultArgs {
		C.go_godot_variant_add_element(variantArray, variant.getBase(), C.int(i))
	}
	signal.base.default_args = *variantArray

	// Register the signal with Godot.
	C.go_godot_nativescript_register_signal(
		n.api,
		n.handle,
		C.CString(name),
		signal.getBase(),
	)
}

// nativeScriptInit will be called when `godot_nativescript_init` is called by
// Godot. You can use `SetNativeScriptInit` to set the function that will be called
// when NativeScript initializes.
var nativeScriptInit = []func(){}

// SetNativeScriptInit will configure the given function to be called when
// `godot_nativescript_init` is called by Godot upon NativeScript initialization.
// This is used so you can define a function that will run to register all of the
// classes that you want exposed to Godot.
func SetNativeScriptInit(initFunc ...func()) {
	for _, init := range initFunc {
		nativeScriptInit = append(nativeScriptInit, init)
	}
}

/*------------------------------------------------------------------------------
//	  			Exported C Functions
//
//   The methods below are special C exported functions. They can be called by
//   Godot directly or by one of the C gateway functions defined in
//   nativescript.c.
//
//----------------------------------------------------------------------------*/

/** Script entry (Registering all the classes and stuff) **/
// godot_nativescript_init is the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes,
// etc. The `unsafe.Pointer` type is used to represent a void C pointer.
//export godot_nativescript_init
func godot_nativescript_init(hdl unsafe.Pointer) {
	if debug {
		log.Println("Initializing NativeScript")
	}
	NativeScript.handle = hdl

	// Call the user-defined nativeScriptInit function
	if nativeScriptInit == nil {
		err := "NativeScript initialization function was not set! Use `gdnative.SetNativeScriptInit` to define the function that will run to register classes."
		log.Println(err)
		Log.Error(err)

		return
	}

	// Loop through any defined nativeScriptInit methods and execute them.
	for _, init := range nativeScriptInit {
		init()
	}
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in nativescript.c. It will be ultimately called by
// Godot, where it will pass us the Godot object and the MethodData defined in
// CreateFunc. We will need to return UserData, which can be used to track the
// actual instance that was created.
//export go_create_func
func go_create_func(godotObject *C.godot_object, methodData unsafe.Pointer) unsafe.Pointer {
	// Convert the method data into a Go string.
	methodDataString := unsafeToGoString(methodData)
	if debug {
		log.Println("Create function called for:", methodDataString)
	}

	// Look up the creation function in our CreateFuncRegistry for the function
	// to call.
	constructor := CreateFuncRegistry[methodDataString]

	// Call the constructor and return the user data string. The user data
	// returned by the create func will be passed to the method function as
	// userData.
	userData := constructor(Object{base: godotObject}, methodDataString)

	return unsafe.Pointer(C.CString(userData))
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in nativescript.c. It will use the userData passed to it,
// which is a CString of the instance id, which we can use to delete the instance
// from the instance registry. This will make the instance available to be garbage
// collected.
//export go_destroy_func
func go_destroy_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer) {
	// Convert the method data and user data into a Go string
	methodDataString := unsafeToGoString(methodData)
	userDataString := unsafeToGoString(userData)
	if debug {
		log.Println("Destroy function called for:", methodDataString)
	}

	// Look up the destroy function in our DestroyFuncRegistry for the function
	// to call.
	destructor := DestroyFuncRegistry[methodDataString]

	// Call the destructor function. We pass the methodData and userData to
	// the destructor so it knows which class and instance to destroy.
	destructor(Object{base: godotObject}, methodDataString, userDataString)
}

//export go_free_func
func go_free_func(methodData unsafe.Pointer) {
	// Convert the method data and user data into a Go string
	methodDataString := unsafeToGoString(methodData)
	if debug {
		log.Println("Free function called for:", methodDataString)
	}

	// Look up the free function in our FreeFuncRegistry for the function
	// to call.
	freer := FreeFuncRegistry[methodDataString]

	// Call the free function. We pass the methodData to the free
	// function so it knows which clas to free.
	freer(methodDataString)
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in nativescript.c.
//export go_method_func
func go_method_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant) C.godot_variant {
	// Convert the method data and user data into a Go string
	methodDataString := unsafeToGoString(methodData)
	userDataString := unsafeToGoString(userData)

	// Create a slice of Variants for the arguments
	variantArgs := []Variant{}

	// Get the size of each godot_variant object pointer.
	if debug {
		log.Println("  Getting size of argument pointer")
	}
	size := unsafe.Sizeof(*args)

	// Panic if something's wrong.
	if int(numArgs) > 50 {
		panic("Too many arguments. Invalid method.")
	}

	// If we have arguments, append the first argument.
	if int(numArgs) > 0 {
		arg := *args
		// Loop through all our arguments.
		for i := 0; i < int(numArgs); i++ {
			// Convert the variant into a Go Variant
			variant := Variant{base: arg}

			// Append the variant to our list of variants
			variantArgs = append(variantArgs, variant)

			// Convert the pointer into a uintptr so we can perform artithmetic on it.
			arrayPtr := uintptr(unsafe.Pointer(args))

			// Add the size of the godot_variant pointer to our array pointer to get the position
			// of the next argument.
			arg = (*C.godot_variant)(unsafe.Pointer(arrayPtr + size))
		}
	}

	// Look up the method function in our MethodFuncRegistry for the function
	// to call.
	method := MethodFuncRegistry[methodDataString]

	// Call the method
	ret := method(Object{base: godotObject}, methodDataString, userDataString, int(numArgs), variantArgs)

	return *ret.getBase()
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in nativescript.c.
//export go_set_property_func
func go_set_property_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, property *C.godot_variant) {
	// Convert the method data and user data into a Go string
	methodDataString := unsafeToGoString(methodData)
	userDataString := unsafeToGoString(userData)

	// Convert the property into a Go variant
	variant := Variant{base: property}

	// Look up the set property function in our SetPropertyFuncRegistry for
	// the function to call.
	setFunc := SetPropertyFuncRegistry[methodDataString]

	// Call the method
	setFunc(Object{base: godotObject}, methodDataString, userDataString, variant)
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in nativescript.c.
//export go_get_property_func
func go_get_property_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer) C.godot_variant {
	// Convert the method data and user data into a Go string
	methodDataString := unsafeToGoString(methodData)
	userDataString := unsafeToGoString(userData)

	// Look up the get property function in our GetPropertyFuncRegistry for
	// the function to call.
	getFunc := GetPropertyFuncRegistry[methodDataString]

	// Call the method
	ret := getFunc(Object{base: godotObject}, methodDataString, userDataString)

	return *ret.getBase()
}
