package godot

// Class is an interface for any objects that can have Godot
// inheritance.
type Class interface {
	Inherits() string
}

type Node struct{}

func (n *Node) Inherits() string {
	return "Node"
}

type Node2D struct{}

func (n *Node2D) Inherits() string {
	return "Node2D"
}
