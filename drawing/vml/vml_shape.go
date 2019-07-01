package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Shape is direct mapping of CT_Shape
type Shape struct {
	XMLName xml.Name `xml:"shape"`
	Type    string   `xml:"type,attr,omitempty"`
	ShapeType
}

func (s *Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
