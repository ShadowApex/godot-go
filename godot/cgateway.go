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

// This is a gateway function for the destroy method.
void *go_godot_instance_destroy_func_cgo(godot_object *obj, void *method_data, void *user_data)
{
	printf("CGO: C.go_godot_instance_destroy_func_cgo()\n");
	void go_godot_instance_destroy_func(godot_object *, void *, void *);
	go_godot_instance_destroy_func(obj, method_data, user_data); // Execute our Go function.
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

// This is a gateway function for the method
//GDCALLINGCONV godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);
godot_variant go_godot_instance_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args)
{
	printf("CGO: C.go_godot_instance_method_func_cgo()\n");
	void go_godot_instance_method_func(void *);
	go_godot_instance_method_func(method_data); // Execute our Go function.

	godot_variant ret;
    godot_variant_new_nil(&ret);
	return ret;
}
*/
import "C"
