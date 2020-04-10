package vngine

import (
	"fmt"
	"github.com/faiface/pixel"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// newAssetManager is a simple constructor for the assetManager struct.
func newAssetManager() assetManager {
	return assetManager{make(map[string]*pixel.PictureData)}
}

// assetManager is an entity that loads the textures and frees them from the memory.
type assetManager struct {
	textures map[string]*pixel.PictureData
	// TODO: Sounds and music
}

// LoadTexture loads a given texture.
func (am *assetManager) LoadTexture(texName, path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("crititical failure - could not load the texture: %+v", err)
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		err = fmt.Errorf("crititical failure - could not decode the texture: %+v", err)
		return
	}
	am.textures[texName] = pixel.PictureDataFromImage(img)
	return
}

// RmTexture removes a given texture from memory.
func (am *assetManager) RmTexture(texName string) {
	delete(am.textures, texName)
}

// GetTexture returns the texture with the given name.
func (am *assetManager) GetTexture(texName string) *pixel.PictureData {
	return am.textures[texName]
}