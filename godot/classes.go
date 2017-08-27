package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>

void **build_array(int length);
void **build_array(int length) {
	void *ptr;
	void **arr = malloc(sizeof(void *) * length);
	for (int i = 0; i < length; i++) {
	    arr[i] = ptr;
	}

	return arr;
}

void add_element(void**, void*, int);
void add_element(void **array, void *element, int index) {
	printf("CGO: Array %p %p %p %p %p\n", &array, array, &array[index], *array, array[index]);
    array[index] = element;
	printf("CGO: Index %i %p\n", index, element);
	printf("CGO: Array %p %p %p %p %p\n", &array, array, &array[index], *array, array[index]);
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
	variantArgs := []unsafe.Pointer{} // void*
	for _, arg := range args {
		log.Println("  Argument type: ", arg.Type().String())
		var argValue unsafe.Pointer
		switch arg.Type().String() {
		case "bool":
			boolArg := C.godot_bool(arg.Interface().(bool))
			argValue = unsafe.Pointer(&boolArg)
		case "int64":
			var intArg C.int64_t
			intArg = C.int64_t(arg.Interface().(int64))
			argValue = unsafe.Pointer(&intArg)
		case "uint64":
			var intArg C.uint64_t
			intArg = C.uint64_t(arg.Interface().(uint64))
			argValue = unsafe.Pointer(&intArg)
		case "float64":
			floatArg := C.double(arg.Interface().(float64))
			argValue = unsafe.Pointer(&floatArg)
		case "string":
			stringArg := stringAsGodotString(arg.Interface().(string))
			argValue = unsafe.Pointer(stringArg)
		default:
			log.Fatal("Unknown type of argument value")
		}
		variantArgs = append(variantArgs, argValue)
	}
	log.Println("  Built variant arguments: ", variantArgs)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(len(variantArgs)))
	log.Println("    C Array: ", cArgsArray)

	// Loop through and add each argument to our C args array.
	for i, arg := range variantArgs {
		C.add_element(cArgsArray, arg, C.int(i))
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
	case "Node":
		ret = unsafe.Pointer(C.CString("")) // TODO: Make this correct?
	default:
		log.Fatal("Unknown return type specified.")
	}

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		unsafe.Pointer(o.owner),
		cArgsArray, // void**
		ret,        // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	var retValue reflect.Value
	switch returns {
	case "string":
		gdString := (*C.godot_string)(ret)
		retValue = reflect.ValueOf(C.GoString(C.godot_string_c_str(gdString)))
	case "Node":
		// TODO: What the fuck is going to be returned here?
		gdObject := (*C.godot_object)(ret)
		nodeObject := &Node{
			Object: Object{
				owner: gdObject,
			},
		}
		retValue = reflect.ValueOf(nodeObject)
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

func (n *Node) GetChild(index int64) *Node {
	log.Println("Calling Node.GetChild()!")
	ret := n.callParentMethod(n.BaseClass(), "get_child", []reflect.Value{reflect.ValueOf(index)}, "Node")
	log.Println("Got return value!")
	value := ret.Interface().(*Node)
	log.Println("Converted return value into string: ", value)

	return value
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
	ret := n.callParentMethod(n.BaseClass(), "_get_node", []reflect.Value{}, "NodePath")
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
