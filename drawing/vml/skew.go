package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//Skew is direct mapping of CT_Skew
type Skew struct {
	XMLName xml.Name `xml:"skew,omitempty" namespace:"o"`
	On      bool     `xml:"on,attr,omitempty"`
	ID      string   `xml:"id,attr,omitempty"`
	Offset  string   `xml:"offset,attr,omitempty"`
	Origin  string   `xml:"origin,attr,omitempty"`
	Matrix  string   `xml:"matrix,attr,omitempty"`
	Ext     Ext      `xml:"ext,attr,omitempty" namespace:"v"`
}

func (s *Skew) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
