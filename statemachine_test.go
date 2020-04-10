package vngine

import (
	"fmt"
	"testing"
)

type testState struct {
	val int
}

func (ts *testState) init() {
	fmt.Println("init call")
}

func (ts *testState) handleInput() {
	fmt.Println("handle input call")
}

func (ts *testState) update(dt float32) {
	ts.val++
	fmt.Println("update call")
}

func (ts *testState) draw(dt float32) {
	fmt.Println("draw call")
}

func (ts *testState) pause() {
	fmt.Println("pause call")
}

func (ts *testState) resume() {
	fmt.Println("resume call")
}

// TestStateMachine tests whether state machine is working correctly.
func TestStateMachine(t *testing.T) {
	var st State
	sm := stateMachine{isRemoving:false,isAdding:false,isReplacing:false}
	st = &testState{}
	sm.addState(st,true)
	sm.processStateChanges()
	sm.rmTopState()
	sm.addState(st,false)
	sm.addState(st,false)
	sm.rmTopState()
}
