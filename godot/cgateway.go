package godot

/*
#include <stdio.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>

// This is a gateway function for the create method.
void *go_godot_instance_create_func_cgo(godot_object *obj, void *method_data)
{
	printf("CGO: C.go_godot_instance_create_func_cgo()\n");
	void go_godot_instance_create_func(godot_object *, void *);
	go_godot_instance_create_func(obj, method_data); // Execute our Go function.
	return 0;
}

// This is a gateway function for the free method.
void *go_godot_instance_free_func_cgo(void *method_data)
{
	printf("CGO: C.go_godot_instance_free_func_cgo()\n");
	void go_godot_instance_free_func(void *);
	go_godot_instance_free_func(method_data); // Execute our Go function.
	return 0;
}

*/
import "C"
