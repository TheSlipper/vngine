package vngine

import (
	"fmt"
	"testing"
	"time"
)

func newTestState(data *gameData) (ts testState) {
	ts.gd = data
	return
}

type testState struct {
	gd *gameData
	val      int
	initTime time.Time
}

func (ts *testState) init() {
	ts.initTime = time.Now()
	fmt.Println("init call")
}

func (ts *testState) handleInput() {
	fmt.Println("handle input call")
}

func (ts *testState) update(dt time.Duration) {
	ts.val++
	fmt.Println("update number", ts.val)
	if time.Now().Sub(ts.initTime).Seconds() >= 1 {
		fmt.Println("Turning it off")
		ts.gd.StateMachine.rmTopState()
	}
}

func (ts *testState) draw(dt time.Duration) {
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
	sm := stateMachine{isRemoving: false, isAdding: false, isReplacing: false}
	st = &testState{}
	sm.addState(st, true)
	sm.processStateChanges()
	sm.rmTopState()
	sm.addState(st, false)
	sm.addState(st, false)
	sm.rmTopState()
}
