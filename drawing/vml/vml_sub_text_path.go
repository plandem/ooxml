package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//TextPath is direct mapping of CT_TextPath
type TextPath struct {
	XMLName  xml.Name        `xml:"textpath"`
	Style    string          `xml:"style,attr,omitempty"`
	Text     string          `xml:"string,attr,omitempty"`
	On       ml.TriStateType `xml:"on,attr,omitempty"`
	FitShape ml.TriStateType `xml:"fitshape,attr,omitempty"`
	FitPath  ml.TriStateType `xml:"fitpath,attr,omitempty"`
	Trim     ml.TriStateType `xml:"trim,attr,omitempty"`
	XScale   ml.TriStateType `xml:"xscale,attr,omitempty"`
	ml.ReservedAttributes
}

func (s *TextPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
