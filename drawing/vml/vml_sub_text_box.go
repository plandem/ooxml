package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//TextBox is direct mapping of CT_TextBox
type TextBox struct {
	XMLName   xml.Name      `xml:"textbox"`
	Style     string        `xml:"style,attr,omitempty"`
	Inset     string        `xml:"inset,attr,omitempty"`
	InsetMode InsetModeType `xml:"insetmode,attr,omitempty"`
	Text      string        `xml:",innerxml"`
	ml.ReservedAttributes
}

func (s *TextBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
