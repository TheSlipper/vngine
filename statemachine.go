package vngine

// stateMachine is an entity responsible for management of the game states.
type stateMachine struct {
	stack       stateStack
	newState    State
	isRemoving  bool
	isAdding    bool
	isReplacing bool
}

// addState queues the putting of the given state on top of the stateMachine's stack.
func (s *stateMachine) addState(st State, replaces bool) {
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
		s.stack.pop()
		if !s.stack.isEmpty() {
			st := *s.stack.peek()
			st.resume()
		}
		s.isRemoving = false
	}
	if s.isAdding {
		if !s.stack.isEmpty() {
			if s.isReplacing {
				_ = s.stack.pop()
			} else {
				st := *s.stack.peek()
				st.pause()
			}
		}
		s.stack.push(&s.newState)
		st := *s.stack.peek()
		st.init()
		s.isAdding = false
	}
}

// getActiveState returns the pointer to the top state of the stack.
func (s *stateMachine) getActiveState() *State {
	return s.stack.peek()
}

func (s *stateMachine) hasStates() bool {
	return !(s.stack.length == 0)
}
