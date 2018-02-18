#include "util.h"
#include <gdnative/gdnative.h>

// Helper functions for accessing C arrays.
godot_gdnative_api_struct *cgo_get_ext(godot_gdnative_api_struct **ext, int i) {
	return ext[i];
}
