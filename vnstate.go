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
	"github.com/faiface/pixel/pixelgl"
	"strconv"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains an implementation of the state interface that is responsible for the main state of
// the visual novel (the game - reading, animations, etc.).

func NewVNState(scpath string, gd *GameData) (vns VNState, err error) {
	var interp interpreter
	interp, err = newInterpreter(scpath)
	if err != nil {
		return
	}
	vns.interp = &interp
	vns.gd = gd
	vns.currEntry, err = vns.interp.nextEntry()
	if err != nil {
		return
	}
	DebugLog("Entry ID: " + strconv.Itoa(vns.currEntry.ID) + "; Text content: " + vns.currEntry.Texts[0].InnerTxt)
	return
}

// VNState is a state that uses the vngine interpreter and reads the story from it.
type VNState struct {
	gd                *GameData
	firstScenarioPath string
	interp            *interpreter
	currEntry         Entry
	name              string
}

// Init initializes the visual novel state.
func (vns *VNState) Init() {
	var err error
	vns.name = "vngine interpreter state"
	vns.currEntry, err = vns.interp.nextEntry()
	if err != nil {
		DebugLog(err.Error())
		vns.gd.StateMachine.rmTopState()
	}
}

// HandleInput executes given actions on given user input events.
func (vns *VNState) HandleInput() {
	if vns.gd.Window.JustPressed(pixelgl.MouseButtonLeft) {
		var err error
		vns.currEntry, err = vns.interp.nextEntry()
		if err != nil {
			DebugLog(err.Error())
		}
		DebugLog("Entry ID: " + strconv.Itoa(vns.currEntry.ID) + "; Text content: " + vns.currEntry.Texts[0].InnerTxt)
		// TODO: Some kind of check for clicking the UI
	} else if vns.gd.Window.JustPressed(pixelgl.MouseButtonRight) {
		DebugLog("Right mouse button clicked")
	} else if vns.gd.Window.JustPressed(pixelgl.KeyLeftControl) || vns.gd.Window.JustPressed(pixelgl.KeyRightControl) {
		DebugLog("Control button clicked")
	} else if vec := vns.gd.Window.MouseScroll(); vec.X != 0 || vec.Y != 0 {
		DebugLog("Mouse scrolled: " + vec.String())
	}
}

// Update updates some data of the game when called.
func (vns *VNState) Update(dt float64) {
	// TODO: Maybe delete since there is no physics here?
}

// Draw is a method called when the vngine is ready to draw the state.
func (vns *VNState) Draw(dt float64) {

}

// Pause pauses the execution of the state.
func (vns *VNState) Pause() {

}

// Resume resumes the execution of the state.
func (vns *VNState) Resume() {

}

// Name returns the name of the state.
func (vns *VNState) Name() string {
	return vns.name
}

func (vns *VNState) loadAssets() {

}
