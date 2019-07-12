// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
)

//StrokeAttributes is direct mapping of AG_StrokeAttributes
type StrokeAttributes struct {
	Color            string            `xml:"color,attr,omitempty"`
	Color2           string            `xml:"color2,attr,omitempty"`
	DashStyle        StrokeDashStyle   `xml:"dashstyle,attr,omitempty"`
	EndArrow         StrokeArrowType   `xml:"endarrow,attr,omitempty"`
	EndArrowLength   StrokeArrowLength `xml:"endarrowlength,attr,omitempty"`
	EndArrowWidth    StrokeArrowWidth  `xml:"endarrowwidth,attr,omitempty"`
	EndCap           StrokeEndCap      `xml:"endcap,attr,omitempty"`
	FillType         FillType          `xml:"filltype,attr,omitempty"`
	ImageAlignShape  ml.TriStateType   `xml:"imagealignshape,attr,omitempty"`
	ImageAspect      ImageAspect       `xml:"imageaspect,attr,omitempty"`
	ImageSize        string            `xml:"imagesize,attr,omitempty"`
	ImageSrc         string            `xml:"src,attr,omitempty"`
	InsetPenOK       ml.TriStateType   `xml:"insetpen,attr,omitempty"`
	JoinStyle        StrokeJoinStyle   `xml:"joinstyle,attr,omitempty"`
	LineStyle        StrokeLineStyle   `xml:"linestyle,attr,omitempty"`
	On               ml.TriStateType   `xml:"on,attr,omitempty"`
	Opacity          css.Fraction      `xml:"opacity,attr,omitempty"`
	StartArrowLength StrokeArrowLength `xml:"startarrowlength,attr,omitempty"`
	StartArrow       StrokeArrowType   `xml:"startarrow,attr,omitempty"`
	StartArrowWidth  StrokeArrowWidth  `xml:"startarrowwidth,attr,omitempty"`
	Weight           css.Number        `xml:"weight,attr,omitempty"`
	ml.ReservedAttributes
}

//Stroke is direct mapping of CT_Stroke
type Stroke struct {
	XMLName xml.Name          `xml:"stroke"`
	Left    *StrokeAttributes `xml:"left,omitempty"`
	Top     *StrokeAttributes `xml:"top,omitempty"`
	Right   *StrokeAttributes `xml:"right,omitempty"`
	Bottom  *StrokeAttributes `xml:"bottom,omitempty"`
	Column  *StrokeAttributes `xml:"column,omitempty"`
	StrokeAttributes
}

func (s *Stroke) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}

func (s *StrokeAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVMLOffice, start.Name)})
}
