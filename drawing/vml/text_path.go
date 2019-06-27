package vml

import (
	"encoding/xml"
	css2 "github.com/plandem/ooxml/drawing/vml/css"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//TextPath is direct mapping of CT_TextPath
type TextPath struct {
	XMLName  xml.Name   `xml:"textpath,omitempty" namespace:"v"`
	ID       string     `xml:"id,attr,omitempty"`
	Style    css2.Style `xml:"style,attr,omitempty"`
	Text     string     `xml:"string,attr,omitempty"`
	On       bool       `xml:"on,attr,omitempty"`
	FitShape bool       `xml:"fitshape,attr,omitempty"`
	FitPath  bool       `xml:"fitpath,attr,omitempty"`
	Trim     *bool      `xml:"trim,attr,omitempty"`
	XScale   *bool      `xml:"xscale,attr,omitempty"`
}

func (s *TextPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
