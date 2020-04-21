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
// scripting language. Models with the 'S' prefix are used in vnscripts file and
// models with the 'A' prefix are used in vnassets file.

////////////////////////////////////////////////////////////// VNSCRIPT MODELS

// SChapter represents the chapter element in the vnscript file.
type SChapter struct {
	XMLName   xml.Name    `xml:"chapter"`
	Title     string      `xml:"title,attr"`
	Scenarios []SScenario `xml:"scenario"`
}

// SScenario represents the scenario elements in the vnscript file.
type SScenario struct {
	XMLName xml.Name   `xml:"scenario"`
	ID      int        `xml:"id,attr"`
	Assets  SAssetPath `xml:"asset-path"`
	Entries []SEntry   `xml:"entry"`
}

// SAssetPath represents the asset-path elements in the vnscript file.
type SAssetPath struct {
	XMLName xml.Name `xml:"asset-path"`
	Source  string   `xml:"src,attr"`
}

// SEntry represents the entry elements in the vnscript file.
type SEntry struct {
	XMLName          xml.Name      `xml:"entry"`
	ID               int           `xml:"id,attr"`
	Transmission     string        `xml:"transmission,attr"`
	Interactable     bool          `xml:"interactable,attr"`
	DeferInteraction bool          `xml:"defer-interaction,attr"`
	ForwardTo        string        `xml:"forward-to,attr"`
	BG               SBackground   `xml:"background"`
	MusicEvents      []SMusicEvent `xml:"music-event"`
	Chars            []SCharacter  `xml:"character"`
	Texts            []SText       `xml:"text"`
	Choices          SChoices      `xml:"choices"`
}

// SBackground represents the background elements in the vnscript file.
type SBackground struct {
	XMLName xml.Name `xml:"background"`
	Name    string   `xml:"name,attr"`
}

// SMusicEvent represents the music-event elements in the vnscript file.
type SMusicEvent struct {
	XMLName     xml.Name `xml:"music-event"`
	Action      string   `xml:"action,attr"`
	Track       string   `xml:"track,attr"`
	Repeatable  bool     `xml:"repeatable,attr"`
	EffectNames string   `xml:"effect-names,attr"`
}

// SCharacter represents the character elements in the vnscript file.
type SCharacter struct {
	XMLName xml.Name `xml:"character"`
	ID      int      `xml:"id,attr"`
	State   string   `xml:"state,attr"`
	XPos    float32  `xml:"x-pos,attr"`
	YPos    float32  `xml:"y-pos,attr"`
}

// SText represents the text elements in the vnscript file.
type SText struct {
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

// SChoices represents the choices elements in the vnscript file.
type SChoices struct {
	XMLName xml.Name  `xml:"choices"`
	Entries []SChoice `xml:"choice"`
}

// SChoice represents the choice elements in the vnscript file.
type SChoice struct {
	XMLName    xml.Name `xml:"choice"`
	ForwardsTo string   `xml:"forwards-to,attr"`
}

////////////////////////////////////////////////////////////// VNASSETS MODELS

type AAssets struct {

}

