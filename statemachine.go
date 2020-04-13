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

import "runtime"

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains a basic state machine implementation used for management of the loaded
// game states (e.g. menu state).

// stateMachine is an entity responsible for management of the game states.
type stateMachine struct {
	stack       stateStack
	newState    *State
	isRemoving  bool
	isAdding    bool
	isReplacing bool
}

// addState queues the putting of the given state on top of the stateMachine's stack.
func (s *stateMachine) addState(st *State, replaces bool) {
	s.isAdding = true
	s.isReplacing = replaces

	s.newState = st
}

// rmTopState queues the removal of the state on top of the state machine's stack.
func (s *stateMachine) rmTopState() {
	s.isRemoving = true
}

// processStateChanges processes the queued events.
func (s *stateMachine) processStateChanges() {

	if s.isRemoving && !s.stack.isEmpty() {
		runtime.GC()
		s.stack.pop()
		if !s.stack.isEmpty() {
			st := *s.stack.peek()
			st.Resume()
		}
		s.isRemoving = false
	}
	if s.isAdding {
		runtime.GC()
		if !s.stack.isEmpty() {
			if s.isReplacing {
				_ = s.stack.pop()
			} else {
				st := *s.stack.peek()
				st.Pause()
			}
		}
		s.stack.push(s.newState)
		st := *s.stack.peek()
		st.Init()
		s.isAdding = false
	}
}

// getActiveState returns the pointer to the top state of the stack.
func (s *stateMachine) getActiveState() *State {
	return s.stack.peek()
}

// hasStates checks whether the state machine has any loaded states onto its internal stack.
func (s *stateMachine) hasStates() bool {
	return !(s.stack.length == 0)
}
