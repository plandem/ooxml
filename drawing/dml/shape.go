// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//BlackWhiteMode is a direct mapping of XSD ST_BlackWhiteMode
type BlackWhiteMode string //enum

//Shape is a direct mapping of XSD CT_ShapeProperties
type Shape struct {
	Transform *Transform2D   `xml:"xfrm,omitempty"`
	Line      *Line          `xml:"ln,omitempty"`
	Mode      BlackWhiteMode `xml:"bwMode,attr,omitempty"`
	ml.ReservedElements
	geometry
	fill
}

//Go1.12 has limited support of namespace prefixes, so use special type with hardcoded prefixes for marshalling
type shape struct {
	Transform *Transform2D   `xml:"a:xfrm,omitempty"`
	Line      *Line          `xml:"a:ln,omitempty"`
	Mode      BlackWhiteMode `xml:"bwMode,attr,omitempty"`
	ml.ReservedElements
	geometry
	fill
}

func (n *Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(shape(*n), start)
}
