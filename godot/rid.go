package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

import (
	"unsafe"
)

// NewRID returns a new resource ID.
func NewRID() *RID {
	rid := &RID{}
	var gdRID C.godot_rid
	C.godot_rid_new(&gdRID)
	rid.rid = &gdRID

	return rid
}

// NewRIDWithResource returns the RID of an object.
func NewRIDWithResource(object Class) *RID {
	rid := &RID{}
	var gdRID C.godot_rid
	C.godot_rid_new_with_resource(&gdRID, unsafe.Pointer(object.getOwner()))
	rid.rid = &gdRID

	return rid
}

// RID is a resource ID.
type RID struct {
	rid *C.godot_rid
}

// GetID returns the godot resource ID
func (r *RID) GetID() int64 {
	return int64(C.godot_rid_get_id(r.rid))
}

// OperatorEqual returns true if r.rid == rid.rid
func (r *RID) OperatorEqual(rid *RID) bool {
	return bool(C.godot_rid_operator_equal(r.rid, rid.rid))
}

// OperatorGreater returns true if r.rid > rid.rid
func (r *RID) OperatorGreater(rid *RID) bool {
	if !r.OperatorEqual(rid) && !r.OperatorLess(rid) {
		return true
	}
	return false
}

// OperatorGreater returns true if r.rid < rid.rid
func (r *RID) OperatorLess(rid *RID) bool {
	return bool(C.godot_rid_operator_less(r.rid, rid.rid))
}
