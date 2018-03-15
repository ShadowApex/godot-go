package main

import (
	gd "github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
	"log"
	"math/rand"
)

// NewMob is a mob constructor that we will register with Godot.
func NewMob() godot.Class {
	mob := &Mob{}

	return mob
}

// Mob is a structure for enemy mobs.
type Mob struct {
	godot.RigidBody2D
	MinSpeed       gd.Real
	MaxSpeed       gd.Real
	animatedSprite godot.AnimatedSpriteImplementer
}

// X_Ready will be called as soon as the mob enters the scene.
func (m *Mob) X_Ready() {
	log.Println("X_Ready called!")

	// Get the AnimatedSprite child node.
	animatedSpritePath := gd.NewNodePath("AnimatedSprite")
	animatedSpriteNode := m.GetNode(animatedSpritePath)
	m.animatedSprite = animatedSpriteNode.(godot.AnimatedSpriteImplementer)

	// Set up different mob types
	mobTypes := []gd.String{"walk", "swim", "fly"}

	// Randomly select a mob type
	m.animatedSprite.SetAnimation(mobTypes[rand.Int()%len(mobTypes)])
}

func (m *Mob) X_OnVisibilityScreenExited() {
	m.QueueFree()
}

// X_Process will be called every frame.
func (m *Mob) X_Process(delta gd.Double) {
}
