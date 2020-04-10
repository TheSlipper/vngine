package vngine

type State interface {
	init()
	handleInput()
	update(dt float32)
	draw(dt float32)
	pause()
	resume()
}

// stateStack is a stack that contains State interfaces.
type stateStack struct {
	top    *stateNode
	length int
}

// stateNode is a single node used in the stateStack.
type stateNode struct {
	val *State
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

