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

//ShapeProperties is a direct mapping of XSD CT_ShapeProperties
type ShapeProperties struct {
	Transform      *Transform2D    `xml:"xfrm,omitempty"`
	LineProperties *LineProperties `xml:"ln,omitempty"`
	Mode           BlackWhiteMode  `xml:"bwMode,attr,omitempty"`
	*Geometry
	*FillProperties
	ml.ReservedElements
}

//Go1.12 has limited support of namespace prefixes, so use special type with hardcoded prefixes for marshalling
type shapeProperties struct {
	Transform      *Transform2D    `xml:"a:xfrm,omitempty"`
	LineProperties *LineProperties `xml:"a:ln,omitempty"`
	Mode           BlackWhiteMode  `xml:"bwMode,attr,omitempty"`
	*Geometry
	*FillProperties
	ml.ReservedElements
}

func (t *ShapeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(shapeProperties(*t), start)
}
