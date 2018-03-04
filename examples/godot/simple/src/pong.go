package main

import (
	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
)

// NewPongClass is a constructor that we can pass to godot.
func NewPongClass() godot.Class {
	pongClass := &PongClass{}

	return pongClass
}

// PongClass is a simple go struct that can be attached to a Godot Node2D object.
type PongClass struct {
	owner gdnative.Object
}

func (p *PongClass) BaseClass() string {
	return "Node"
}

func (p *PongClass) SetOwner(object gdnative.Object) {
	p.owner = object
}

func (p *PongClass) GetOwner() gdnative.Object {
	return p.owner
}

// Xready is called as soon as the node enters the scene.
func (p *PongClass) X_Ready() {
	// Log when the class is ready
	gdnative.Log.Warning("Ready called from pong.go!")
}

func (p *PongClass) X_Process(delta gdnative.Double) {
	gdnative.Log.Warning("Got delta:", delta)
}

func (p *PongClass) CustomThing(myPhrase string) string {
	gdnative.Log.Println(myPhrase)
	return "The phrase: '" + myPhrase + "' was returned ."
}

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// Enable debug logging
	godot.EnableDebug()

	// AutoRegister will register the given class constructor with Godot.
	godot.AutoRegister(NewPongClass)
}

// This never gets called, but it necessary to export as a shared library.
func main() {
}
