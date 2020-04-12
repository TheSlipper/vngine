package vngine

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// getSettingsModelFromFile parses the settings file and returns the settings model with data from it.
func getSettingsModelFromFile(file string) (s SettingsModel, err error) {
	var xmlData []byte
	xmlData, err = ioutil.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("error while trying to read from the file '%s': %+v", file, err)
		return
	}
	err = xml.Unmarshal(xmlData, &s)
	if err != nil {
		err = fmt.Errorf("error while trying to interpret settings file '%s': %+v", file, err)
	}
	return
}

// SettingsModel represents the settings file and its elements.
type SettingsModel struct {
	XMLName     xml.Name `xml:"settings"`
	Name        string   `xml:"name"`
	VSync       bool     `xml:"vsync"`
	Resizable   bool     `xml:"resizable"`
	Width       float64  `xml:"width"`
	Height      float64  `xml:"height"`
	Undecorated bool     `xml:"undecorated"`
	AlwaysOnTop bool     `xml:"ontop"`
	Icon        string   `xml:"icon"`
	Fullscreen  bool     `xml:"fullscreen"`
}
