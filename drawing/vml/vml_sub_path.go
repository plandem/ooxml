package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

//Path is direct mapping of CT_Path
type Path struct {
	XMLName         xml.Name        `xml:"path"`
	Value           string          `xml:"v,attr,omitempty"`
	ConnectType     ConnectType     `xml:"connecttype,attr,omitempty"`
	FillOK          ml.TriStateType `xml:"fillok,attr,omitempty"`
	StrokeOK        ml.TriStateType `xml:"strokeok,attr,omitempty"`
	ShadowOK        ml.TriStateType `xml:"shadowok,attr,omitempty"`
	ArrowOK         ml.TriStateType `xml:"arrowok,attr,omitempty"`
	GradientShapeOK ml.TriStateType `xml:"gradientshapeok,attr,omitempty"`
	ml.ReservedAttributes
}

func (s *Path) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
