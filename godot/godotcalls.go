package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>

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
	"unsafe"
)

func getGodotMethod(baseClass string, methodName string) *C.godot_method_bind {
	// Convert the base class and method names to C strings.
	log.Println("  Using base class: ", baseClass)
	classCString := C.CString(baseClass)
	defer C.free(unsafe.Pointer(classCString))

	log.Println("  Using method name: ", methodName)
	methodCString := C.CString(methodName)
	defer C.free(unsafe.Pointer(methodCString))

	// Get the Godot method bind pointer so we can pass it to godot_method_bind_ptrcall.
	var methodBind *C.godot_method_bind
	methodBind = C.godot_method_bind_get_method(classCString, methodCString)
	log.Println("  Using method bind pointer: ", methodBind)
	return methodBind
}

func GodotCall_String_Dictionary(o Class, methodName string, arg0 *Dictionary) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.dictionary
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_Dictionary_Vector3_Vector3_Array_int(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Array, arg3 int64) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_dictionary
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{ret}

}
func GodotCall_void_RID_Rect2_Rect2_Color_bool_Object_bool(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Rect2, arg3 *Color, arg4 bool, arg5 *Object, arg6 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.godot_bool(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_String_Color(o Class, methodName string, arg0 string, arg1 string, arg2 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_Rect2_bool_Color_bool_Object(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 bool, arg3 *Color, arg4 bool, arg5 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_Object_int(o Class, methodName string, arg0 *Object, arg1 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_RID_RID_Vector3_RID_Vector3(o Class, methodName string, arg0 *RID, arg1 *Vector3, arg2 *RID, arg3 *Vector3) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_Vector3_RID(o Class, methodName string, arg0 *RID) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_PoolByteArray_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) *PoolByteArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_byte_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolByteArray{ret}

}
func GodotCall_Variant_String_Array(o Class, methodName string, arg0 string, arg1 *Array) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_Variant_NodePath(o Class, methodName string, arg0 *NodePath) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_Rect2_Color_bool(o Class, methodName string, arg0 *Rect2, arg1 *Color, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_float_bool_bool(o Class, methodName string, arg0 int64, arg1 float64, arg2 bool, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_float_float_bool(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_bool(o Class, methodName string, arg0 *Object, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector2_int_int(o Class, methodName string, arg0 int64, arg1 int64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_Vector2_Vector3(o Class, methodName string, arg0 *Vector3) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_void_int_Object_Transform2D_bool_Vector2(o Class, methodName string, arg0 int64, arg1 *Object, arg2 *Transform2D, arg3 bool, arg4 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_Color_bool(o Class, methodName string, arg0 *Object, arg1 *Color, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_int_Variant(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_float_RID(o Class, methodName string, arg0 *RID) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_void_int_Color(o Class, methodName string, arg0 int64, arg1 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Plane_int(o Class, methodName string, arg0 int64) *Plane {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_plane
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Plane{ret}

}
func GodotCall_void_int_Plane(o Class, methodName string, arg0 int64, arg1 *Plane) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.plane
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Variant(o Class, methodName string, arg0 *RID, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Vector3_Vector3(o Class, methodName string, arg0 *RID, arg1 *Vector3, arg2 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_RID_RID(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_Vector3(o Class, methodName string, arg0 *Vector3) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_Object_Vector2_Color_Object(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 *Color, arg3 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_PoolVector2Array_bool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_bool(o Class, methodName string, arg0 string, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Vector2_Variant_Object(o Class, methodName string, arg0 *Vector2, arg1 *Variant, arg2 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_RID_Rect2_int(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_PoolIntArray_int(o Class, methodName string, arg0 int64) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_int_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{ret}

}
func GodotCall_Object(o Class, methodName string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_Object_int_Vector2(o Class, methodName string, arg0 *Object, arg1 int64, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_Object_Transform_Object(o Class, methodName string, arg0 *Object, arg1 *Transform, arg2 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_String_String_Object_Object(o Class, methodName string, arg0 string, arg1 string, arg2 *Object, arg3 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Plane(o Class, methodName string, arg0 *Plane) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.plane
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_Object_String(o Class, methodName string, arg0 string, arg1 *Object, arg2 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Vector2_float_Color(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 float64, arg3 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_PoolStringArray(o Class, methodName string, arg0 string, arg1 *PoolStringArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Dictionary_bool(o Class, methodName string, arg0 bool) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_dictionary
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{ret}

}
func GodotCall_int_PoolByteArray(o Class, methodName string, arg0 *PoolByteArray) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Transform2D_Object(o Class, methodName string, arg0 *Object) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform2d
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{ret}

}
func GodotCall_void_RID_RID_Transform2D(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_Vector2_float(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Vector2, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_RID_RID_RID_int_int(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID, arg3 *RID, arg4 int64, arg5 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.int64_t(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_AABB(o Class, methodName string, arg0 *Object, arg1 *AABB) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.aabb
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector3_int_int_int(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_RID_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_bool_Object_String_Variant_Object_String_float_int_int_float(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant, arg3 *Object, arg4 string, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := stringAsGodotString(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.double(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.int64_t(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := C.int64_t(arg7)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	cArg8 := C.double(arg8)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg8), C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_RID_PoolIntArray_PoolVector2Array_PoolColorArray_PoolVector2Array_RID_int_RID(o Class, methodName string, arg0 *RID, arg1 *PoolIntArray, arg2 *PoolVector2Array, arg3 *PoolColorArray, arg4 *PoolVector2Array, arg5 *RID, arg6 int64, arg7 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.int64_t(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := arg7.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_int(o Class, methodName string, arg0 string, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolVector3Array(o Class, methodName string, arg0 *PoolVector3Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_String_Array(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_int_float(o Class, methodName string, arg0 int64, arg1 float64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_PoolVector2Array_Vector2_Vector2(o Class, methodName string, arg0 *Vector2, arg1 *Vector2) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{ret}

}
func GodotCall_void_Object_Object_int(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector2_int_int_Object_Vector2(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Object, arg3 *Vector2) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_void_Dictionary(o Class, methodName string, arg0 *Dictionary) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.dictionary
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Color_float(o Class, methodName string, arg0 float64) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_color
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{ret}

}
func GodotCall_void_RID_RID_int(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Vector2_float_Color(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_String_int_String_int(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_int_int_float_Vector3_Quat_Vector3(o Class, methodName string, arg0 int64, arg1 float64, arg2 *Vector3, arg3 *Quat, arg4 *Vector3) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.quat
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_PoolVector3Array_Object_bool(o Class, methodName string, arg0 *PoolVector3Array, arg1 *Object, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolStringArray_int(o Class, methodName string, arg0 *PoolStringArray, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector2(o Class, methodName string) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_void_Vector3(o Class, methodName string, arg0 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_PoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 *PoolByteArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_String_Variant(o Class, methodName string, arg0 string, arg1 string, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_int_int_int_int(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_RID_Vector2_Vector2_RID_RID(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *RID, arg3 *RID) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_void_RID_Vector3(o Class, methodName string, arg0 *RID, arg1 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Rect2_Object_int(o Class, methodName string, arg0 *Object, arg1 int64) *Rect2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rect2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Rect2{ret}

}
func GodotCall_int_String_String_int(o Class, methodName string, arg0 string, arg1 string, arg2 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_Vector2_Vector2_Vector2_int(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Vector2, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_PoolRealArray(o Class, methodName string, arg0 int64, arg1 *PoolRealArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolVector2Array_Color_PoolVector2Array_Object_Object_bool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *Color, arg2 *PoolVector2Array, arg3 *Object, arg4 *Object, arg5 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.godot_bool(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_Transform(o Class, methodName string, arg0 int64, arg1 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_String_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_int_Object_bool_String(o Class, methodName string, arg0 *Object, arg1 bool, arg2 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_int_int_Transform2D(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Object(o Class, methodName string, arg0 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_RID_float(o Class, methodName string, arg0 *RID, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_int_float(o Class, methodName string, arg0 *RID, arg1 int64, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_Object_Vector3(o Class, methodName string, arg0 *Object, arg1 *Vector3) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_void_String_String_int(o Class, methodName string, arg0 string, arg1 string, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Variant_Object_String(o Class, methodName string, arg0 *Object, arg1 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_String_int_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Array(o Class, methodName string, arg0 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_int_String_PoolStringArray_PoolByteArray(o Class, methodName string, arg0 int64, arg1 string, arg2 *PoolStringArray, arg3 *PoolByteArray) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_bool_String_Object_String(o Class, methodName string, arg0 string, arg1 *Object, arg2 string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_bool_RID_Transform2D_Vector2_float_Object(o Class, methodName string, arg0 *RID, arg1 *Transform2D, arg2 *Vector2, arg3 float64, arg4 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_bool_bool(o Class, methodName string, arg0 bool, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Vector2_Color_bool_Object(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Color, arg3 bool, arg4 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int(o Class, methodName string, arg0 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_String_String_int(o Class, methodName string, arg0 string, arg1 string, arg2 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Transform(o Class, methodName string) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{ret}

}
func GodotCall_bool_Object_Object(o Class, methodName string, arg0 *Object, arg1 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_Vector2_Vector2(o Class, methodName string, arg0 *Vector2, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_String_String(o Class, methodName string, arg0 string, arg1 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Variant_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_bool_String_int_String_int(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_PoolVector2Array_PoolColorArray_float_bool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Vector2_float_Vector2(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_String_int(o Class, methodName string, arg0 string, arg1 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_int_Variant(o Class, methodName string, arg0 *Variant) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_RID_bool(o Class, methodName string, arg0 *RID, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_float_float_float_bool(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 float64, arg4 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_int_int(o Class, methodName string, arg0 int64, arg1 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_NodePath(o Class, methodName string) *NodePath {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_node_path
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &NodePath{ret}

}
func GodotCall_void_PoolRealArray(o Class, methodName string, arg0 *PoolRealArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_float_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_void_int_int(o Class, methodName string, arg0 int64, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolStringArray_bool_String_int(o Class, methodName string, arg0 *PoolStringArray, arg1 bool, arg2 string, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector2_float_bool(o Class, methodName string, arg0 float64, arg1 bool) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_bool_Transform_Vector3(o Class, methodName string, arg0 *Transform, arg1 *Vector3) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_Object_Object_bool(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_String_int(o Class, methodName string, arg0 *Object, arg1 string, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_RID_RID(o Class, methodName string, arg0 *RID) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_RID_RID_Transform_RID_Transform(o Class, methodName string, arg0 *RID, arg1 *Transform, arg2 *RID, arg3 *Transform) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_void_int_int_Variant(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_float_float(o Class, methodName string, arg0 float64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_void_Object_Object_int_bool(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 int64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_int_Array_Array_int(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Array, arg3 *Array, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_int_RID(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Transform_RID(o Class, methodName string, arg0 *RID) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{ret}

}
func GodotCall_void_int_int_Rect2_Vector2_float(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Rect2, arg3 *Vector2, arg4 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.double(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_RID(o Class, methodName string, arg0 *RID) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_float_float_float(o Class, methodName string, arg0 float64, arg1 float64, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Object_NodePath_Object_NodePath_Variant_float_int_int_float(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Object, arg3 *NodePath, arg4 *Variant, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.double(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.int64_t(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := C.int64_t(arg7)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	cArg8 := C.double(arg8)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg8), C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_bool_Vector2_Vector2(o Class, methodName string, arg0 bool, arg1 *Vector2, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_bool_Rect2(o Class, methodName string, arg0 *RID, arg1 bool, arg2 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_NodePath(o Class, methodName string, arg0 *NodePath) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_void_RID_int_Transform(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Variant_int_String_String(o Class, methodName string, arg0 int64, arg1 string, arg2 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_PoolVector3Array_int_int(o Class, methodName string, arg0 int64, arg1 int64) *PoolVector3Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector3_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector3Array{ret}

}
func GodotCall_PoolIntArray(o Class, methodName string) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_int_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{ret}

}
func GodotCall_bool_Transform2D_Object_Transform2D(o Class, methodName string, arg0 *Transform2D, arg1 *Object, arg2 *Transform2D) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_Rect2_bool(o Class, methodName string, arg0 *Rect2, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_String_Variant(o Class, methodName string, arg0 int64, arg1 string, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_float_float_float(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_PoolVector2Array(o Class, methodName string, arg0 int64, arg1 *PoolVector2Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_PoolVector2Array_PoolColorArray_float_bool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 float64, arg4 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String(o Class, methodName string, arg0 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_PoolVector3Array(o Class, methodName string) *PoolVector3Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector3_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector3Array{ret}

}
func GodotCall_int_Object(o Class, methodName string, arg0 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_bool_String_String(o Class, methodName string, arg0 string, arg1 string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Dictionary_String(o Class, methodName string, arg0 string) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_dictionary
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{ret}

}
func GodotCall_Vector3_Vector3(o Class, methodName string, arg0 *Vector3) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_void_Vector3_Vector3_Vector3_int(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Vector3, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_String_int_int_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64, arg4 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Variant_int_int(o Class, methodName string, arg0 int64, arg1 int64) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_Object_Vector2(o Class, methodName string, arg0 *Vector2) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_int_String_int_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Transform2D_int_int(o Class, methodName string, arg0 int64, arg1 int64) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform2d
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{ret}

}
func GodotCall_AABB_RID(o Class, methodName string, arg0 *RID) *AABB {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_aabb
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &AABB{ret}

}
func GodotCall_void_int_float_Variant_float(o Class, methodName string, arg0 int64, arg1 float64, arg2 *Variant, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_PoolStringArray(o Class, methodName string) *PoolStringArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_string_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolStringArray{ret}

}
func GodotCall_void_RID_PoolVector2Array_PoolColorArray_PoolVector2Array_RID_float_RID(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *RID, arg5 float64, arg6 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.double(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := arg6.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_String_RID(o Class, methodName string, arg0 *RID, arg1 string, arg2 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Basis(o Class, methodName string) *Basis {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_basis
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Basis{ret}

}
func GodotCall_void_Object_Object_Vector3_Vector3_int(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 *Vector3, arg3 *Vector3, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Dictionary(o Class, methodName string) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_dictionary
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{ret}

}
func GodotCall_Vector3_Vector2(o Class, methodName string, arg0 *Vector2) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_PoolRealArray(o Class, methodName string) *PoolRealArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_real_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolRealArray{ret}

}
func GodotCall_Object_NodePath(o Class, methodName string, arg0 *NodePath) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_RID_Rect2_RID_bool_Color_bool_RID(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *RID, arg3 bool, arg4 *Color, arg5 bool, arg6 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.godot_bool(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := arg6.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_AABB(o Class, methodName string, arg0 *RID, arg1 *AABB) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.aabb
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector3_int(o Class, methodName string, arg0 int64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_void_int_RID_int_int_int(o Class, methodName string, arg0 int64, arg1 *RID, arg2 int64, arg3 int64, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector2_String_int(o Class, methodName string, arg0 string, arg1 int64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_void_String_int_Vector2(o Class, methodName string, arg0 string, arg1 int64, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_Object_String_Variant(o Class, methodName string, arg0 string, arg1 *Object, arg2 string, arg3 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_Object(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_String_int_int(o Class, methodName string, arg0 *Object, arg1 string, arg2 int64, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Transform2D_Vector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_String_bool(o Class, methodName string, arg0 *Object, arg1 string, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Dictionary_Object(o Class, methodName string, arg0 *Object) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_dictionary
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{ret}

}
func GodotCall_String_Variant(o Class, methodName string, arg0 *Variant) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_String_String_Object(o Class, methodName string, arg0 string, arg1 *Object) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_void_int_Object_int(o Class, methodName string, arg0 int64, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector3_Vector3_Vector3_float_int_float(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 float64, arg3 int64, arg4 float64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.double(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_int_int_int_int(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_String_String_String_bool_int_int(o Class, methodName string, arg0 string, arg1 string, arg2 bool, arg3 int64, arg4 int64) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_Dictionary_int(o Class, methodName string, arg0 int64) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_dictionary
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{ret}

}
func GodotCall_void_String_Variant(o Class, methodName string, arg0 string, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_Object_Vector2(o Class, methodName string, arg0 int64, arg1 *Object, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_String_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_String_Object_bool(o Class, methodName string, arg0 string, arg1 *Object, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_int_String(o Class, methodName string, arg0 int64, arg1 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Object_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_float_String(o Class, methodName string, arg0 string) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_PoolVector2Array(o Class, methodName string) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{ret}

}
func GodotCall_void_int_int_bool(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_String_PoolStringArray_bool_int_String(o Class, methodName string, arg0 string, arg1 *PoolStringArray, arg2 bool, arg3 int64, arg4 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := stringAsGodotString(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_Object_Vector2_String_Color_int(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 string, arg3 *Color, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_PoolStringArray_String(o Class, methodName string, arg0 string) *PoolStringArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_string_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolStringArray{ret}

}
func GodotCall_Vector2_Vector2(o Class, methodName string, arg0 *Vector2) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_bool_Vector2_Variant(o Class, methodName string, arg0 *Vector2, arg1 *Variant) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_Object_Object_String_Variant(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 string, arg3 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Vector2_String_Color_int(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 string, arg3 *Color, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_Color(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_String_Object(o Class, methodName string, arg0 string, arg1 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_String_float_float_bool(o Class, methodName string, arg0 string, arg1 float64, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_int_int(o Class, methodName string, arg0 int64, arg1 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_String_int_String_int(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_Object(o Class, methodName string, arg0 *Object) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_Vector2_float(o Class, methodName string, arg0 float64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_Vector2_int(o Class, methodName string, arg0 int64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_void_RID_int_bool(o Class, methodName string, arg0 *RID, arg1 int64, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_Rect2_Rect2_Color_bool_Object_bool(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 *Rect2, arg3 *Color, arg4 bool, arg5 *Object, arg6 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.godot_bool(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_PoolStringArray_int(o Class, methodName string, arg0 int64) *PoolStringArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_string_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolStringArray{ret}

}
func GodotCall_void_int_Object_String(o Class, methodName string, arg0 int64, arg1 *Object, arg2 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_Vector3_float(o Class, methodName string, arg0 int64, arg1 *Vector3, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_String_String(o Class, methodName string, arg0 string) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_float_int(o Class, methodName string, arg0 int64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_Variant_Array(o Class, methodName string, arg0 *Array) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_Object_String_int(o Class, methodName string, arg0 string, arg1 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_bool_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Object_float_bool(o Class, methodName string, arg0 float64, arg1 bool) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_Object_Rect2(o Class, methodName string, arg0 *Rect2) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_int_Vector2_bool(o Class, methodName string, arg0 *Vector2, arg1 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_PoolByteArray(o Class, methodName string) *PoolByteArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_byte_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolByteArray{ret}

}
func GodotCall_int_Vector2(o Class, methodName string, arg0 *Vector2) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_NodePath_Variant(o Class, methodName string, arg0 *NodePath, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Object_NodePath_Variant_Variant_float_int_int_float(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Variant, arg3 *Variant, arg4 float64, arg5 int64, arg6 int64, arg7 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.double(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.int64_t(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.int64_t(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := C.double(arg7)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_Object(o Class, methodName string, arg0 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_Array(o Class, methodName string, arg0 int64, arg1 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_RID_Vector2_RID_RID(o Class, methodName string, arg0 *Vector2, arg1 *RID, arg2 *RID) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_void_String_int_int_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Rect2_Rect2_RID_Vector2_Vector2_int_int_bool_Color_RID(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Rect2, arg3 *RID, arg4 *Vector2, arg5 *Vector2, arg6 int64, arg7 int64, arg8 bool, arg9 *Color, arg10 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(11))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.int64_t(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := C.int64_t(arg7)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	cArg8 := C.godot_bool(arg8)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg8), C.int(8))

	cArg9 := arg9.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg9), C.int(9))

	cArg10 := arg10.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg10), C.int(10))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Object_int(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_float_bool(o Class, methodName string, arg0 float64, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Transform_int(o Class, methodName string, arg0 int64) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{ret}

}
func GodotCall_void_int_int_bool_int_PoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool, arg3 int64, arg4 *PoolByteArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_RID(o Class, methodName string, arg0 *RID) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_void_PoolVector2Array(o Class, methodName string, arg0 *PoolVector2Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector3_int_float(o Class, methodName string, arg0 int64, arg1 float64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_PoolVector2Array_int_float(o Class, methodName string, arg0 int64, arg1 float64) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{ret}

}
func GodotCall_void_int_Variant_bool(o Class, methodName string, arg0 int64, arg1 *Variant, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Transform2D_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform2d
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{ret}

}
func GodotCall_bool_Transform2D_Vector2_Object_Transform2D_Vector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2, arg2 *Object, arg3 *Transform2D, arg4 *Vector2) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_PoolVector2Array_PoolColorArray_PoolVector2Array_Object_Object_bool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 *PoolVector2Array, arg3 *Object, arg4 *Object, arg5 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.godot_bool(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Variant_Object(o Class, methodName string, arg0 *Variant, arg1 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_Object_int(o Class, methodName string, arg0 *Object, arg1 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_void_Vector3_Vector3_Vector3(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_Variant(o Class, methodName string, arg0 int64, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_Object(o Class, methodName string, arg0 *Object) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_Transform_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{ret}

}
func GodotCall_void_Vector2_int_bool_bool_bool(o Class, methodName string, arg0 *Vector2, arg1 int64, arg2 bool, arg3 bool, arg4 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_String(o Class, methodName string) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_Object_String(o Class, methodName string, arg0 string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_String_String_Variant(o Class, methodName string, arg0 int64, arg1 string, arg2 string, arg3 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_RID_int_int_float(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_AABB_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) *AABB {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_aabb
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &AABB{ret}

}
func GodotCall_String_RID(o Class, methodName string, arg0 *RID) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_int_NodePath(o Class, methodName string, arg0 *NodePath) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Rect2_int(o Class, methodName string, arg0 int64) *Rect2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rect2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Rect2{ret}

}
func GodotCall_void_int_int_int_int(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_bool_int(o Class, methodName string, arg0 int64, arg1 bool, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Vector2_float_Object(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_RID_PoolVector2Array_PoolColorArray_PoolVector2Array_RID_RID_bool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *RID, arg5 *RID, arg6 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.godot_bool(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_float(o Class, methodName string) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_bool_String(o Class, methodName string, arg0 string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Array_String(o Class, methodName string, arg0 string) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_void_RID_PoolVector2Array(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Vector3_Vector3(o Class, methodName string, arg0 *Vector3, arg1 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_float_Vector2(o Class, methodName string, arg0 *Vector2) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_Object_RID(o Class, methodName string, arg0 *RID) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_String_String_Color_bool(o Class, methodName string, arg0 string, arg1 string, arg2 *Color, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_String_Object(o Class, methodName string, arg0 string, arg1 string, arg2 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Color(o Class, methodName string, arg0 *RID, arg1 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_Object_Rect2_Vector2(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 *Rect2, arg3 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_bool_int(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Plane(o Class, methodName string) *Plane {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_plane
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Plane{ret}

}
func GodotCall_int_bool(o Class, methodName string, arg0 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_int_int_float_bool(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_AABB(o Class, methodName string, arg0 *AABB) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.aabb
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Dictionary_Vector2_Vector2_Array_int(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Array, arg3 int64) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_dictionary
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{ret}

}
func GodotCall_Vector3_float_bool(o Class, methodName string, arg0 float64, arg1 bool) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_void_String_Variant_bool(o Class, methodName string, arg0 string, arg1 *Variant, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_Object_Variant(o Class, methodName string, arg0 string, arg1 *Object, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_bool_RID(o Class, methodName string, arg0 *Object, arg1 bool, arg2 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_String_bool_bool(o Class, methodName string, arg0 string, arg1 bool, arg2 bool) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_RID_int_int_float(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Transform_bool(o Class, methodName string, arg0 bool) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{ret}

}
func GodotCall_PoolVector3Array_int_float(o Class, methodName string, arg0 int64, arg1 float64) *PoolVector3Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector3_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector3Array{ret}

}
func GodotCall_void_float_Color(o Class, methodName string, arg0 float64, arg1 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_Object_bool(o Class, methodName string, arg0 int64, arg1 *Object, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Rect2_RID_Rect2_Color_bool_RID_bool(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *RID, arg3 *Rect2, arg4 *Color, arg5 bool, arg6 *RID, arg7 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.godot_bool(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := arg6.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := C.godot_bool(arg7)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_String_Object(o Class, methodName string, arg0 *Object) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_void_int_int_int(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_NodePath(o Class, methodName string, arg0 *NodePath) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_Vector3_float(o Class, methodName string, arg0 *Vector3, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_PoolStringArray_PoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 *PoolStringArray, arg3 *PoolByteArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_Transform2D_Vector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_PoolColorArray(o Class, methodName string) *PoolColorArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_color_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolColorArray{ret}

}
func GodotCall_Vector2_String(o Class, methodName string, arg0 string) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_Rect2(o Class, methodName string) *Rect2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rect2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Rect2{ret}

}
func GodotCall_void_Object_Rect2_Vector2(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Vector3(o Class, methodName string, arg0 *Vector3) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_PoolVector2Array_Color_float_bool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *Color, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolColorArray(o Class, methodName string, arg0 *PoolColorArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_RID_int_bool(o Class, methodName string, arg0 int64, arg1 bool) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_Variant_Array_Array_int_Array(o Class, methodName string, arg0 *Array, arg1 *Array, arg2 int64, arg3 *Array) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_int_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Array_RID_int(o Class, methodName string, arg0 *RID, arg1 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_int_int(o Class, methodName string, arg0 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Color(o Class, methodName string) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_color
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{ret}

}
func GodotCall_void_RID_int_int_int_int(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_Color(o Class, methodName string, arg0 string, arg1 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_float(o Class, methodName string, arg0 float64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_int_Array_Array_int(o Class, methodName string, arg0 int64, arg1 *Array, arg2 *Array, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolStringArray(o Class, methodName string, arg0 *PoolStringArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Object_String_Object_String_Variant_float_int_int_float(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Object, arg3 string, arg4 *Variant, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := stringAsGodotString(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.double(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.int64_t(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := C.int64_t(arg7)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	cArg8 := C.double(arg8)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg8), C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_bool_String_int_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Array_Array_int(o Class, methodName string, arg0 *Array, arg1 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_void_bool_bool_int_int(o Class, methodName string, arg0 bool, arg1 bool, arg2 int64, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_String_int_bool_bool(o Class, methodName string, arg0 string, arg1 int64, arg2 bool, arg3 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_bool_Object_String(o Class, methodName string, arg0 *Object, arg1 string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Variant_int_String(o Class, methodName string, arg0 int64, arg1 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_RID_int_int_bool(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Rect2(o Class, methodName string, arg0 *RID, arg1 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_String(o Class, methodName string, arg0 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_int_Vector2_float_float_int_int(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 float64, arg3 int64, arg4 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_String_String_int(o Class, methodName string, arg0 string, arg1 int64) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_void_Object_Rect2(o Class, methodName string, arg0 *Object, arg1 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Variant_Transform2D_Vector2_Object_Transform2D_Vector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2, arg2 *Object, arg3 *Transform2D, arg4 *Vector2) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_String_Dictionary(o Class, methodName string, arg0 string, arg1 *Dictionary) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.dictionary
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_NodePath(o Class, methodName string, arg0 int64, arg1 *NodePath) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_NodePath(o Class, methodName string, arg0 *NodePath) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Transform2D_Vector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_PoolVector2Array_int(o Class, methodName string, arg0 *PoolVector2Array, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_String_int(o Class, methodName string, arg0 string, arg1 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_Variant(o Class, methodName string, arg0 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Object_String_Variant(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 string, arg3 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Vector2_bool(o Class, methodName string, arg0 *Vector2, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_String_String(o Class, methodName string, arg0 string, arg1 string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_PoolVector2Array_Vector2_Vector2_bool(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 bool) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{ret}

}
func GodotCall_void_Object_String_PoolStringArray(o Class, methodName string, arg0 *Object, arg1 string, arg2 *PoolStringArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_Color_bool(o Class, methodName string, arg0 int64, arg1 *Color, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_bool(o Class, methodName string, arg0 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector3_float(o Class, methodName string, arg0 float64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_void_RID_Vector2_Vector2_Color_float_bool(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Vector2, arg3 *Color, arg4 float64, arg5 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.double(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.godot_bool(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_String_Dictionary(o Class, methodName string, arg0 string, arg1 *Dictionary) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.dictionary
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Vector2_Vector2_Vector2_float_int_float(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 float64, arg3 int64, arg4 float64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.double(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_int_Object_Transform2D_Object(o Class, methodName string, arg0 *Object, arg1 *Transform2D, arg2 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_PoolVector2Array_int(o Class, methodName string, arg0 int64) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{ret}

}
func GodotCall_void_String_Object_int_String_Variant(o Class, methodName string, arg0 string, arg1 *Object, arg2 int64, arg3 string, arg4 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := stringAsGodotString(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID(o Class, methodName string, arg0 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_String_PoolStringArray(o Class, methodName string, arg0 string, arg1 string, arg2 *PoolStringArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Variant_String(o Class, methodName string, arg0 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_int_Rect2(o Class, methodName string, arg0 int64, arg1 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_Vector2_Object(o Class, methodName string, arg0 *Vector2, arg1 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Vector3_Vector3_Vector3_bool(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 bool) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_bool_int(o Class, methodName string, arg0 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_bool_Vector2(o Class, methodName string, arg0 *Vector2) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Object_Vector3(o Class, methodName string, arg0 *Vector3) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_int_Object_int(o Class, methodName string, arg0 *Object, arg1 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_Basis(o Class, methodName string, arg0 *Basis) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.basis
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Vector2(o Class, methodName string, arg0 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_int_String_PoolStringArray_String(o Class, methodName string, arg0 int64, arg1 string, arg2 *PoolStringArray, arg3 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := stringAsGodotString(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_bool_String_Variant(o Class, methodName string, arg0 string, arg1 *Variant) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_int_String_int(o Class, methodName string, arg0 int64, arg1 string, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_PoolIntArray_int_float_float(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_int_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{ret}

}
func GodotCall_void_String_PoolByteArray_bool(o Class, methodName string, arg0 string, arg1 *PoolByteArray, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_int_bool_bool_bool_Vector2(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 bool, arg4 bool, arg5 bool, arg6 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.godot_bool(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := arg6.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_RID_Object_int(o Class, methodName string, arg0 *Object, arg1 int64) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_void_int_bool(o Class, methodName string, arg0 int64, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolVector3Array_PoolVector2Array_PoolColorArray_PoolVector2Array_PoolVector3Array_Array(o Class, methodName string, arg0 *PoolVector3Array, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *PoolVector3Array, arg5 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_int_Object(o Class, methodName string, arg0 string, arg1 int64, arg2 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_String(o Class, methodName string, arg0 *RID, arg1 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_String_Object(o Class, methodName string, arg0 string, arg1 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_String_int(o Class, methodName string, arg0 int64) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_void_int_PoolIntArray(o Class, methodName string, arg0 int64, arg1 *PoolIntArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_RID(o Class, methodName string, arg0 *RID, arg1 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_RID_int_int(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Array_PoolByteArray(o Class, methodName string, arg0 *PoolByteArray) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_bool_int_int(o Class, methodName string, arg0 int64, arg1 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Transform2D(o Class, methodName string) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform2d
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{ret}

}
func GodotCall_bool_bool(o Class, methodName string, arg0 bool) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_RID_int_int(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_float_int_int(o Class, methodName string, arg0 int64, arg1 int64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_Color_String_String(o Class, methodName string, arg0 string, arg1 string) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_color
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{ret}

}
func GodotCall_NodePath_int(o Class, methodName string, arg0 int64) *NodePath {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_node_path
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &NodePath{ret}

}
func GodotCall_float_String_String(o Class, methodName string, arg0 string, arg1 string) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_void_int_String(o Class, methodName string, arg0 int64, arg1 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_NodePath_bool(o Class, methodName string, arg0 string, arg1 *NodePath, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Transform2D(o Class, methodName string, arg0 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_RID(o Class, methodName string, arg0 *RID) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Vector3(o Class, methodName string) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector3
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{ret}

}
func GodotCall_Object_int(o Class, methodName string, arg0 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_void_RID_Rect2_bool_Color_bool_Object(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 bool, arg3 *Color, arg4 bool, arg5 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolByteArray(o Class, methodName string, arg0 *PoolByteArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_Vector2_int_Array_int(o Class, methodName string, arg0 *Vector2, arg1 int64, arg2 *Array, arg3 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_void_PoolVector3Array_bool_bool(o Class, methodName string, arg0 *PoolVector3Array, arg1 bool, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Variant(o Class, methodName string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_bool_float(o Class, methodName string, arg0 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_String_Array(o Class, methodName string, arg0 string, arg1 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_float_RID_int_int(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_Array(o Class, methodName string) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_PoolIntArray_int_int(o Class, methodName string, arg0 int64, arg1 int64) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_int_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{ret}

}
func GodotCall_float_Object_Vector2_String_String_Color(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 string, arg3 string, arg4 *Color) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := stringAsGodotString(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_void_Color(o Class, methodName string, arg0 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_String_Vector2(o Class, methodName string, arg0 *Vector2) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_float_RID_Vector2_int_int_Color(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 int64, arg3 int64, arg4 *Color) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_real
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(*ret)

}
func GodotCall_void_Object_int(o Class, methodName string, arg0 *Object, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Object_String(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_RID_Vector2_Vector2_Vector2_RID_RID(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Vector2, arg3 *RID, arg4 *RID) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_void_int_float(o Class, methodName string, arg0 int64, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_String_float(o Class, methodName string, arg0 string, arg1 string, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Rect2_Color(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_NodePath_int_bool(o Class, methodName string, arg0 int64, arg1 bool) *NodePath {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_node_path
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &NodePath{ret}

}
func GodotCall_void_int_bool_bool(o Class, methodName string, arg0 int64, arg1 bool, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_float_float_float_float(o Class, methodName string, arg0 float64, arg1 float64, arg2 float64, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_float(o Class, methodName string, arg0 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.double(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_RID(o Class, methodName string) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_PoolVector3Array_Vector3_Vector3_bool(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 bool) *PoolVector3Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_vector3_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector3Array{ret}

}
func GodotCall_int_String_Object_String_Array_int(o Class, methodName string, arg0 string, arg1 *Object, arg2 string, arg3 *Array, arg4 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_String_Vector2(o Class, methodName string, arg0 string, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_int_float(o Class, methodName string, arg0 int64, arg1 float64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_int_Object(o Class, methodName string, arg0 int64, arg1 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_float(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_float(o Class, methodName string, arg0 string, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_bool_Object_float_String_Variant_Variant_Variant_Variant_Variant(o Class, methodName string, arg0 *Object, arg1 float64, arg2 string, arg3 *Variant, arg4 *Variant, arg5 *Variant, arg6 *Variant, arg7 *Variant) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := arg6.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := arg7.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Variant_int(o Class, methodName string, arg0 int64) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_bool_float(o Class, methodName string, arg0 bool, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Vector2_Variant_Object(o Class, methodName string, arg0 *Vector2, arg1 *Variant, arg2 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int(o Class, methodName string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Vector2_int_float(o Class, methodName string, arg0 int64, arg1 float64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_Variant_String_String_Variant(o Class, methodName string, arg0 string, arg1 string, arg2 *Variant) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_Variant_RID_String(o Class, methodName string, arg0 *RID, arg1 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_bool(o Class, methodName string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_String_int_bool(o Class, methodName string, arg0 string, arg1 int64, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_int_float_bool(o Class, methodName string, arg0 int64, arg1 float64, arg2 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_int_String_String_Dictionary_Array_Array(o Class, methodName string, arg0 string, arg1 string, arg2 *Dictionary, arg3 *Array, arg4 *Array) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.dictionary
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_int_int_int(o Class, methodName string, arg0 int64, arg1 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Color_int_int(o Class, methodName string, arg0 int64, arg1 int64) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_color
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{ret}

}
func GodotCall_void_Object_int_Transform(o Class, methodName string, arg0 *Object, arg1 int64, arg2 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_Object_int_bool_String(o Class, methodName string, arg0 int64, arg1 *Object, arg2 int64, arg3 bool, arg4 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.godot_bool(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := stringAsGodotString(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_String_bool_String(o Class, methodName string, arg0 string, arg1 bool, arg2 string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_Variant_Variant(o Class, methodName string, arg0 *Variant) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_int_bool_String_String(o Class, methodName string, arg0 int64, arg1 bool, arg2 string, arg3 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := stringAsGodotString(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolVector2Array_PoolColorArray_PoolVector2Array_Object_float_Object(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 *PoolVector2Array, arg3 *Object, arg4 float64, arg5 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.double(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := arg5.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_int_int_int_int(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.int64_t(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_int_Transform2D(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_RID_Transform(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_int_String(o Class, methodName string, arg0 string, arg1 int64, arg2 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := stringAsGodotString(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Rect2(o Class, methodName string, arg0 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_NodePath_Object_int(o Class, methodName string, arg0 *NodePath, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_int_int_int_float(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.double(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Variant_String_bool(o Class, methodName string, arg0 string, arg1 bool) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_Transform(o Class, methodName string, arg0 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Vector2_Vector2_Color_float_bool(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Color, arg3 float64, arg4 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.double(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_Array_bool(o Class, methodName string, arg0 string, arg1 *Array, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_int_bool(o Class, methodName string, arg0 *Object, arg1 int64, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.godot_bool(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Vector2_Vector2_bool(o Class, methodName string, arg0 *Vector2, arg1 bool) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_vector2
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{ret}

}
func GodotCall_void_String_String(o Class, methodName string, arg0 string, arg1 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Color_int(o Class, methodName string, arg0 int64) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_color
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{ret}

}
func GodotCall_Transform2D_int(o Class, methodName string, arg0 int64) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform2d
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{ret}

}
func GodotCall_void_Object_String(o Class, methodName string, arg0 *Object, arg1 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_RID_int(o Class, methodName string, arg0 int64) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_int_int_String_int(o Class, methodName string, arg0 int64, arg1 string, arg2 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_Variant_RID(o Class, methodName string, arg0 *RID) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void(o Class, methodName string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_AABB(o Class, methodName string) *AABB {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_aabb
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &AABB{ret}

}
func GodotCall_PoolRealArray_int(o Class, methodName string, arg0 int64) *PoolRealArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_real_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolRealArray{ret}

}
func GodotCall_void_RID_Transform(o Class, methodName string, arg0 *RID, arg1 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.transform
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Variant_String_String(o Class, methodName string, arg0 string, arg1 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_void_String_Object_int(o Class, methodName string, arg0 string, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_PoolIntArray_String_int_int_int(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := C.int64_t(arg3)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_pool_int_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{ret}

}
func GodotCall_void_int_Vector2(o Class, methodName string, arg0 int64, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Object_float(o Class, methodName string, arg0 *Object, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.double(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Object_bool(o Class, methodName string, arg0 bool) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.godot_bool(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_bool_Object_String_Variant_Variant_float_int_int_float(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant, arg3 *Variant, arg4 float64, arg5 int64, arg6 int64, arg7 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.double(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.int64_t(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.int64_t(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := C.double(arg7)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_Object_Object_String(o Class, methodName string, arg0 *Object, arg1 string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_object
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{ret}

}
func GodotCall_Variant_String_String_Array(o Class, methodName string, arg0 string, arg1 string, arg2 *Array) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_bool_Vector2_Rect2(o Class, methodName string, arg0 *Vector2, arg1 *Rect2) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rect2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_void_int_Vector3(o Class, methodName string, arg0 int64, arg1 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector3
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Array_int(o Class, methodName string, arg0 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_array
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{ret}

}
func GodotCall_Variant_Vector2_Object(o Class, methodName string, arg0 *Vector2, arg1 *Object) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_NodePath_Object(o Class, methodName string, arg0 *Object) *NodePath {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_node_path
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &NodePath{ret}

}
func GodotCall_void_RID_Vector2(o Class, methodName string, arg0 *RID, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_Vector2_Variant(o Class, methodName string, arg0 *Vector2, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_RID_RID_RID(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID, arg3 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_String_Variant(o Class, methodName string, arg0 *RID, arg1 string, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_String_int_int(o Class, methodName string, arg0 int64, arg1 int64) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_string
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(ret)

}
func GodotCall_void_int_Transform2D(o Class, methodName string, arg0 int64, arg1 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Variant_Transform2D_Object_Transform2D(o Class, methodName string, arg0 *Transform2D, arg1 *Object, arg2 *Transform2D) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_variant
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{ret}

}
func GodotCall_bool_Object_NodePath_Variant_Object_NodePath_float_int_int_float(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Variant, arg3 *Object, arg4 *NodePath, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := arg4.nodePath
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.double(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := C.int64_t(arg6)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := C.int64_t(arg7)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	cArg8 := C.double(arg8)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg8), C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_bool
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(*ret)

}
func GodotCall_RID_RID_String(o Class, methodName string, arg0 *RID, arg1 string) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_rid
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{ret}

}
func GodotCall_int_Object_bool(o Class, methodName string, arg0 *Object, arg1 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_int
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(*ret)

}
func GodotCall_void_PoolVector2Array_PoolIntArray(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolIntArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_int_Object_Vector2(o Class, methodName string, arg0 string, arg1 int64, arg2 *Object, arg3 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.int64_t(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_String_Object(o Class, methodName string, arg0 string, arg1 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := stringAsGodotString(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_PoolIntArray(o Class, methodName string, arg0 *PoolIntArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.array
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_Transform2D_RID(o Class, methodName string, arg0 *RID) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret *C.godot_transform2d
	retPtr := unsafe.Pointer(ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{ret}

}
func GodotCall_void_Object_String_Variant(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := stringAsGodotString(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.variant
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Vector2_Vector2(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_RID_Vector2(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := arg2.vector2
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_int_bool_int_Color_bool_int_Color_Object_Object(o Class, methodName string, arg0 int64, arg1 bool, arg2 int64, arg3 *Color, arg4 bool, arg5 int64, arg6 *Color, arg7 *Object, arg8 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := C.int64_t(arg0)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := C.godot_bool(arg1)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	cArg2 := C.int64_t(arg2)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg2), C.int(2))

	cArg3 := arg3.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg3), C.int(3))

	cArg4 := C.godot_bool(arg4)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg4), C.int(4))

	cArg5 := C.int64_t(arg5)
	C.add_element(cArgsArray, unsafe.Pointer(&cArg5), C.int(5))

	cArg6 := arg6.color
	C.add_element(cArgsArray, unsafe.Pointer(&cArg6), C.int(6))

	cArg7 := arg7.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg7), C.int(7))

	cArg8 := arg8.owner
	C.add_element(cArgsArray, unsafe.Pointer(&cArg8), C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func GodotCall_void_RID_Transform2D(o Class, methodName string, arg0 *RID, arg1 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := arg0.rid
	C.add_element(cArgsArray, unsafe.Pointer(&cArg0), C.int(0))

	cArg1 := arg1.transform2d
	C.add_element(cArgsArray, unsafe.Pointer(&cArg1), C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
