package vml

import (
	"encoding/xml"
	css2 "github.com/plandem/ooxml/drawing/vml/css"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//TextBox is direct mapping of CT_TextBox
type TextBox struct {
	XMLName     xml.Name   `xml:"textbox,omitempty" namespace:"v"`
	ID          string     `xml:"id,attr,omitempty"`
	Style       css2.Style `xml:"style,attr,omitempty"`
	Inset       string     `xml:"inset,attr,omitempty"`
	SingleClick bool       `xml:"singleclick,attr,omitempty" namespace:"o"`
	InsetMode   string     `xml:"insetmode,attr,omitempty" namespace:"o"` //enum
	Text        string     `xml:",innerxml"`
}

func (s *TextBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
