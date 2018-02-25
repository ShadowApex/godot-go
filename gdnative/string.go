package gdnative

/*
#include <gdnative/string.h>
#include "gdnative.gen.h"
*/
import "C"

func NewString() *String {
	var str C.godot_string
	C.go_godot_string_new(GDNative.api, &str)

	return &String{base: &str}
}

func NewStringWithWideString(str string) *String {
	return &String{base: stringAsGodotString(str)}
}
