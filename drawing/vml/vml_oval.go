package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//oval is direct mapping of CT_Oval
type oval struct {
	XMLName xml.Name `xml:"oval"`
	shapeAttributes
	shapeElements
}

//Oval creates a new object with default values
func Oval() *oval {
	return &oval{}
}

func (s *oval) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	s.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
