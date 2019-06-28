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
	XMLName       xml.Name      `xml:"fill"`
	AlignShape    *bool         `xml:"alignshape,attr,omitempty"`
	Angle         *int          `xml:"angle,attr,omitempty"`
	Aspect        ImageAspect   `xml:"aspect,attr,omitempty"`
	Color         string        `xml:"color,attr,omitempty"`
	Color2        string        `xml:"color2,attr,omitempty"`
	Colors        string        `xml:"colors,attr,omitempty"`
	Focus         string        `xml:"focus,attr,omitempty"`
	FocusPosition string        `xml:"focusposition,attr,omitempty"`
	FocusSize     string        `xml:"focussize,attr,omitempty"`
	ID            string        `xml:"id,attr,omitempty"`
	Method        FillMethod    `xml:"method,attr,omitempty"`
	On            *bool         `xml:"on,attr,omitempty"`
	Opacity       string        `xml:"opacity,attr,omitempty"`
	Origin        string        `xml:"origin,attr,omitempty"`
	Position      string        `xml:"position,attr,omitempty"`
	Recolor       *bool         `xml:"recolor,attr,omitempty"`
	Rotate        *bool         `xml:"rotate,attr,omitempty"`
	Size          string        `xml:"size,attr,omitempty"`
	Src           string        `xml:"src,attr,omitempty"`
	Type          FillType      `xml:"type,attr,omitempty"`
	Extended      *FillExtended `xml:"fill,omitempty"`
	ml.ReservedAttributes
}

func (s *Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.Attrs)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
