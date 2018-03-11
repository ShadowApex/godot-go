package main

import (
	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
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
	animatedSprite *godot.AnimatedSprite
}

// X_Ready will be called as soon as the player enters the scene.
func (p *Player) X_Ready() {
	godot.Log.Println("X_Ready called!")

	// Get the node path.
	nodePath := p.GetPath()
	godot.Log.Println("Node path: ", nodePath.AsString())
	godot.Log.Println("  Empty:", nodePath.IsEmpty())

	// Get the class type
	godot.Log.Println("Class: ", p.GetClass())
	godot.Log.Println("Children:", p.GetChildCount())
	//child := p.GetChild(1)
	//godot.Log.Println("  Child name:", child.GetName())

	// Get the animated sprite
	animatedSpritePath := gdnative.NewNodePath("AnimatedSprite")
	animatedSpriteNode := p.GetNode(animatedSpritePath)
	p.animatedSprite = animatedSpriteNode.(*godot.AnimatedSprite)

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

	// Get the AnimatedSprite child node.
	//nodePath := gdnative.NewNodePath("AnimatedSprite")
	//animatedSpriteNode := p.GetNode(nodePath)
	//godot.Log.Println("Fetched Node ID: ", animatedSpriteNode.GetBaseObject().ID())

	// "Cast" the node to an AnimatedSprite
	//animatedSprite := godot.AnimatedSprite{}
	//animatedSprite.SetBaseObject(animatedSpriteNode.GetBaseObject())
	if p.velocity.Length() > 0 {
		normal := p.velocity.Normalized()
		p.velocity = normal.OperatorMultiplyScalar(p.Speed)
		//godot.Log.Println("Animated Sprite ID: ", animatedSprite.GetBaseObject().ID())
		//animatedSprite.SetAnimation("right")
		//animatedSprite.Play("right")
	} else {
		//animatedSprite.Stop()
	}

}

func init() {
	// AutoRegister our Player class.
	godot.AutoRegister(NewPlayer)
}

func main() {
}
