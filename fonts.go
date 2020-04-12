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
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"os"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains helper functions used for font data manipulation.

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
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}
