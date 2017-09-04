package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Object
import (
	"unsafe"
)

func NewRID(object Class) *RID {
	rid := &RID{}
	var gdRID C.godot_rid
	C.godot_rid_new_with_resource(&gdRID, unsafe.Pointer(object.getOwner()))
	rid.rid = &gdRID

	return rid
}

// NewEmptyRID returns a new resource ID.
func NewEmptyRID() *RID {
	rid := &RID{}
	var gdRID C.godot_rid
	C.godot_rid_new(&gdRID)
	rid.rid = &gdRID

	return rid
}

// RID is a resource ID.
type RID struct {
	rid *C.godot_rid
}

func (r *RID) ID() int64 {
	return int64(C.godot_rid_get_id(r.rid))
}

func (r *RID) Equals(rid *RID) bool {
	return bool(C.godot_rid_operator_equal(r.rid, rid.rid))
}

func (r *RID) Less(rid *RID) bool {
	return bool(C.godot_rid_operator_less(r.rid, rid.rid))
}

func (r *RID) Greater(rid *RID) bool {
	return !r.Less(rid)
}
