package main

import (
	"github.com/shadowapex/godot-go/cmd/generate/gdnative"
	"github.com/shadowapex/godot-go/cmd/generate/types"
)

func main() {
	// Generate the gdnative bindings
	gdnative.Generate()
	types.Generate()
}
