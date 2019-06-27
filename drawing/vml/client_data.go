package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//ClientData is direct mapping of CT_ClientData
type ClientData struct {
	XMLName         xml.Name    `xml:"ClientData,omitempty" namespace:"x"`
}

func (s *ClientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
