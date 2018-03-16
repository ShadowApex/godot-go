package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

// Signal is a structure definition that defines a custom Godot signal.
type Signal struct {
	Name        string
	Args        []SignalArg
	DefaultArgs []SignalDefaultArg
}

// SignalArg is any valid Godot class or variant type.
type SignalArg struct {
	Name       gdnative.String
	Value      interface{}
	Hint       gdnative.PropertyHint
	HintString gdnative.String
	Usage      gdnative.PropertyUsageFlags
	//Type         Int
	//DefaultValue Variant
}

type SignalDefaultArg interface{}
