package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//ShapeType is direct mapping of CT_ShapeType
type ShapeType struct {
	XMLName xml.Name `xml:"shapetype"`
	Path    string   `xml:"path,attr,omitempty"`
	Adj     string   `xml:"adj,attr,omitempty"`
	Complex *Complex `xml:"complex,omitempty"`
	ml.ReservedAttributes
	coreAttributes
	shapeAttributes
	shapeElements
}

func (s *ShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.Attrs)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
