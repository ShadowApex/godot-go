package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

func NewNodePath(path string) *NodePath {
	nodePath := &NodePath{}
	gdString := stringAsGodotString(path)

	// Create our godot node path object
	var godotNodePath C.godot_node_path

	// Create our node path from the godotstring
	C.godot_node_path_new(&godotNodePath, gdString)

	return nodePath
}

type NodePath struct {
	nodePath *C.godot_node_path
}

func (n *NodePath) Copy() *NodePath {
	var godotNodePathCopy C.godot_node_path
	C.godot_node_path_new_copy(&godotNodePathCopy, n.nodePath)

	return &NodePath{
		nodePath: &godotNodePathCopy,
	}
}

func (n *NodePath) destroy() {
	C.godot_node_path_destroy(n.nodePath)
}

func (n *NodePath) String() string {
	gdString := C.godot_node_path_as_string(n.nodePath)
	defer C.godot_free(&gdString)

	return godotStringAsString(&gdString)
}

func (n *NodePath) IsAbsolute() bool {
	return bool(C.godot_node_path_is_absolute(n.nodePath))
}

func (n *NodePath) GetNameCount() int64 {
	return int64(C.godot_node_path_get_name_count(n.nodePath))
}

func (n *NodePath) GetName(index int64) string {
	gdString := C.godot_node_path_get_name(n.nodePath, C.godot_int(index))
	defer C.godot_free(&gdString)

	return godotStringAsString(&gdString)
}

func (n *NodePath) GetSubnameCount() int64 {
	return int64(C.godot_node_path_get_subname_count(n.nodePath))
}

func (n *NodePath) GetSubname(index int64) string {
	gdString := C.godot_node_path_get_subname(n.nodePath, C.godot_int(index))
	defer C.godot_free(&gdString)

	return godotStringAsString(&gdString)
}

func (n *NodePath) GetProperty() string {
	gdString := C.godot_node_path_get_property(n.nodePath)
	defer C.godot_free(&gdString)

	return godotStringAsString(&gdString)
}

func (n *NodePath) IsEmpty() bool {
	if n.nodePath == nil {
		return true
	}
	return bool(C.godot_node_path_is_empty(n.nodePath))
}

func (n *NodePath) Equals(node *NodePath) bool {
	return bool(C.godot_node_path_operator_equal(n.nodePath, node.nodePath))
}
