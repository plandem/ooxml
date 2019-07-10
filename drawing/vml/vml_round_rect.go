// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
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
	s.ReservedAttributes.ResolveNamespacePrefixes()
	s.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
