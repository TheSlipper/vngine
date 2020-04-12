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
	"log"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains a basic interpreter of the xml-based scripting language used by the engine called 'vnscript'.

// newInterpreter is a simple constructor for the interpreter struct.
func newInterpreter(scenarioPath string) (i interpreter, err error) {
	fp, sc := scenarioPathToFilePath(scenarioPath)
	cm, err := GetChapterFromFile(fp)
	if err != nil {
		return
	}
	i.currChapter = &cm
	for _, v := range cm.Scenarios {
		if v.Name == sc {
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
	currChapter  *ChapterModel
	currScenario *ScenarioModel
	currEntryID  int
}

// nextEntry is a function that returns the next entry model.
func (i *interpreter) nextEntry() (em EntryModel) {
	// If it is the end of the scenario then handle scenario switching
	if i.currEntryID+1 == len(i.currScenario.Entries) {
		i.currEntryID = -1
		rPath := i.currScenario.Entries[i.currEntryID].RedirectPath
		if rPath != "" {
			fp, sc := scenarioPathToFilePath(rPath)
			// if the next scenario is in a different chapter then load it
			if fp != "" {
				// Load up the different chapter
				cm, err := GetChapterFromFile(fp)
				if err != nil {
					log.Fatal(err)
				}
				i.currChapter = &cm
			}
			// Search for the scenario in the current chapter
			for _, v := range i.currChapter.Scenarios {
				if v.Name == sc {
					i.currScenario = &v
				}
			}
		}
	}
	// Iterate to the next entry and return it
	i.currEntryID++
	em = i.currScenario.Entries[i.currEntryID]
	return
}

// scenarioPathToFilePath is a function that extracts the file path and scenario name from the scenario path.
func scenarioPathToFilePath(sp string) (fp, sc string) {
	parts := strings.Split(sp, "/")
	sb := strings.Builder{}
	sb.WriteString(parts[0])
	for i := 1; i < len(parts)-1; i++ {
		sb.WriteString("/")
		sb.WriteString(parts[i])
	}
	sc = parts[len(parts)-1]
	fp = sb.String() + ".vnscript"
	return
}
