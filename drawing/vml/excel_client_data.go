package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//ClientData is direct mapping of CT_ClientData
type ClientData struct {
	XMLName xml.Name `xml:"ClientData"`
	ml.ReservedAttributes
	ml.ReservedElements
}

func (s *ClientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	resolveElementsName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceExcelPrefix, start.Name)})
}
