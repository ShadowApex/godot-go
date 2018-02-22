package gdnative

/*
#include <stdio.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
#include <gdnative_api_struct.gen.h>

// This is a gateway function for the create method.
void *go_create_func_cgo(godot_object *obj, void *method_data)
{
	printf("CGO: C.go_create_func_cgo()\n");
	void *ret;
	void *go_create_func(godot_object *, void *);
	ret = go_create_func(obj, method_data); // Execute our Go function.
	return ret;
}

// This is a gateway function for the destroy method.
void *go_destroy_func_cgo(godot_object *obj, void *method_data, void *user_data)
{
	printf("CGO: C.go_destroy_func_cgo()\n");
	void *ret;
	void *go_destroy_func(godot_object *, void *, void *);
	ret = go_destroy_func(obj, method_data, user_data); // Execute our Go function.
	return ret;
}

// This is a gateway function for the free method.
void *go_free_func_cgo(void *method_data)
{
	printf("CGO: C.go_free_func_cgo()\n");
	void *ret;
	void *go_free_func(void *);
	ret = go_free_func(method_data); // Execute our Go function.
	return ret;
}

// This is a gateway function for the method
//GDCALLINGCONV godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);
//func go_method_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.uint, args **C.godot_variant) {
godot_variant go_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args)
{
	printf("CGO: C.go_method_func_cgo()\n");
	printf("CGO: Number of arguments: %d\n", num_args);
	godot_variant ret;
	godot_variant go_method_func(godot_object *, void *, void *, int, godot_variant **);
	ret = go_method_func(obj, method_data, user_data, num_args, args); // Execute our Go function.

	return ret;
}
*/
import "C"
