#include <gdnative/gdnative.h>
#include <stdlib.h>

#ifndef CHELPER_H
#define CHELPER_H
godot_gdnative_api_struct *cgo_get_ext(godot_gdnative_api_struct **ext, int i);
void **go_void_build_array(int length);
void go_void_add_element(void **array, void *element, int index);
#endif
