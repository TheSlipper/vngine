//////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////// LICENCE
// VNgine - a simple robust visual novel engine.
// Copyright© 2020 Kornel Domeradzki
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
package vngine

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains all of the structs and functions associated with the asset manager.

// newAssetManager is a simple constructor for the assetManager struct.
func newAssetManager() assetManager {
	return assetManager{
		make(map[string]*pixel.PictureData),
		make(map[string]*text.Atlas),
	}
}

// assetManager is an entity that loads the textures and frees them from the memory.
type assetManager struct {
	textures map[string]*pixel.PictureData
	atlases  map[string]*text.Atlas
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
