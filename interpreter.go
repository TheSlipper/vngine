package vngine

import (
	"fmt"
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
	return
}

// interpreter is an entity responsible for loading the data from the VNgine scripting language to corresponding models.
type interpreter struct {
	currChapter  *ChapterModel
	currScenario *ScenarioModel
}

// scenarioPathToFilePath is a function that extracts the file path and scenario name from the scenario path.
func scenarioPathToFilePath(sp string) (fp, sc string) {
	parts := strings.Split(sp, "/")
	sb := strings.Builder{}
	for i := 0; i < len(parts)-1; i++ {
		sb.WriteString(parts[i])
	}
	sc = parts[len(parts)-1]
	fp = sb.String() + ".vnscript"
	return
}
