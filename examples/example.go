package main

import (
	"fmt"
	"github.com/shadowapex/go-godot/godot"
)

func gdNativeInit(options *godot.GodotGDNativeInitOptions) {
}

func main() {
	godot.SetGodotGDNativeInit(gdNativeInit)
	fmt.Println("vim-go")
}
