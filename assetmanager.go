package vngine

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
)

// newAssetManager is a simple constructor for the assetManager struct.
func newAssetManager() assetManager {
	return assetManager{
		make(map[string]*pixel.PictureData),
		make(map[string]*text.Atlas),
	}
}

// assetManager is an entity that loads the textures and frees them from the memory.
type assetManager struct {
	// TODO: Perhaps a mutex for loading the textures so that the game does not stop
	textures map[string]*pixel.PictureData
	atlases map[string]*text.Atlas
	// TODO: Sounds and music
}

// LoadTexture loads a given texture.
func (am *assetManager) LoadTexture(texName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("crititical failure - could not load the texture: %+v", err)
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		err = fmt.Errorf("crititical failure - could not decode the texture: %+v", err)
		return err
	}
	am.textures[texName] = pixel.PictureDataFromImage(img)
	return nil
}

// RmTexture removes a given texture from memory.
func (am *assetManager) RmTexture(texName string) {
	delete(am.textures, texName)
}

// GetTexture returns the pointer to the texture corresponding to the passed name.
func (am *assetManager) GetTexture(texName string) *pixel.PictureData {
	return am.textures[texName]
}

// LoadAtlas loads the specified font and creates an atlas for it.
func (am *assetManager) LoadAtlas(atlasName, path string, size float64) error {
	face, err := loadTTF(path, size)
	if err != nil {
		err = fmt.Errorf("crititical failure - could not load the font: %+v", err)
		return err
	}
	am.atlases[atlasName] = text.NewAtlas(face, text.ASCII)
	return nil
}

// RmAtlas deletes the atlas with the given name from the memory.
func (am *assetManager) RmAtlas(atlasName string) {
	delete(am.atlases, atlasName)
}

// GetAtlas returns the pointer to the atlas corresponding to the passed name.
func (am *assetManager) GetAtlas(atlasName string) *text.Atlas {
	return am.atlases[atlasName]
}

// loadTTF loads the font's face from a TTF file.
func loadTTF(path string, size float64) (font.Face, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	ft, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ft, &truetype.Options{
		Size: size,
		GlyphCacheEntries: 1,
	}), nil
}