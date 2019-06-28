package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//image is direct mapping of CT_Image
type image struct {
	XMLName xml.Name `xml:"image"`
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
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
