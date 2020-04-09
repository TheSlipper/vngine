package vnginelib

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestScenarioParsing(t *testing.T) {
	chapter := ChapterModel {}
	xmlData, err := ioutil.ReadFile("source examples/chapter.xml")
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(xmlData, &chapter)
	if err != nil {
		panic(err)
	}
}
