package main

import (
	"github.com/shadowapex/godot-go/godot"
	"log"
)

// NewPongClass is a constructor that we can pass to godot.
func NewPongClass() godot.Class {
	pongClass := &PongClass{}

	return pongClass
}

// PongClass is a simple Godot class that can be registered.
type PongClass struct {
	godot.Node2D
}

// Xready is called as soon as the node enters the scene.
func (p *PongClass) X_ready() {
	godot.Log.Println("Pong is ready!")
	name := p.GetName()
	godot.Log.Warning("Got name from parent method!: ", name)

	godot.Log.Warning(p.GetFilename())
	godot.Log.Warning("Google's IP Address: ", godot.IP.ResolveHostname("google.com", 1))

	//child := p.GetChild(0)
	//godot.Log.Warning("Got child!: ", child)
	//childName := child.GetName()
	//godot.Log.Warning("Got child name: ", childName)
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
