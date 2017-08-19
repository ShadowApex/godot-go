package godot

import "C"

import (
	"unsafe"
)

// unsafeToGoString will convert a null pointer with an underlying CString into
// a Go string.
func unsafeToGoString(p unsafe.Pointer) string {
	// Cast the pointer to a CString.
	cString := (*C.char)(p)
	goString := C.GoString(cString)

	return goString
}
