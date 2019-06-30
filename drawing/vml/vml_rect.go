package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//rect is direct mapping of CT_Rect
type rect struct {
	XMLName xml.Name `xml:"rect"`
	shapeAttributes
	shapeElements
}

//Rect creates a new object with default values
func Rect() *rect {
	return &rect{}
}

func (s *rect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	resolveElementsName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
