package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//image is direct mapping of CT_Image
type image struct {
	XMLName xml.Name `xml:"image"`
	ml.ReservedAttributes
	imageAttributes
	shapeAttributes
	shapeElements
}

//Image creates a new object with default values
func Image() *image {
	return &image{}
}

func (s *image) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	resolveNestedName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
