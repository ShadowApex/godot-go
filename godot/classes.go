package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

import (
	"log"
	"reflect"
	"unsafe"
)

// Class is an interface for any objects that can have Godot
// inheritance.
type Class interface {
	BaseClass() string
	SetOwner(object *C.godot_object)
}

// Object is the base structure for any Godot object.
type Object struct {
	owner *C.godot_object
}

// BaseClass will return a string of the Godot object's base class. This
// is used during class registration to register the correct base class.
func (o *Object) BaseClass() string {
	return "Object"
}

// SetOwner will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *Object) SetOwner(object *C.godot_object) {
	o.owner = object
}

// callParentMethod will call this object's method with the given method name.
func (o *Object) callParentMethod(methodName string, args []reflect.Value) reflect.Value {
	// Convert the base class and method names to C strings.
	classCString := C.CString(o.BaseClass())
	methodCString := C.CString(methodName)

	// Get the Godot method bind pointer so we can pass it to godot_method_bind_ptrcall.
	methodBind := C.godot_method_bind_get_method(classCString, methodCString)

	// Loop through the given arguments and see what type they are. When we know what
	// type it is, we need to convert them to godot_variant objects.
	// TODO: Probably pull this out into its own function?
	variantArgs := []*C.godot_variant{}
	for _, arg := range args {
		var argValue *C.godot_variant
		switch arg.Type().String() {
		case "bool":
			argValue = boolAsVariant(arg.Interface().(bool))
		case "int64":
			argValue = intAsVariant(arg.Interface().(int64))
		case "uint64":
			argValue = uIntAsVariant(arg.Interface().(uint64))
		case "float64":
			argValue = realAsVariant(arg.Interface().(float64))
		case "string":
			argValue = stringAsVariant(arg.Interface().(string))
		default:
			log.Fatal("Unknown type of argument value")
		}
		variantArgs = append(variantArgs, argValue)
	}

	// Construct a C array that will contain pointers to our arguments.
	//size := uintptr(len(variantArgs))
	cArgsArray := C.malloc(C.size_t(len(variantArgs)) * C.size_t(unsafe.Sizeof(uintptr(0)))) // pointer to allocated memory
	//if len(variantArgs) > 0 {
	//	arg := unsafe.Pointer(cArgsArray)
	//	for i := 0; i < len(variantArgs); i++ {
	//		// Write our argument
	//		arg = unsafe.Pointer(&variantArgs[i])

	//		// Convert the pointer into a uintptr so we can perform artithmetic on it.
	//		arrayPtr := uintptr(unsafe.Pointer(cArgsArray))

	//		// Add the size of the godot_variant pointer to our array pointer to get the position
	//		// of the next argument.
	//		arg = (*C.godot_variant)(unsafe.Pointer(arrayPtr + size))

	//	}
	//}

	// Construct our return object that will be populated by the method call.
	var ret C.godot_variant

	// Call the parent method. "ret" will be populated with the return value.
	C.godot_method_bind_ptrcall(
		methodBind,
		unsafe.Pointer(&o.owner),
		&cArgsArray,
		unsafe.Pointer(&ret),
	)

	// Get the return value type
	retType := C.godot_variant_get_type(&ret)

	// Convert the return value to a godot variant of some type.
	// Then convert the variant to a Go structure or builtin.
	// TODO: This is duplicated in godot.go. We should pull this out into its
	// own function.
	var retValue reflect.Value
	switch retType {
	case C.godot_variant_type(VariantTypeBool):
		retValue = reflect.ValueOf(variantAsBool(&ret))
	case C.godot_variant_type(VariantTypeInt):
		retValue = reflect.ValueOf(variantAsInt(&ret))
	case C.godot_variant_type(VariantTypeReal):
		retValue = reflect.ValueOf(variantAsReal(&ret))
	case C.godot_variant_type(VariantTypeString):
		retValue = reflect.ValueOf(variantAsString(&ret))
	default:
		log.Fatal("Unknown type of return value")
	}

	// Return the converted variant.
	return retValue

	// TODO: Does *p_ret give us back a godot_variant?
	//godot_method_bind GDAPI *godot_method_bind_get_method(const char *p_classname, const char *p_methodname);
	//void GDAPI godot_method_bind_ptrcall(godot_method_bind *p_method_bind, godot_object *p_instance, const void **p_args, void *p_ret);
}

type Node struct {
	Object
}

func (n *Node) BaseClass() string {
	return "Node"
}

func (n *Node) GetName() string {
	ret := n.callParentMethod("get_name", []reflect.Value{})
	value := ret.Interface().(string)

	return value
}

type Node2D struct {
	Node
}

func (n *Node2D) BaseClass() string {
	return "Node2D"
}
