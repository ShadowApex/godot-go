package main

import (
	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
	"github.com/shadowapex/godot-go/godot/class"
	"runtime/debug"
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
	godot.Log.Println("X_Process called with delta:", delta)
	godot.Log.Println("  Creating vector...")
	p.velocity = gdnative.NewVector2(0, 0)

	// Works
	godot.Log.Println("Setting mouse mode")
	class.Input.SetMouseMode(1)

	// Works
	godot.Log.Println("Getting mouse button mask...")
	mask := class.Input.GetMouseButtonMask()
	godot.Log.Println("Got button mask: ", mask)

	// Doesn't work
	godot.Log.Println("Checking if key pressed")
	isPressed := class.Input.IsKeyPressed(1)
	godot.Log.Println("Got key pressed: ", isPressed)

	// Doesn't work
	godot.Log.Println("Checking if joy is known...")
	joyKnown := class.Input.IsJoyKnown(1)
	godot.Log.Println("Got joy known: ", joyKnown)

	// Doesn't work
	godot.Log.Println("Getting audio bus count...")
	busCount := class.AudioServer.GetBusCount()
	godot.Log.Println("Got num buses: ", busCount)

	godot.Log.Println("  Checking if ui_right pressed...")
	if class.Input.IsActionPressed("ui_right") {
		godot.Log.Println("  p.velocity.x += 1...")
		p.velocity.SetX(p.velocity.GetX() + 1)
	}
	godot.Log.Println("  Checking if ui_left pressed...")
	if class.Input.IsActionPressed("ui_left") {
		godot.Log.Println("  p.velocity.x -= 1...")
		p.velocity.SetX(p.velocity.GetX() - 1)
	}
	godot.Log.Println("  Checking if ui_down pressed...")
	if class.Input.IsActionPressed("ui_down") {
		godot.Log.Println("  p.velocity.y += 1...")
		p.velocity.SetY(p.velocity.GetY() + 1)
	}
	godot.Log.Println("  Checking if ui_up pressed...")
	if class.Input.IsActionPressed("ui_up") {
		godot.Log.Println("  p.velocity.y -= 1...")
		p.velocity.SetY(p.velocity.GetY() - 1)
	}

	// Get the AnimatedSprite child node.
	godot.Log.Println("  Getting AnimatedSprite...")
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
	debug.SetGCPercent(-1)
	// AutoRegister our Player class.
	godot.AutoRegister(NewPlayer)
}

func main() {
}
