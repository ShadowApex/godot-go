package main

import (
	gd "github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
	"log"
	"math/rand"
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
	//gd.EnableDebug()

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
	Mob           godot.PackedSceneImplementer `hint:"ResourceType" usage:"Default"`
	mobPath       godot.Path2DImplementer
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

	// Get the mob path
	mobPathPath := gd.NewNodePath("MobPath")
	mobPathNode := m.GetNode(mobPathPath)
	m.mobPath = mobPathNode.(godot.Path2DImplementer)

	// Start the game
	m.NewGame()
}

func (m *Main) GameOver() {
	log.Println("Game over!")
	m.scoreTimer.Stop()
	m.mobTimer.Stop()
}

func (m *Main) NewGame() {
	log.Println("New Game")
	m.score = 0
	m.player.Start(m.startPosition.GetPosition())
	m.startTimer.Start()
}

func (m *Main) X_OnStartTimerTimeout() {
	log.Println("Start timer timeout")
	m.mobTimer.Start()
	m.scoreTimer.Start()
}

func (m *Main) X_OnScoreTimerTimeout() {
	log.Println("Score timer timeout")
	m.score += 1
}

func (m *Main) X_OnMobTimerTimeout() {
	log.Println("Mob timer timeout")
	mobSpawnLocationPath := gd.NewNodePath("MobSpawnLocation")
	mobSpawnLocation := (m.mobPath.GetNode(mobSpawnLocationPath)).(godot.PathFollow2DImplementer)
	mobSpawnLocation.SetOffset(gd.Real(rand.Int()))

	// Create a mob instance and add it to the scene
	if m.Mob.CanInstance() {
		mob := m.Mob.Instance(gd.Int(godot.PackedSceneGenEditStateDisabled))
		m.AddChild(mob, false)
	}
}

/*
func _on_MobTimer_timeout():
    # choose a random location on Path2D
    $MobPath/MobSpawnLocation.set_offset(randi())
    # create a Mob instance and add it to the scene
    var mob = Mob.instance()
    add_child(mob)
    # set the mob's direction perpendicular to the path direction
    var direction = $MobPath/MobSpawnLocation.rotation + PI/2
    # set the mob's position to a random location
    mob.position = $MobPath/MobSpawnLocation.position
    # add some randomness to the direction
    direction += rand_range(-PI/4, PI/4)
    mob.rotation = direction
    # choose the velocity
    mob.set_linear_velocity(Vector2(rand_range(mob.MIN_SPEED, mob.MAX_SPEED), 0).rotated(direction))

*/
