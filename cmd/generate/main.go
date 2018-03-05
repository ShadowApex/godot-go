package main

import (
	"github.com/shadowapex/godot-go/cmd/generate/classes"
	"github.com/shadowapex/godot-go/cmd/generate/gdnative"
	"github.com/shadowapex/godot-go/cmd/generate/types"
	"os"
)

func main() {
	// Environment variables will tell us if we only want to do one.
	only := os.Getenv("ONLY")

	// Generate the gdnative bindings
	switch only {
	case "gdnative":
		gdnative.Generate()
	case "types":
		types.Generate()
	case "classes":
		classes.Generate()
	default:
		gdnative.Generate()
		types.Generate()
		classes.Generate()
	}
}
