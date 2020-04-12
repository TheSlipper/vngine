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

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains an implementation of the state interface that is responsible for the main state of
// the visual novel (the game - reading, animations, etc.).

func NewVNState(scenPath string) (vns VNState, err error) {
	var interp interpreter
	interp, err = newInterpreter(scenPath)
	if err != nil {
		return
	}
	vns.interp = &interp
	return
}

// VNState is a state that uses the vngine interpreter and reads the story from it.
type VNState struct {
	gd                *GameData
	firstScenarioPath string
	interp            *interpreter
	currEntry         EntryModel
}

func (vns *VNState) Init() {
	vns.currEntry = vns.interp.nextEntry()
}

func (vns *VNState) HandleInput() {
	// TODO:
}

func (vns *VNState) Update(dt float64) {
	// TODO: Maybe delete since there is no physics here?
}

func (vns *VNState) Draw(dt float64) {

}

func (vns *VNState) Pause() {

}

func (vns *VNState) Resume() {

}

func (vns *VNState) handleMusic() {

}

func (vns *VNState) handleSounds() {

}
