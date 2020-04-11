package vngine

import (
	"github.com/faiface/pixel/pixelgl"
	"testing"
)

func TestGame(t *testing.T) {
	foo := func() {
		g, err := NewGame("", "source examples/chapter/introduction_5")
		if err != nil {
			panic(err)
		}
		err = g.LoadSettings()
		if err != nil {
			panic(err)
		}
		st := newTestState(g.GameData)
		g.Run(&st)
	}
	pixelgl.Run(foo)
}
