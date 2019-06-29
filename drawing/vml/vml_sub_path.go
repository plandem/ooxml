package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//Path is direct mapping of CT_Path
type Path struct {
	XMLName         xml.Name    `xml:"path"`
	ID              string      `xml:"id,attr,omitempty"`
	V               string      `xml:"v,attr,omitempty"`
	Limo            string      `xml:"limo,attr,omitempty"`
	TextBoxRect     string      `xml:"textboxrect,attr,omitempty"`
	ConnectType     ConnectType `xml:"connecttype,attr,omitempty" namespace:"o"`
	ConnectLocs     string      `xml:"connectlocs,attr,omitempty" namespace:"o"`
	ConnectAngles   string      `xml:"connectangles,attr,omitempty" namespace:"o"`
	FillOK          *bool       `xml:"fillok,attr,omitempty"`
	StrokeOK        *bool       `xml:"strokeok,attr,omitempty"`
	ShadowOK        *bool       `xml:"shadowok,attr,omitempty"`
	ArrowOK         *bool       `xml:"arrowok,attr,omitempty"`
	GradientShapeOK *bool       `xml:"gradientshapeok,attr,omitempty"`
	TextpathOK      *bool       `xml:"textpathok,attr,omitempty"`
	InsetPenOK      *bool       `xml:"insetpenok,attr,omitempty"`
	ExtrusionOK     *bool       `xml:"extrusionok,attr,omitempty" namespace:"o"`
	ml.ReservedAttributes
}

func (s *Path) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
