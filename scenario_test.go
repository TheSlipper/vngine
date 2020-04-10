package vnginelib

import (
	"testing"
)

// TestScenarioParsing the parsing of a correctly formatted vngine script file.
func TestScenarioParsing(t *testing.T) {
	filePath := "source examples/chapter.xml"
	_, err := GetChapterFromFile(filePath)
	if err != nil {
		panic(err)
	}
	_, err = GetScenarioFromFileByID(0, filePath)
	if err != nil {
		panic(err)
	}
	_, err = GetScenarioFromFileByName("introduction_5", filePath)
	if err != nil {
		panic(err)
	}
}
