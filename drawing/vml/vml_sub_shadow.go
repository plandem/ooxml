package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//ShadowType is direct mapping of ST_ShadowType
type ShadowType string //enum

//Shadow is direct mapping of CT_Shadow
type Shadow struct {
	XMLName  xml.Name   `xml:"shadow"`
	ID       string     `xml:"id,attr,omitempty"`
	On       bool       `xml:"on,attr,omitempty"`
	Type     ShadowType `xml:"type,attr,omitempty"`
	Color    string     `xml:"color,attr,omitempty"`
	Color2   string     `xml:"color2,attr,omitempty"`
	Obscured bool       `xml:"obscured,attr,omitempty"`
	Opacity  string     `xml:"opacity,attr,omitempty"`
	Offset   string     `xml:"offset,attr,omitempty"`
	Offset2  string     `xml:"offset2,attr,omitempty"`
	Origin   string     `xml:"origin,attr,omitempty"`
	Matrix   string     `xml:"matrix,attr,omitempty"`
	ml.ReservedAttributes
}

func (s *Shadow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
