package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
)

//TextPath is direct mapping of CT_TextPath
type TextPath struct {
	XMLName  xml.Name   `xml:"textpath"`
	ID       string     `xml:"id,attr,omitempty"`
	Style    *css.Style `xml:"style,attr,omitempty"`
	Text     string     `xml:"string,attr,omitempty"`
	On       bool       `xml:"on,attr,omitempty"`
	FitShape bool       `xml:"fitshape,attr,omitempty"`
	FitPath  bool       `xml:"fitpath,attr,omitempty"`
	Trim     *bool      `xml:"trim,attr,omitempty"`
	XScale   *bool      `xml:"xscale,attr,omitempty"`
	ml.ReservedAttributes
}

func (s *TextPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.Attrs)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
