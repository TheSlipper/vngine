package vngine

import "testing"

func TestGame(t *testing.T) {
	g, err := NewGame("", "source examples/chapter/introduction_5")
	if err != nil {
		panic(err)
	}
	err = g.LoadSettings()
	if err != nil {
		panic(err)
	}
	st := newTestState(g.gameData)
	g.Run(&st)
}
