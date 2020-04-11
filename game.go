package vngine

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

const dt = time.Second / 60.0

// NewGame is a simple constructor for the game struct.
func NewGame(settings, scenario string) (g game, err error) {
	g.GameData = new(GameData)
	g.settingsPath = settings
	am := newAssetManager()
	g.GameData.AssetManager = &am
	g.GameData.StateMachine = &stateMachine{}
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

	var newTime time.Time
	var frameTime, accumulator, interpolation time.Duration
	currentTime := time.Now()

	for !g.GameData.Window.Closed() {
		g.GameData.StateMachine.processStateChanges()

		if !g.GameData.StateMachine.hasStates() {
			fmt.Println("No more states. Closing the window")
			return
		}

		newTime = time.Now()
		frameTime = newTime.Sub(currentTime)

		if frameTime.Milliseconds() > 250 {
			frameTime = time.Millisecond * 250
		}

		currentTime = time.Now()
		accumulator += frameTime

		s := *g.GameData.StateMachine.getActiveState()

		for accumulator >= dt {
			s.HandleInput()
			s.Update(dt)

			accumulator -= dt
		}

		interpolation = accumulator / dt
		s.Draw(interpolation)
	}
}
