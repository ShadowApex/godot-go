#ifndef CGDNATIVE_NATIVESCRIPT_GATEWAY_H
#define CGDNATIVE_NATIVESCRIPT_GATEWAY_H

#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>

/* GDNative NATIVESCRIPT C Gateway */
void *cgo_gateway_create_func(godot_object *obj, void *method_data);
void *cgo_gateway_destroy_func(godot_object *obj, void *method_data,
			       void *user_data);
void *cgo_gateway_free_func(void *method_data);
godot_variant cgo_gateway_method_func(godot_object *obj, void *method_data,
				      void *user_data, int num_args,
				      godot_variant **args);
void cgo_gateway_property_set_func(godot_object *obj, void *method_data,
				   void *user_data, godot_variant *property);
godot_variant cgo_gateway_property_get_func(godot_object *obj,
					    void *method_data, void *user_data);

// Type definitions for any function pointers. Some of these are not defined in
// the godot headers when they are part of a typedef struct.
typedef void (*create_func)(godot_object *, void *);
typedef void (*destroy_func)(godot_object *, void *, void *);
typedef void (*free_func)(void *);
typedef godot_variant (*method)(godot_object *, void *, void *, int,
				godot_variant **);
typedef void (*set_property_func)(godot_object *, void *, void *,
				  godot_variant *);
typedef godot_variant (*get_property_func)(godot_object *, void *, void *);
godot_signal_argument **go_godot_signal_argument_build_array(int length);
void go_godot_signal_argument_add_element(godot_signal_argument **array,
					  godot_signal_argument *element,
					  int index);
#endif
