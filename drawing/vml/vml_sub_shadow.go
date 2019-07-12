// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
)

//Shadow is direct mapping of CT_Shadow
type Shadow struct {
	XMLName  xml.Name        `xml:"shadow"`
	On       ml.TriStateType `xml:"on,attr,omitempty"`
	Type     ShadowType      `xml:"type,attr,omitempty"`
	Color    string          `xml:"color,attr,omitempty"`
	Color2   string          `xml:"color2,attr,omitempty"`
	Obscured ml.TriStateType `xml:"obscured,attr,omitempty"`
	Opacity  css.Fraction    `xml:"opacity,attr,omitempty"`
	Offset   string          `xml:"offset,attr,omitempty"`
	Offset2  string          `xml:"offset2,attr,omitempty"`
	Origin   string          `xml:"origin,attr,omitempty"`
	Matrix   string          `xml:"matrix,attr,omitempty"`
	ml.ReservedAttributes
}

func (s *Shadow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
