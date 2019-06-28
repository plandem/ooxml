package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//oval is direct mapping of CT_Oval
type oval struct {
	XMLName xml.Name `xml:"oval"`
	ml.ReservedAttributes
	coreAttributes
	shapeAttributes
	shapeElements
}

//Oval creates a new object with default values
func Oval() *oval {
	return &oval{}
}

func (s *oval) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.Attrs)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
