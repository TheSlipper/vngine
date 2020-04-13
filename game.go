//////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////// LICENCE
// VNgine - a simple robust visual novel engine.
// Copyright© 2020 Kornel Domeradzki
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
package vngine

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"log"
	"time"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains the constructor and definition of the game singleton and its state-shared
// subcomponent called gamedata.

const dt = 1.0 / 60.0
const VERSION = "ALPHA 0.0.1"

// NewGame is a simple constructor for the game struct.
func NewGame(settings, scenario string) (g game, err error) {
	g.GameData = new(GameData)
	g.settingsPath = settings
	am := newAssetManager()
	g.GameData.AssetManager = &am
	g.GameData.StateMachine = &stateMachine{}
	g.GameData.StartTime = time.Now()

	return
}

// DebugData contains the information on the debugging session.
type DebugData struct {
	dbgDataFormat string
	dbgTxt        *text.Text
	fps           int
	fpsAcc        int
	tick          <-chan time.Time
}

// GameData is a struct that contains all of the data used for managing the game flow.
type GameData struct {
	AssetManager *assetManager
	StateMachine *stateMachine
	Window       *pixelgl.Window
	StartTime    time.Time
}

// game is a struct that represents the game entity.
type game struct {
	settingsPath string
	GameData     *GameData
	cfg          pixelgl.WindowConfig
	sett         SettingsModel
	dbg          DebugData
}

// LoadSettings loads settings from the specified file.
func (g *game) LoadSettings() (err error) {
	// Load up the settings model from the file
	g.sett, err = getSettingsModelFromFile(g.settingsPath)
	if err != nil {
		log.Fatal(err)
	}
	// Apply it to the windowconfig struct
	// TODO: Fullscreen and icon
	g.cfg = pixelgl.WindowConfig{
		Title:       g.sett.Name,
		Bounds:      pixel.R(0, 0, g.sett.Width, g.sett.Height),
		VSync:       g.sett.VSync,
		Resizable:   g.sett.Resizable,
		Undecorated: g.sett.Undecorated,
		AlwaysOnTop: g.sett.AlwaysOnTop,
	}
	// Create the window and save the pointer to it in the game struct
	g.GameData.Window, err = pixelgl.NewWindow(g.cfg)
	if err != nil {
		return
	}

	// If the game is in debug mode then also load stuff for debugging
	if g.sett.Debugging {
		g.GameData.AssetManager.atlases["debug_atlas"] = text.NewAtlas(basicfont.Face7x13, text.ASCII)
		g.dbg.dbgTxt = text.New(pixel.V(0, 0), g.GameData.AssetManager.atlases["debug_atlas"])
		g.dbg.dbgTxt.Color = colornames.Yellow
		g.dbg.dbgDataFormat = "Build %s\r\nFramerate: %d FPS\r\nCurrently loaded state: %s"
		g.dbg.tick = time.Tick(time.Second)
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

		// Get the state and handle the input, game logic and drawing
		s = *g.GameData.StateMachine.getActiveState()
		s.HandleInput()
		s.Update(dt)
		s.Draw(dt)

		// If debugging then also draw the debug data
		if g.sett.Debugging {
			g.DebugDataProcessing()
		}
		g.GameData.Window.Update()
	}
}

func (g *game) DebugDataProcessing() {
	g.dbg.fpsAcc++
	select {
	case <-g.dbg.tick:
		g.dbg.fps = g.dbg.fpsAcc
		g.dbg.fpsAcc = 0
	default:
	}
	g.dbg.dbgTxt.Clear()
	_, _ = fmt.Fprintf(g.dbg.dbgTxt, g.dbg.dbgDataFormat, VERSION, g.dbg.fps, "todo: load the state name")
	g.dbg.dbgTxt.Draw(g.GameData.Window, pixel.IM.Moved(pixel.V(0, g.sett.Height - g.dbg.dbgTxt.LineHeight)))
}
