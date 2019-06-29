package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
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
	resolveAttributesName(s.ReservedAttributes)
	resolveElementsName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
