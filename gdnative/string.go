package gdnative

/*
#include <gdnative/string.h>
#include "gdnative.gen.h"
*/
import "C"

func NewStringWithWideString(str string) *String {
	return &String{base: stringAsGodotString(str)}
}
