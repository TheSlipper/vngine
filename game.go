package vngine

import (
	"time"
)

const dt = time.Second / 60.0

// NewGame is a simple constructor for the game struct.
func NewGame(settings, scenario string) (g game, err error) {
	g.gameData = new(gameData)
	g.settingsPath = settings
	am := newAssetManager()
	g.gameData.AssetManager = &am
	g.gameData.StateMachine = &stateMachine{}
	i, err := newInterpreter(scenario)
	if err != nil {
		return
	}
	g.gameData.Interpreter = &i
	return
}

// gameData is a struct that contains all of the data used for managing the game flow.
type gameData struct {
	AssetManager *assetManager
	StateMachine *stateMachine
	Interpreter *interpreter
}

// game is a struct that represents the game entity.
type game struct {
	settingsPath string
	gameData *gameData
}

// LoadSettings loads settings from the specified file.
func (g *game) LoadSettings() (err error) {
	// TODO: Make a model with all the necessary settings. Load the settings from an xml file into that model.
	return
}

// Run starts the game.
func (g *game) Run(st State) {
	// Load the passed state
	g.gameData.StateMachine.addState(st, true)

	// Do magic TODO: Document this
	var newTime time.Time
	var frameTime, accumulator, interpolation time.Duration
	currentTime := time.Now()

	// TODO: While window is open
	// TODO: Test this out
	for {
		if !g.gameData.StateMachine.hasStates() {
			return
		}

		g.gameData.StateMachine.processStateChanges()

		newTime = time.Now()
		frameTime = newTime.Sub(currentTime)

		if frameTime.Milliseconds() > 250 {
			frameTime = time.Millisecond * 250
		}

		currentTime = time.Now()
		accumulator += frameTime

		s := *g.gameData.StateMachine.getActiveState()

		for accumulator >= dt {
			s.handleInput()
			s.update(dt)

			accumulator -= dt
		}

		interpolation = accumulator / dt
		s.draw(interpolation)
	}
}
