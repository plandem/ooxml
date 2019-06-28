package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//oval is direct mapping of CT_Oval
type oval struct {
	XMLName xml.Name `xml:"oval"`
	coreAttributes
	shapeAttributes
	shapeElements
}

//Oval creates a new object with default values
func Oval() *oval {
	return &oval{}
}

func (s *oval) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
