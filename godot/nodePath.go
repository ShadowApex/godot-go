package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

//TODO: godot_free is causing crashes currently. Need to fix this.
//import (
//	"unsafe"
//)

//DEPENDS: Built-In types
func NewNodePath(path string) *NodePath {
	nodePath := &NodePath{}

	// Create our godot node path object
	var godotNodePath C.godot_node_path

	// Create our node path from the godotstring
	C.godot_node_path_new(
		&godotNodePath,
		stringAsGodotString(path),
	)

	// Set the internal Godot NodePath
	nodePath.nodePath = &godotNodePath

	return nodePath
}

type NodePath struct {
	nodePath *C.godot_node_path
}

func (n *NodePath) nilCheck() {
	if n.nodePath == nil {
		panic("NodePath is nil! Are you not using NewNodePath() to create this?")
	}
}

func (n *NodePath) Copy() *NodePath {
	n.nilCheck()
	var godotNodePathCopy C.godot_node_path
	C.godot_node_path_new_copy(&godotNodePathCopy, n.nodePath)

	return &NodePath{
		nodePath: &godotNodePathCopy,
	}
}

func (n *NodePath) destroy() {
	n.nilCheck()
	C.godot_node_path_destroy(n.nodePath)
}

func (n *NodePath) AsString() string {
	n.nilCheck()
	gdString := C.godot_node_path_as_string(n.nodePath)
	//defer C.godot_free(unsafe.Pointer(&gdString))

	return godotStringAsString(&gdString)
}

func (n *NodePath) IsAbsolute() bool {
	n.nilCheck()
	return bool(C.godot_node_path_is_absolute(n.nodePath))
}

func (n *NodePath) GetNameCount() int64 {
	n.nilCheck()
	return int64(C.godot_node_path_get_name_count(n.nodePath))
}

func (n *NodePath) GetName(index int64) string {
	n.nilCheck()
	gdString := C.godot_node_path_get_name(n.nodePath, C.godot_int(index))
	//defer C.godot_free(unsafe.Pointer(&gdString))

	return godotStringAsString(&gdString)
}

func (n *NodePath) GetSubnameCount() int64 {
	n.nilCheck()
	return int64(C.godot_node_path_get_subname_count(n.nodePath))
}

func (n *NodePath) GetSubname(index int64) string {
	n.nilCheck()
	gdString := C.godot_node_path_get_subname(n.nodePath, C.godot_int(index))
	//defer C.godot_free(unsafe.Pointer(&gdString))

	return godotStringAsString(&gdString)
}

func (n *NodePath) IsEmpty() bool {
	if n.nodePath == nil {
		return true
	}
	return bool(C.godot_node_path_is_empty(n.nodePath))
}

func (n *NodePath) Equals(node *NodePath) bool {
	n.nilCheck()
	return bool(C.godot_node_path_operator_equal(n.nodePath, node.nodePath))
}
