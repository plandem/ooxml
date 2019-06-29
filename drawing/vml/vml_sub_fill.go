package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//FillType is direct mapping of ST_FillType
type FillType string //enum

//FillMethod is direct mapping of ST_FillMethod
type FillMethod string //enum

//Fill is direct mapping of CT_Fill
type Fill struct {
	XMLName       xml.Name   `xml:"fill"`
	Color         string     `xml:"color,attr,omitempty"`
	Color2        string     `xml:"color2,attr,omitempty"`
	Colors        string     `xml:"colors,attr,omitempty"`
	Focus         string     `xml:"focus,attr,omitempty"`
	FocusPosition string     `xml:"focusposition,attr,omitempty"`
	FocusSize     string     `xml:"focussize,attr,omitempty"`
	Method        FillMethod `xml:"method,attr,omitempty"`
	Opacity       string     `xml:"opacity,attr,omitempty"`
	Origin        string     `xml:"origin,attr,omitempty"`
	Position      string     `xml:"position,attr,omitempty"`
	Size          string     `xml:"size,attr,omitempty"`
	Type          FillType   `xml:"type,attr,omitempty"`
	ml.ReservedElements
	ml.ReservedAttributes
}

func (s *Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	resolveNestedName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
