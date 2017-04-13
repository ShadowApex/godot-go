package main

import (
	"fmt"
	"github.com/shadowapex/go-godot/godot"
)

var start = godot.SetNativeInit(NativeInit)

func NativeInit(options *godot.InitOptions) {
	fmt.Println("My custom native init!")
}

// main is needed to compile as a shared library.
func main() {}
