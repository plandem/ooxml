package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//Complex is direct mapping of CT_Complex
type Complex struct {
	XMLName xml.Name `xml:"complex,omitempty" namespace:"o"`
	Ext     Ext      `xml:"ext,attr,omitempty" namespace:"v"`
}

func (s *Complex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
