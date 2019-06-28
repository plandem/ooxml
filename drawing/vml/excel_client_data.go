package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//ClientData is direct mapping of CT_ClientData
type ClientData struct {
	XMLName xml.Name `xml:"ClientData,omitempty"`
}

func (s *ClientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix("x", start.Name)})
}
