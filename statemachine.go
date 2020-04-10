package vngine

type stateMachine struct {
	stack       stateStack
	newState    State
	isRemoving  bool
	isAdding    bool
	isReplacing bool
}

func (s *stateMachine) addState(st State, replaces bool) {
	s.isAdding = true
	s.isReplacing = replaces

	s.newState = st
}

func (s *stateMachine) rmTopState() {
	s.isRemoving = true
}

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

func (s *stateMachine) getActiveState() *State {
	return s.stack.peek()
}
