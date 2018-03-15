package main

import (
	gd "github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
	"log"
)

// main is required to be exported as a shared library
func main() {
}

// Init will be called when this library is loaded by Godot.
func init() {
	// Set up logging to log to Godot.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(godot.Log)

	// Enable debug output
	//godot.EnableDebug()

	// AutoRegister our Player and Mob classes.
	godot.AutoRegister(NewMain)
	godot.AutoRegister(NewPlayer)
	godot.AutoRegister(NewMob)
}

// NewMain is a Main constructor that we will register with Godot.
func NewMain() godot.Class {
	m := &Main{}

	return m
}

// Main is a structure for the main game script.
type Main struct {
	godot.Node
	Mob           godot.PackedSceneImplementer `hint:"ResourceType"`
	mobTimer      godot.TimerImplementer
	player        *Player
	score         int
	scoreTimer    godot.TimerImplementer
	startTimer    godot.TimerImplementer
	startPosition godot.Node2DImplementer
}

// X_Ready will be called as soon as the main node enters the scene.
func (m *Main) X_Ready() {
	log.Println("X_Ready called!")
	log.Println("Registry:", godot.InstanceRegistry)

	// Get our score and mob timers
	scoreTimerPath := gd.NewNodePath("ScoreTimer")
	scoreTimerNode := m.GetNode(scoreTimerPath)
	m.scoreTimer = scoreTimerNode.(godot.TimerImplementer)

	startTimerPath := gd.NewNodePath("StartTimer")
	startTimerNode := m.GetNode(startTimerPath)
	m.startTimer = startTimerNode.(godot.TimerImplementer)

	mobTimerPath := gd.NewNodePath("MobTimer")
	mobTimerNode := m.GetNode(mobTimerPath)
	m.mobTimer = mobTimerNode.(godot.TimerImplementer)

	// Get the player
	playerPath := gd.NewNodePath("Player")
	playerNode := m.GetNode(playerPath)
	m.player = playerNode.(*Player)

	// Get the starting position
	startPosPath := gd.NewNodePath("StartPosition")
	startPosNode := m.GetNode(startPosPath)
	m.startPosition = startPosNode.(godot.Node2DImplementer)
}

func (m *Main) GameOver() {
	m.scoreTimer.Stop()
	m.mobTimer.Stop()
}

func (m *Main) NewGame() {
	m.score = 0
	m.player.Start(m.startPosition.GetPosition())
	m.startTimer.Start()
}
