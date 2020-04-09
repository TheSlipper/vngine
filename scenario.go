package vnginelib

import "encoding/xml"

// ChapterModel represents a chapter in a novel scripting file.
type ChapterModel struct {
	XMLName   xml.Name        `xml:"chapter"`
	ID        int             `xml:"id,attr"`
	Name      string          `xml:"name,attr"`
	Scenarios []ScenarioModel `xml:"scenario"`
}

// ScenarioModel represents a part of a chapter.
type ScenarioModel struct {
	XMLName xml.Name     `xml:"scenario"`
	ID      int          `xml:"id,attr"`
	Name    string       `xml:"name,attr"`
	Entries []EntryModel `xml:"entry"`
}

// EntryModel represents a single frame in the novel.
type EntryModel struct {
	XMLName         xml.Name         `xml:"entry"`
	ID              int              `xml:"id,attr"`
	TransitionStyle string           `xml:"transition-style,attr"`
	Characters      []CharacterModel `xml:"character"`
	MusicEvents     []MusicModel     `xml:"music"`
	Sounds          []SoundModel     `xml:"sound"`
	Text            TextModel        `xml:"text"`
	ChoiceBox       ChoiceBoxModel   `xml:"choice-box"`
}

// MusicModel represents a music event entry that occurs in a given frame.
type MusicModel struct {
	XMLName            xml.Name `xml:"music"`
	ActionName         string   `xml:"action,attr"`
	AppliedEffect      string   `xml:"effect,attr"`
	ChannelName        string   `xml:"channel-name,attr"`
	MusicName          string   `xml:"name,attr"`
	StartTimestamp     string   `xml:"start,attr"`
	Loops              bool     `xml:"loop,attr"`
	LoopStartTimestamp string   `xml:"loop-start,attr"`
	LoopEndTimestamp   string   `xml:"loop-end,attr"`
	StartEffect        string   `xml:"start-effect,attr"`
	EndEffect          string   `xml:"end-effect,attr"`
}

// SoundModel represents a sound that occurs in a given frame.
type SoundModel struct {
	Name   string `xml:"name,attr"`
	Repeat bool   `xml:"repeat,attr"`
	Delay  uint8  `xml:"delay,attr"`
}

// EffectModel represents a definition of either a visual effect or an animation that will take place in the given frame.
type EffectModel struct {
	XMLName      xml.Name `xml:"effect"`
	ID           int      `xml:"id,attr"`
	Asynchronous bool     `xml:"asynchronous,attr"`
	Repeat       bool     `xml:"repeat,attr"`
	QueueIndex   uint8    `xml:"queue,attr"`
}

// CharacterModel represents a setup of the character on the screen in a given frame.
type CharacterModel struct {
	XMLName   xml.Name `xml:"character"`
	ID        int      `xml:"id,attr"`
	State     string   `xml:"state,attr"`
	Blinking  bool     `xml:"blinking,attr"`
	PositionX uint8    `xml:"position-x,attr"`
	PositionY uint8    `xml:"position-y,attr"`
	Priority  uint8    `xml:"priority,attr"`
}

// TextModel represents text data that will be displayed in the given novel's frame.
type TextModel struct {
	XMLName   xml.Name `xml:"text"`
	SpeakerID int      `xml:"speaker-id,attr"`
	Content   string   `xml:",chardata"`
}

// ChoiceBoxModel represents all the data that the user will be able to choose.
type ChoiceBoxModel struct {
	XMLName xml.Name      `xml:"choice-box"`
	Choices []ChoiceModel `xml:"choice"`
}

// ChoiceModel contains information on a single choice that a user will be able to choose.
type ChoiceModel struct {
	XMLName      xml.Name `xml:"choice"`
	RedirectPath string   `xml:"redirect,attr"`
	Value        string   `xml:"value,attr"`
	Content      string   `xml:",chardata"`
}
