package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//PathConnectType is direct mapping of ST_ConnectType
type PathConnectType string //enum

//Path is direct mapping of CT_Path
type Path struct {
	XMLName         xml.Name        `xml:"path" namespace:"v"`
	ID              string          `xml:"id,attr,omitempty"`
	V               string          `xml:"v,attr,omitempty"`
	Limo            string          `xml:"limo,attr,omitempty"`
	TextBoxRect     string          `xml:"textboxrect,attr,omitempty"`
	ConnectType     PathConnectType `xml:"connecttype,attr,omitempty" namespace:"o"`
	ConnectLocs     string          `xml:"connectlocs,attr,omitempty" namespace:"o"`
	ConnectAngles   string          `xml:"connectangles,attr,omitempty" namespace:"o"`
	FillOK          *bool           `xml:"fillok,attr,omitempty"`
	StrokeOK        *bool           `xml:"strokeok,attr,omitempty"`
	ShadowOK        *bool           `xml:"shadowok,attr,omitempty"`
	ArrowOK         *bool           `xml:"arrowok,attr,omitempty"`
	GradientShapeOK *bool           `xml:"gradientshapeok,attr,omitempty"`
	TextpathOK      *bool           `xml:"textpathok,attr,omitempty"`
	InsetPenOK      *bool           `xml:"insetpenok,attr,omitempty"`
	ExtrusionOK     *bool           `xml:"extrusionok,attr,omitempty" namespace:"o"`
}

func (s *Path) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
