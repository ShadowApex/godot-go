package godot

/*
#include <stdio.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>

// The gateway function
void *go_godot_instance_create_func_cgo(godot_object *obj, void *method_data)
{
	printf("C.go_godot_instance_create_func_cgo()\n");
	void go_godot_instance_create_func(godot_object *, void *);
	go_godot_instance_create_func(obj, method_data); // Execute our Go function.
	return 0;
}
*/
import "C"
