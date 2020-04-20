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
	"strings"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains a basic interpreter of the xml-based scripting language used by the engine called 'vnscript'.

// newInterpreter is a simple constructor for the interpreter struct.
func newInterpreter(scenarioPath string, scenarioID int) (i interpreter, err error) {
	c, err := loadChapter(scenarioPath)
	i.currChapter = &c
	for _, v := range c.Scenarios {
		if v.ID == scenarioID {
			i.currScenario = &v
		}
	}
	if i.currScenario == nil {
		err = fmt.Errorf("critical error - could not find the specified scenario")
	}
	i.currEntryID = -1
	return
}

// interpreter is an entity responsible for loading the data from the VNgine scripting language to corresponding models.
type interpreter struct {
	currChapter  *Chapter
	currScenario *Scenario
	currEntryID  int
}

// nextEntry is a function that returns the next entry model.
func (i *interpreter) nextEntry() (em Entry) {
	//// If it is the end of the scenario then handle scenario switching
	//if i.currEntryID+1 == len(i.currScenario.Entries) {
	//	i.currEntryID = -1
	//	rPath := i.currScenario.Entries[i.currEntryID].RedirectPath
	//	if rPath != "" {
	//		fp, sc := scenarioPathToFilePath(rPath)
	//		// if the next scenario is in a different chapter then load it
	//		if fp != "" {
	//			// Load up the different chapter
	//			cm, err := GetChapterFromFile(fp)
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//			i.currChapter = &cm
	//		}
	//		// Search for the scenario in the current chapter
	//		for _, v := range i.currChapter.Scenarios {
	//			if v.Name == sc {
	//				i.currScenario = &v
	//			}
	//		}
	//	}
	//}
	//// Iterate to the next entry and return it
	//i.currEntryID++
	//em = i.currScenario.Entries[i.currEntryID]
	return
}

// scpathToFpath is a function that extracts the file path from the scenario path.
func scpathToFpath(sp string) (path string, err error) {
	// Tokenize the file path and if it's too short then return an error
	parts := strings.Split(sp, "/")
	if len(parts) < 2 {
		err = fmt.Errorf("invalid format of the scenario path: '%s'", sp)
		return
	}
	// Extract the filepath
	var sb strings.Builder
	sb.WriteString(parts[0])
	for i := 1; i < len(parts)-1; i++ {
		sb.WriteString("/")
		sb.WriteString(parts[i])
	}
	sb.WriteString(".vnscript")
	path = sb.String()
	return
}

// loadChapter loads the chapter from the specified chapter path.
func loadChapter(path string) (ch Chapter, err error) {
	// Extract the file path from the scenario path
	var fp string
	fp, err = scpathToFpath(path)
	if err != nil {
		return
	}

	// Read the file
	var script []byte
	script, err = ioutil.ReadFile(fp)

	// Unmarshall it into the chapter
	err = xml.Unmarshal(script, &ch)
	return
}
