#include "variant.h"
#include <gdnative/variant.h>
#include <stdlib.h>

godot_variant **go_godot_variant_build_array(int length) {
	godot_variant **arr = malloc(sizeof(godot_variant *) * length);

	return arr;
}

void go_godot_variant_add_element(godot_variant **array, godot_variant *element,
				  int index) {
	array[index] = element;
}
