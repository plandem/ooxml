// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
)

//Fill is direct mapping of CT_Fill
type Fill struct {
	XMLName         xml.Name        `xml:"fill"`
	Color           string          `xml:"color,attr,omitempty"`
	Color2          string          `xml:"color2,attr,omitempty"`
	Colors          string          `xml:"colors,attr,omitempty"`
	Focus           css.Fraction    `xml:"focus,attr,omitempty"`
	FocusPosition   string          `xml:"focusposition,attr,omitempty"`
	FocusSize       string          `xml:"focussize,attr,omitempty"`
	ImageAlignShape ml.TriStateType `xml:"alignshape,attr,omitempty"`
	ImageAspect     ImageAspect     `xml:"aspect,attr,omitempty"`
	ImageSize       string          `xml:"size,attr,omitempty"`
	ImageSrc        string          `xml:"src,attr,omitempty"`
	Method          FillMethod      `xml:"method,attr,omitempty"`
	On              ml.TriStateType `xml:"on,attr,omitempty"`
	Opacity         css.Fraction    `xml:"opacity,attr,omitempty"`
	Origin          string          `xml:"origin,attr,omitempty"`
	Position        string          `xml:"position,attr,omitempty"`
	Type            FillType        `xml:"type,attr,omitempty"`
	ml.ReservedElements
	ml.ReservedAttributes
}

func (s *Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	s.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
