package main

import (
	"log"

	"github.com/shadowapex/godot-go/godot"
	"github.com/shadowapex/godot-go/godot/class"
)

// NewPongClass is a constructor that we can pass to godot.
func NewPongClass() godot.Class {
	pongClass := &PongClass{}

	return pongClass
}

// PongClass is a simple go struct that can be attached to a Godot Node2D object.
type PongClass struct {
	class.Node2D
}

// Xready is called as soon as the node enters the scene.
func (p *PongClass) X_Ready() {
	godot.Log.Warning("X_Ready was called in pong.go!")
}

func (p *PongClass) CustomThing(myPhrase string) string {
	godot.Log.Println(myPhrase)
	return "The phrase: '" + myPhrase + "' was returned ."
}

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// Register will register the given class constructor with Godot.
	godot.AutoRegister(NewPongClass)
}

// This never gets called, but it necessary to export as a shared library.
func main() {
}
