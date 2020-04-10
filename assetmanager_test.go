package vngine

import "testing"

// TestAssetManager tests whether asset manager is working correctly.
func TestAssetManager(t *testing.T) {
	am := newAssetManager()
	err := am.LoadTexture("testTex", "textures/hiking.png")
	if err != nil {
		panic(err)
	}
	_ = am.GetTexture("testTex")
	am.RmTexture("testTex")
	_ = am.GetTexture("testTex")
}
