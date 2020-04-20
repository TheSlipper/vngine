package vngine

import (
	"fmt"
	"testing"
)

// TestScenarioParsing the parsing of a correctly formatted vngine script file.
func TestScenarioParsing(t *testing.T) {
	filePath := "source examples/chapter/0"
	x, err := loadChapter(filePath)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(x)
}
