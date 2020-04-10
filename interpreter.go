package vngine

import (
	"fmt"
	"log"
	"strings"
)

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
	for i := 1; i < len(parts)-2; i++ {
		sb.WriteString("/")
		sb.WriteString(parts[i])
	}
	sc = parts[len(parts)-1]
	fp = sb.String() + ".vnscript"
	return
}
