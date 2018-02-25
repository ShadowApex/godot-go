package gdnative

/*
#include <gdnative/string.h>
#include <gdnative/variant.h>
#include "gdnative.gen.h"
*/
import "C"

func NewVariantWithString(str *String) Variant {
	var variant C.godot_variant
	C.go_godot_variant_new_string(GDNative.api, &variant, str.base)

	return Variant{base: &variant}
}
