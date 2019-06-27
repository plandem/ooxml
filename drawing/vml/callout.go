package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

type CalloutType string //enum

//CalloutDrop is direct mapping of ST_CalloutDrop
type CalloutDrop string //enum

//Callout is direct mapping of CT_Callout
type Callout struct {
	XMLName         xml.Name    `xml:"callout,omitempty" namespace:"o"`
	On              bool        `xml:"on,attr,omitempty"`
	Type            CalloutType `xml:"type,attr,omitempty"`
	Gap             float64     `xml:"gap,attr,omitempty"`
	Angle           string      `xml:"angle,attr,omitempty"`
	DropAuto        bool        `xml:"dropauto,attr,omitempty"`
	Drop            CalloutDrop `xml:"drop,attr,omitempty"`
	Distance        string      `xml:"distance,attr,omitempty"`
	LengthSpecified bool        `xml:"lengthspecified,attr,omitempty"`
	Length          float64     `xml:"length,attr,omitempty"`
	AccentBar       bool        `xml:"accentbar,attr,omitempty"`
	TextBorder      *bool       `xml:"textborder,attr,omitempty"`
	MinusX          bool        `xml:"minusx,attr,omitempty"`
	MinusY          bool        `xml:"minusy,attr,omitempty"`
	Ext             Ext         `xml:"ext,attr,omitempty" namespace:"v"`
}

func (s *Callout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
