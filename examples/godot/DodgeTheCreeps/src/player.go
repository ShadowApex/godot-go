package main

import (
	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
	"github.com/shadowapex/godot-go/godot/class"
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
	class.Area2D
	Speed      gdnative.Real
	velocity   gdnative.Vector2
	screenSize gdnative.Rect2
}

// X_Ready will be called as soon as the player enters the scene.
func (p *Player) X_Ready() {
	godot.Log.Println("X_Ready called!")

	// Get the viewport size
	//p.screenSize = class.Viewport.GetVisibleRect()
}

// X_Process will be called every frame.
func (p *Player) X_Process(delta gdnative.Double) {
	p.velocity = gdnative.NewVector2(0, 0)

	if class.Input.IsActionPressed("ui_right") {
		p.velocity.SetX(p.velocity.GetX() + 1)
	}
	if class.Input.IsActionPressed("ui_left") {
		p.velocity.SetX(p.velocity.GetX() - 1)
	}
	if class.Input.IsActionPressed("ui_down") {
		p.velocity.SetY(p.velocity.GetY() + 1)
	}
	if class.Input.IsActionPressed("ui_up") {
		p.velocity.SetY(p.velocity.GetY() - 1)
	}

	// Get the AnimatedSprite child node.
	spritePath := gdnative.NewNodePath("AnimatedSprite")
	animatedSpriteNode := p.GetNode(spritePath)

	// "Cast" the node to an AnimatedSprite
	animatedSprite := class.AnimatedSprite{}
	animatedSprite.SetBaseObject(animatedSpriteNode.GetBaseObject())
	if p.velocity.Length() > 0 {
		normal := p.velocity.Normalized()
		p.velocity = normal.OperatorMultiplyScalar(p.Speed)
		animatedSprite.Play("")
	} else {
		animatedSprite.Stop()
	}

}

func init() {
	// AutoRegister our Player class.
	godot.AutoRegister(NewPlayer)
}

func main() {
}
