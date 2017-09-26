package main

import (
	"github.com/shadowapex/godot-go/godot"
	"log"
)

// NewPongClass is a constructor that we can pass to godot.
func NewPongClass() godot.Class {
	pongClass := &PongClass{
		direction: godot.NewVector2(1.0, 0.0),
	}

	return pongClass
}

// PongClass is a simple Godot class that can be registered.
type PongClass struct {
	godot.Node2D
	direction *godot.Vector2
}

// Xready is called as soon as the node enters the scene.
func (p *PongClass) X_ready() {
	// Get the screen size.
	screenSize := p.GetViewportRect()
	godot.Log.Warning("Screen size: ", screenSize.AsString())

	// Resolve Google's IP for funsies
	//godot.Log.Warning("Google's IP Address: ", godot.IP.ResolveHostname("google.com", 1))

	// Get the left paddle node.
	left := p.GetNode(godot.NewNodePath("left"))
	godot.Log.Warning("Left Pad: ", left)
	godot.Log.Warning("  Type: ", left.GetClass())
}

//func (p *PongClass) X_process(delta float64) {
//	godot.Log.Println("Processing in pong.go!")
//	godot.Log.Println("  Delta:", delta)
//}

func (p *PongClass) CustomThing(myPhrase string) string {
	godot.Log.Println(myPhrase)
	return "The phrase: '" + myPhrase + "' was returned ."
}

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// SetGodotGDNativeInit will set the given function to run on library initialization.
	godot.SetGodotGDNativeInit(func(options *godot.GodotGDNativeInitOptions) {
		log.Println("This is being called from pong.go!")
	})

	// Register will register the given class constructor with Godot.
	godot.Register(NewPongClass)
}

// This never gets called, but it necessary to export as a shared library.
func main() {
}
