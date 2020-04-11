package vngine

import (
	"fmt"
	"testing"
	"time"
)

func newTestState(data *GameData) (ts testState) {
	ts.gd = data
	return
}

type testState struct {
	gd *GameData
	val      int
	initTime time.Time
}

func (ts *testState) Init() {
	ts.initTime = time.Now()
	fmt.Println("Init call")
}

func (ts *testState) HandleInput() {
	fmt.Println("handle input call")
}

func (ts *testState) Update(dt time.Duration) {
	ts.val++
	fmt.Println("Update number", ts.val)
	if time.Now().Sub(ts.initTime).Seconds() >= 1 {
		fmt.Println("Turning it off")
		ts.gd.StateMachine.rmTopState()
	}
}

func (ts *testState) Draw(dt time.Duration) {
	//ts.gd.window.Clear(colornames.Skyblue)
	//ts.gd.window.Update()
}

func (ts *testState) Pause() {
	fmt.Println("Pause call")
}

func (ts *testState) Resume() {
	fmt.Println("Resume call")
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
