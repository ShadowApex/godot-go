package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>

void **allocate_array(int);
void **allocate_array(int length) {
    void** array;
    array = (void**)malloc(length * sizeof(void*));

    return array;
}

void add_element(void**, void*, int);
void add_element(void **array, void* element, int index) {
    array[index] = element;
}
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
func (o *Object) callParentMethod(baseClass, methodName string, args []reflect.Value, returns string) reflect.Value {
	log.Println("Calling parent method!")

	// Convert the base class and method names to C strings.
	log.Println("  Using base class: ", baseClass)
	classCString := C.CString(baseClass)
	log.Println("  Using method name: ", methodName)
	methodCString := C.CString(methodName)

	// Get the Godot method bind pointer so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object:", o.owner)
	var methodBind *C.godot_method_bind
	methodBind = C.godot_method_bind_get_method(classCString, methodCString)
	log.Println("  Using method bind pointer: ", methodBind)

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
	log.Println("  Built variant arguments: ", variantArgs)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	//cArgsArray := C.allocate_array(C.int(len(variantArgs)))
	cArgsArray := C.malloc(C.size_t(len(variantArgs)) * C.size_t(unsafe.Sizeof(uintptr(0)))) // pointer to allocated memory

	// Loop through and add each argument to our C args array.
	for i, arg := range variantArgs {
		C.add_element(&cArgsArray, unsafe.Pointer(arg), C.int(i))
	}
	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	// TODO: We need to have the return type passed to us so we know how to convert
	// the return value to its correct type.
	log.Println("  Building return value.")
	var ret unsafe.Pointer
	switch returns {
	case "string":
		ret = unsafe.Pointer(C.CString(""))
	default:
		log.Fatal("Unknown return type specified.")
	}

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		unsafe.Pointer(o.owner),
		&cArgsArray,
		ret,
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	var retValue reflect.Value
	switch returns {
	case "string":
		gdString := (*C.godot_string)(ret)
		retValue = reflect.ValueOf(C.GoString(C.godot_string_c_str(gdString)))
	}

	// Return the converted variant.
	return retValue
}

type Node struct {
	Object
}

func (n *Node) BaseClass() string {
	return "Node"
}

func (n *Node) GetName() string {
	log.Println("Calling Node.GetName()!")
	ret := n.callParentMethod(n.BaseClass(), "get_name", []reflect.Value{}, "string")
	log.Println("Got return value!")
	value := ret.Interface().(string)
	log.Println("Converted return value into string: ", value)

	return value
}

// TODO: Get this working
func (n *Node) GetNode(path *NodePath) Class {
	log.Println("Calling Node.GetNode()!")
	ret := n.callParentMethod(n.BaseClass(), "_get_node", []reflect.Value{}, "string")
	value := ret.Interface().(Class)
	log.Println("Converted return value into: ", value)

	return value
}

type Node2D struct {
	Node
}

func (n *Node2D) BaseClass() string {
	return "Node2D"
}
