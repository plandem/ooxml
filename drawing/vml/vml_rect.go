package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//rect is direct mapping of CT_Rect
type rect struct {
	XMLName xml.Name `xml:"rect"`
	ml.ReservedAttributes
	coreAttributes
	shapeAttributes
	shapeElements
}

//Rect creates a new object with default values
func Rect() *rect {
	return &rect{}
}

func (s *rect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.Attrs)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
