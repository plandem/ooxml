package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//arc is alias of CT_Arc
type arc struct {
	XMLName    xml.Name `xml:"arc"`
	StartAngle int      `xml:"startAngle,attr,omitempty"`
	EndAngle   int      `xml:"endAngle,attr,omitempty"`
	shapeAttributes
	shapeElements
}

//Arc creates a new object with default values
func Arc() *arc {
	return &arc{
		StartAngle: 0,
		EndAngle:   90,
	}
}

func (s *arc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	s.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
