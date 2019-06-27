package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//ClipPath is direct mapping of CT_ClipPath
type ClipPath struct {
	XMLName xml.Name `xml:"clippath,omitempty" namespace:"o"`
	V       string   `xml:"v,attr,omitempty" namespace:"o"`
}

func (s *ClipPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
