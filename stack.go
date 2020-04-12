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
// This file contains a very simple implementation of a stack which is used for storing game states.

// stateStack is a stack that contains State interfaces.
type stateStack struct {
	top    *stateNode
	length int
}

// stateNode is a single node used in the stateStack.
type stateNode struct {
	val  *State
	prev *stateNode
}

// NewStateStack is a simple constructor for a State stack.
func NewStateStack() *stateStack {
	return &stateStack{nil, 0}
}

// isEmpty checks whether the stack is empty.
func (s *stateStack) isEmpty() bool {
	return s.length == 0
}

// peek peeks at the value on top of the stack.
func (s *stateStack) peek() *State {
	if s.length == 0 {
		return nil
	}
	return s.top.val
}

// Pop the top item of the stack.
func (s *stateStack) pop() *State {
	if s.length == 0 {
		return nil
	}
	foo := s.top
	s.top = foo.prev
	s.length--
	return foo.val
}

// Push a value on top of the stack.
func (s *stateStack) push(st *State) {
	foo := &stateNode{st, s.top}
	s.top = foo
	s.length++
}
