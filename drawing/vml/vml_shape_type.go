package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//ShapeType is direct mapping of CT_ShapeType
type ShapeType struct {
	XMLName xml.Name `xml:"shapetype"`
	Path    string   `xml:"path,attr,omitempty"`
	shapeAttributes
	shapeElements
}

func (s *ShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	resolveNestedName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
