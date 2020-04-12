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
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains entities responsible used for loading vngine startup data (e.g. resolution, vsync, etc.).

// getSettingsModelFromFile parses the settings file and returns the settings model with data from it.
func getSettingsModelFromFile(file string) (s SettingsModel, err error) {
	var xmlData []byte
	xmlData, err = ioutil.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("error while trying to read from the file '%s': %+v", file, err)
		return
	}
	err = xml.Unmarshal(xmlData, &s)
	if err != nil {
		err = fmt.Errorf("error while trying to interpret settings file '%s': %+v", file, err)
	}
	return
}

// SettingsModel represents the settings file and its elements.
type SettingsModel struct {
	XMLName     xml.Name `xml:"settings"`
	Name        string   `xml:"name"`
	VSync       bool     `xml:"vsync"`
	Resizable   bool     `xml:"resizable"`
	Width       float64  `xml:"width"`
	Height      float64  `xml:"height"`
	Undecorated bool     `xml:"undecorated"`
	AlwaysOnTop bool     `xml:"ontop"`
	Icon        string   `xml:"icon"`
	Fullscreen  bool     `xml:"fullscreen"`
}
