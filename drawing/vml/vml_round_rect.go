package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/drawing/vml/css"
)

//roundRect is direct mapping of CT_RoundRect
type roundRect struct {
	XMLName xml.Name     `xml:"roundrect"`
	ArcSize css.Fraction `xml:"arcsize,attr,omitempty"`
	shapeAttributes
	shapeElements
}

//RoundRect creates a new object with default values
func RoundRect() *roundRect {
	return &roundRect{
		ArcSize: 0.2,
	}
}

func (s *roundRect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	resolveElementsName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
