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
	collisionShape godot.CollisionShape2DImplementer
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

	// Get the collision shape child node.
	collisionShapePath := gd.NewNodePath("CollisionShape2D")
	collisionShapeNode := p.GetNode(collisionShapePath)
	p.collisionShape = collisionShapeNode.(godot.CollisionShape2DImplementer)

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
		p.animatedSprite.Play("right")
	} else {
		p.animatedSprite.Stop()
	}

	// Set the position based on velocity
	position := p.GetPosition()
	newPosition := position.OperatorAdd(velocity.OperatorMultiplyScalar(dt))

	// Clamp our player's position to the size of the screen.
	newPosition.SetX(godot.Clamp(newPosition.GetX(), 0, p.screenSize.GetX()))
	newPosition.SetY(godot.Clamp(newPosition.GetY(), 0, p.screenSize.GetY()))
	p.SetPosition(newPosition)

	// Flip the sprite if we're moving left or down, since we only have sprites
	// for right and up.
	if velocity.GetX() != 0 {
		p.animatedSprite.SetAnimation("right")
		p.animatedSprite.SetFlipV(false)
		p.animatedSprite.SetFlipH(velocity.GetX() < 0)
	} else if velocity.GetY() != 0 {
		p.animatedSprite.SetAnimation("up")
		p.animatedSprite.SetFlipV(velocity.GetY() > 0)
	}
}

// X_OnPlayerBodyEntered will be called when the "body_entered" signal is
// fired. This is configured in the Player node settings from the Godot editor.
func (p *Player) X_OnPlayerBodyEntered() {
	log.Println("Player body was entered!")
	p.Hide()
	p.EmitSignal("hit")
	p.collisionShape.SetDisabled(true)
}

// Start will be called when we need to reset our player's position.
func (p *Player) Start(pos gd.Vector2) {
	p.SetPosition(pos)
	p.Show()
	p.collisionShape.SetDisabled(false)
}
