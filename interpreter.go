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
	"strconv"
	"strings"
	"unicode"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////// SHORT DESCRIPTION
// This file contains a basic interpreter of the xml-based scripting language used by the engine called 'vnscript'.

// newInterpreter is a simple constructor for the interpreter struct.
func newInterpreter(scpath string) (i interpreter, err error) {
	// Extract the scenario id
	split := strings.Split(scpath, "/")
	scID, err := strconv.Atoi(split[len(split)-1])
	if err != nil {
		return
	}

	// Load the chapter
	c, err := loadChapter(scpath)
	if err != nil {
		return
	}

	// Make a pointer to the chapter
	i.chapter = &c

	// Make also sure that the scenario id is not out of bounds
	if scID >= len(i.chapter.Scenarios) || scID < 0 {
		err = fmt.Errorf("could not find the scenario with '%d' id in scpath '%s'", scID, scpath)
	}
	return
}

// interpreter is an entity responsible for loading the data from the VNgine scripting language to corresponding models.
type interpreter struct {
	chapter *Chapter
	scID    int
	enID    int
	fwdPath string
}

// nextEntry is a function that returns the next entry model.
func (i *interpreter) nextEntry() (em Entry, err error) {
	// If it is the end of the scenario, load the next one
	if i.enID == len(i.chapter.Scenarios[i.scID].Entries) {
		err = i.loadNextScenario()
		if err != nil {
			return
		}
	}
	// Load the entry
	em = i.chapter.Scenarios[i.scID].Entries[i.enID]
	// schedule redirect if it has it
	if em.ForwardTo != "" {
		i.fwdPath = em.ForwardTo
	}
	// Increment the entryID for the next reading
	i.enID++
	return
}

// loadNextScenario loads up the next scenario for the interpreter and if the scenario is outside the current chapter it
// also loads the next chapter.
func (i *interpreter) loadNextScenario() (err error) {
	// if there was a forwarding scheduled then handle that
	if i.fwdPath != "" {
		var ch Chapter
		ch, err = loadChapter(i.fwdPath)
		if err != nil {
			// Check if the error is caused by the fact that the forward is towards a scenario inside
			// the current chapter
			if strings.HasPrefix(err.Error(), "invalid format of the scenario path") {
				parts := strings.Split(i.fwdPath, "/")
				if len(parts) == 1 {
					i.scID, err = strconv.Atoi(i.fwdPath)
				}
				if err != nil { // If the forward path is still incorrect then exit
					err = fmt.Errorf("incorrect format of the scenario id '%s'", i.fwdPath)
					return
				}
			} else { // any other error should cause the function to end
				return
			}
		} else {
			i.chapter = &ch
		}
		i.enID = 0
		i.fwdPath = ""
	} else { // if there was no forwarding scheduled then go onto the scenario with an id that's higher
		i.scID++
		if len(i.chapter.Scenarios) == i.scID {
			err = fmt.Errorf("no forward attribute in the last entry and no next scenario in this chapter")
			return
		}
	}
	return
}

// loadChapter loads the chapter from the specified scenario path.
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

	// Interpret the texts
	for i := 0; i < len(ch.Scenarios); i++ {
		for j := 0; j < len(ch.Scenarios[i].Entries); j++ {
			for k := 0; k < len(ch.Scenarios[i].Entries[j].Texts); k++ {
				// TODO: Benchmark the string to pointer vs this approach (memory usage and performance)
				ch.Scenarios[i].Entries[j].Texts[k].InnerTxt = contToRaw(ch.Scenarios[i].Entries[j].Texts[k].InnerTxt)
			}
		}
	}

	return
}

// scpathToFpath is a function that extracts the file path from the scenario path.
func scpathToFpath(scpath string) (path string, err error) {
	// Tokenize the file path and if it's too short then return an error
	parts := strings.Split(scpath, "/")
	if len(parts) < 2 {
		err = fmt.Errorf("invalid format of the scenario path: '%s'", scpath)
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

// contToRaw is a function that changes the raw string with vnscript tags into raw UTF-8 text.
func contToRaw(content string) (nc string) {
	// Function that deletes spaces at the start of all the lines and new line characters
	// where d is the character that represents the end of the line and str is the passed string
	rmOuterSpaces := func(str, d string) string {
		var sb strings.Builder
		if strings.Contains(str, d) {
			splitter := strings.Split(str, d)
			for _, v := range splitter {
				i, j := -1, -1
				for ii, vv := range v {
					if !unicode.IsSpace(vv) {
						i = ii
						break
					}
				}
				for jj := len(v)-1; jj > 0; jj-- {
					if !unicode.IsSpace([]rune(v)[jj]) {
						j = jj+1
						break
					}
				}
				if i != -1 && j != -1 {
					sb.WriteString(v[i:j])
				}
			}
		}
		return sb.String()
	}

	// Delete spaces and newline characters
	if strings.Contains(content, "\r\n") {
		nc = rmOuterSpaces(content, "\r\n")
	} else if strings.Contains(content, "\n") {
		nc = rmOuterSpaces(content, "\n")
	} else {
		nc = rmOuterSpaces(content, "\r")
	}

	// Replace <br/> to new line characters
	nc = strings.Replace(nc, "<br/>", "\r\n", -1)
	return
}
