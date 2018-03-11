package main

import (
	"github.com/shadowapex/godot-go/gdnative"
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
	Speed          gdnative.Real
	velocity       gdnative.Vector2
	screenSize     gdnative.Rect2
	animatedSprite godot.AnimatedSpriteImplementer
}

// X_Ready will be called as soon as the player enters the scene.
func (p *Player) X_Ready() {
	log.Println("X_Ready called!")

	// Get the AnimatedSprite child node.
	log.Println("Getting animated sprite...")
	animatedSpritePath := gdnative.NewNodePath("AnimatedSprite")
	animatedSpriteNode := p.GetNode(animatedSpritePath)
	log.Println("Got animated sprite with ID:", animatedSpriteNode.GetBaseObject().ID())
	p.animatedSprite = animatedSpriteNode.(godot.AnimatedSpriteImplementer)

	// Get the viewport size
	//p.screenSize = godot.Viewport.GetVisibleRect()
}

// X_Process will be called every frame.
func (p *Player) X_Process(delta gdnative.Double) {
	p.velocity = gdnative.NewVector2(0, 0)

	if godot.Input.IsActionPressed("ui_right") {
		p.velocity.SetX(p.velocity.GetX() + 1)
	}
	if godot.Input.IsActionPressed("ui_left") {
		p.velocity.SetX(p.velocity.GetX() - 1)
	}
	if godot.Input.IsActionPressed("ui_down") {
		p.velocity.SetY(p.velocity.GetY() + 1)
	}
	if godot.Input.IsActionPressed("ui_up") {
		p.velocity.SetY(p.velocity.GetY() - 1)
	}

	if p.velocity.Length() > 0 {
		normal := p.velocity.Normalized()
		p.velocity = normal.OperatorMultiplyScalar(p.Speed)
		p.animatedSprite.SetAnimation("right")
		p.animatedSprite.Play("right")
	} else {
		p.animatedSprite.Stop()
	}

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
