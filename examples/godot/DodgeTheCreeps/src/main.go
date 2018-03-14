package main

import (
	"github.com/shadowapex/godot-go/godot"
	"log"
)

func init() {
	// Set up logging to log to Godot.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(godot.Log)

	// Enable debug output
	//godot.EnableDebug()

	// AutoRegister our Player and Mob classes.
	godot.AutoRegister(NewPlayer)
	godot.AutoRegister(NewMob)
}

func main() {
}
