package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//polyLine is direct mapping of CT_PolyLine
type polyLine struct {
	XMLName xml.Name `xml:"polyline"`
	Points  string   `xml:"points,attr,omitempty"`
	shapeAttributes
	shapeElements
}

//PolyLine creates a new object with default values
func PolyLine() *polyLine {
	return &polyLine{
		Points: "0,0 10,10",
	}
}

func (s *polyLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	resolveElementsName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
