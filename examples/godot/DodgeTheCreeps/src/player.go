package main

import (
	gd "github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
	"log"
)

// NewPlayer is a player constructor that we will register with Godot.
func NewPlayer() godot.Class {
	player := &Player{
		Speed: 400,
	}

	return player
}

// Player is a structure for the player.
type Player struct {
	godot.Area2D
	Speed          gd.Real
	screenSize     gd.Vector2
	animatedSprite godot.AnimatedSpriteImplementer
}

// X_Ready will be called as soon as the player enters the scene.
func (p *Player) X_Ready() {
	log.Println("X_Ready called!")

	// Set the speed.
	p.Speed = 400

	// Get the AnimatedSprite child node.
	log.Println("Getting animated sprite...")
	animatedSpritePath := gd.NewNodePath("AnimatedSprite")
	animatedSpriteNode := p.GetNode(animatedSpritePath)
	log.Println("Got animated sprite with ID:", animatedSpriteNode.GetBaseObject().ID())
	p.animatedSprite = animatedSpriteNode.(godot.AnimatedSpriteImplementer)

	// Get the viewport size
	viewportRect := p.GetViewportRect()
	p.screenSize = viewportRect.GetSize()
}

// X_Process will be called every frame.
func (p *Player) X_Process(delta gd.Double) {
	dt := gd.Real(delta)
	velocity := gd.NewVector2(0, 0)

	if godot.Input.IsActionPressed("ui_right") {
		velocity.SetX(velocity.GetX() + 1)
	}
	if godot.Input.IsActionPressed("ui_left") {
		velocity.SetX(velocity.GetX() - 1)
	}
	if godot.Input.IsActionPressed("ui_down") {
		velocity.SetY(velocity.GetY() + 1)
	}
	if godot.Input.IsActionPressed("ui_up") {
		velocity.SetY(velocity.GetY() - 1)
	}

	if velocity.Length() > 0 {
		normal := velocity.Normalized()
		velocity = normal.OperatorMultiplyScalar(p.Speed)
		p.animatedSprite.SetAnimation("right")
		p.animatedSprite.Play("right")
	} else {
		p.animatedSprite.Stop()
	}

	// Set the position based on velocity
	position := p.GetPosition()
	newPosition := position.OperatorAdd(velocity.OperatorMultiplyScalar(dt))
	p.SetPosition(newPosition)
}

func init() {
	// Set up logging to log to Godot.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(godot.Log)

	// AutoRegister our Player class.
	godot.AutoRegister(NewPlayer)
}

func main() {
}
