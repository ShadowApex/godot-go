#include "util.h"
#include <gdnative/gdnative.h>
#include <stdlib.h>

// Helper functions for accessing C arrays.
godot_gdnative_api_struct *cgo_get_ext(godot_gdnative_api_struct **ext, int i) {
	return ext[i];
}

void **go_void_build_array(int length) {
	void **arr = malloc(sizeof(void *) * length);

	return arr;
}

void go_void_add_element(void **array, void *element, int index) {
	array[index] = element;
}
