package vngine

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

const dt = 1.0 / 60.0

// NewGame is a simple constructor for the game struct.
func NewGame(settings, scenario string) (g game, err error) {
	g.GameData = new(GameData)
	g.settingsPath = settings
	am := newAssetManager()
	g.GameData.AssetManager = &am
	g.GameData.StateMachine = &stateMachine{}
	g.GameData.StartTime = time.Now()
	i, err := newInterpreter(scenario)
	if err != nil {
		return
	}
	g.GameData.Interpreter = &i
	return
}

// GameData is a struct that contains all of the data used for managing the game flow.
type GameData struct {
	AssetManager *assetManager
	StateMachine *stateMachine
	Interpreter *interpreter
	Window *pixelgl.Window
	StartTime time.Time
}

// game is a struct that represents the game entity.
type game struct {
	settingsPath string
	GameData     *GameData
	cfg          pixelgl.WindowConfig
}

// LoadSettings loads settings from the specified file.
func (g *game) LoadSettings() (err error) {
	// TODO: Make a model with all the necessary settings. Load the settings from an xml file into that model.
	g.cfg = pixelgl.WindowConfig {
		Title: "vngine",
		Bounds: pixel.R(0, 0, 1280, 720),
		VSync: true,
	}
	g.GameData.Window, err = pixelgl.NewWindow(g.cfg)
	if err != nil {
		return
	}
	return
}

// Run starts the game.
func (g *game) Run(st State) {
	// Load the passed state
	g.GameData.StateMachine.addState(st, true)

	// Placeholder for the current state
	var s State

	for !g.GameData.Window.Closed() {

		g.GameData.StateMachine.processStateChanges()

		if !g.GameData.StateMachine.hasStates() {
			fmt.Println("No more states. Closing the window")
			return
		}

		s = *g.GameData.StateMachine.getActiveState()
		s.HandleInput()
		s.Update(dt)
		s.Draw(dt)
		g.GameData.Window.Update()
	}
}
