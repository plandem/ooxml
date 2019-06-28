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
	coreAttributes
	shapeAttributes
	imageAttributes
	shapeElements
}

//Image creates a new object with default values
func Image() *image {
	return &image{}
}

func (s *image) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.Attrs)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
