package main

import (
	"log"

	"github.com/shadowapex/godot-go/godot"
)

// NewPongClass is a constructor that we can pass to godot.
func NewPongClass() godot.Class {
	pongClass := &PongClass{
		direction: godot.NewVector2(1.0, 0.0),
	}

	return pongClass
}

// PongClass is a simple go struct that can be attached to a Godot Node2D object.
type PongClass struct {
	godot.Node2D
	direction *godot.Vector2
}

// Xready is called as soon as the node enters the scene.
func (p *PongClass) X_ready() {
	// Get the screen size.
	screenSize := p.GetViewportRect()
	godot.Log.Warning("***Screen size***")
	godot.Log.Warning(screenSize.AsString())

	// Resolve Google's IP for funsies
	//godot.Log.Warning("Google's IP Address: ", godot.IP.ResolveHostname("google.com", 1))

	// Get the left paddle node.
	godot.Log.Warning("***Get Left***")
	left := p.GetNode(godot.NewNodePath("left"))
	godot.Log.Warning(left)

	// Troubleshooting
	//godot.Log.Warning("***pongClass***")
	//godot.Log.Warning(p)
	//godot.Log.Warning("***left***")
	//godot.Log.Warning(left)

	//godot.Log.Warning("***pongClass Owner***")
	//godot.Log.Warning(p.GetOwner())
	//godot.Log.Warning("***Left Owner***")
	//godot.Log.Warning(left.GetOwner())

	//godot.Log.Warning("***pongClass Name***")
	//godot.Log.Warning(p.GetName())
	//godot.Log.Warning("***Left Name***")
	//godot.Log.Warning(left.GetName())

	//godot.Log.Warning("***pongClass Class***")
	//godot.Log.Warning(p.GetClass())
	//godot.Log.Warning("***Left Class***")
	//godot.Log.Warning(left.GetClass())
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
