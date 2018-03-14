package gdnative

/*
#include <gdnative/string.h>
#include "gdnative.gen.h"
*/
import "C"

func NewStringWithWideString(str string) String {
	return String(str)
}

func NewString() String {
	return ""
}

func NewStringCopy(src String) String {
	return ""
}

func truncateString(str string, num int) string {
	return str[0:num]
}
