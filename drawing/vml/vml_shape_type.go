package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//ShapeType is direct mapping of CT_ShapeType
type ShapeType struct {
	XMLName xml.Name `xml:"shapetype"`
	Master  string   `xml:"master,attr,omitempty" namespace:"o"`
	Path    string   `xml:"path,attr,omitempty"`
	Adj     string   `xml:"adj,attr,omitempty"`
	Complex *Complex `xml:"complex,omitempty"`
	coreAttributes
	shapeAttributes
	shapeElements
}

func (s *ShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
