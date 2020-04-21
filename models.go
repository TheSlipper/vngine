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
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains all of the structs used for extraction and representation of the engine's
// scripting language.

// Chapter represents the chapter element in the vnscript file.
type Chapter struct {
	XMLName   xml.Name   `xml:"chapter"`
	Title     string     `xml:"title,attr"`
	Scenarios []Scenario `xml:"scenario"`
}

// Scenario represents the scenario elements in the vnscript file.
type Scenario struct {
	XMLName xml.Name  `xml:"scenario"`
	ID      int       `xml:"id,attr"`
	Assets  AssetPath `xml:"asset-path"`
	Entries []Entry   `xml:"entry"`
}

// AssetPath represents the asset-path elements in the vnscript file.
type AssetPath struct {
	XMLName xml.Name `xml:"asset-path"`
	Source  string   `xml:"src,attr"`
}

// Entry represents the entry elements in the vnscript file.
type Entry struct {
	XMLName          xml.Name     `xml:"entry"`
	ID               int          `xml:"id,attr"`
	Transmission     string       `xml:"transmission,attr"`
	Interactable     bool         `xml:"interactable,attr"`
	DeferInteraction bool         `xml:"defer-interaction,attr"`
	ForwardTo        string       `xml:"forward-to,attr"`
	BG               Background   `xml:"background"`
	MusicEvents      []MusicEvent `xml:"music-event"`
	Chars            []Character  `xml:"character"`
	Texts            []Text       `xml:"text"`
	Choices          Choices      `xml:"choices"`
}

// Background represents the background elements in the vnscript file.
type Background struct {
	XMLName xml.Name `xml:"background"`
	Name    string   `xml:"name,attr"`
}

// MusicEvent represents the music-event elements in the vnscript file.
type MusicEvent struct {
	XMLName     xml.Name `xml:"music-event"`
	Action      string   `xml:"action,attr"`
	Track       string   `xml:"track,attr"`
	Repeatable  bool     `xml:"repeatable,attr"`
	EffectNames string   `xml:"effect-names,attr"`
}

// Character represents the character elements in the vnscript file.
type Character struct {
	XMLName xml.Name `xml:"character"`
	ID      int      `xml:"id,attr"`
	State   string   `xml:"state,attr"`
	XPos    float32  `xml:"x-pos,attr"`
	YPos    float32  `xml:"y-pos,attr"`
}

// Text represents the text elements in the vnscript file.
type Text struct {
	XMLName         xml.Name `xml:"text"`
	Type            string   `xml:"type,attr"`
	CharID          int      `xml:"character-id,attr"`
	CharNamePattern string   `xml:"char-name-pattern,attr"`
	Voice           string   `xml:"voice,attr"`
	XPos            float32  `xml:"x-pos,attr"`
	YPos            float32  `xml:"y-pos,attr"`
	XCenter         bool     `xml:"x-center,attr"`
	YCenter         bool     `xml:"y-center,attr"`
	InnerTxt        string   `xml:",innerxml"`
}

// Choices represents the choices elements in the vnscript file.
type Choices struct {
	XMLName xml.Name `xml:"choices"`
	Entries []Choice `xml:"choice"`
}

// Choice represents the choice elements in the vnscript file.
type Choice struct {
	XMLName    xml.Name `xml:"choice"`
	ForwardsTo string   `xml:"forwards-to,attr"`
}
