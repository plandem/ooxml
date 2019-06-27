package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//Lock is direct mapping of CT_Lock
type Lock struct {
	XMLName       xml.Name `xml:"lock,omitempty" namespace:"o"`
	AdjustHandles bool     `xml:"adjusthandles,attr,omitempty"`
	AspectRatio   bool     `xml:"aspectratio,attr,omitempty"`
	Cropping      bool     `xml:"cropping,attr,omitempty"`
	Grouping      bool     `xml:"grouping,attr,omitempty"`
	Position      bool     `xml:"position,attr,omitempty"`
	Rotation      bool     `xml:"rotation,attr,omitempty"`
	Selection     bool     `xml:"selection,attr,omitempty"`
	ShapeType     bool     `xml:"shapetype,attr,omitempty"`
	Text          bool     `xml:"text,attr,omitempty"`
	Vertices      bool     `xml:"vertices,attr,omitempty"` //vertices(documentation) or verticies(XSD definition)
	Ungrouping    bool     `xml:"ungrouping,attr,omitempty"`
	Ext           Ext      `xml:"ext,attr,omitempty" namespace:"v"`
}

func (s *Lock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
