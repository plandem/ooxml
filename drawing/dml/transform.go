// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Transform2D is a direct mapping of XSD CT_Transform2D
type Transform2D struct {
	Offset         *Point2D        `xml:"off,omitempty"`
	Size           *PositiveSize2D `xml:"ext,omitempty"`
	FlipHorizontal bool            `xml:"flipH,attr,omitempty"`
	FlipVertical   bool            `xml:"flipV,attr,omitempty"`
	Rotation       ml.PropertyInt  `xml:"rot,attr,omitempty"`
}

type transform2D struct {
	Offset         *Point2D        `xml:"a:off,omitempty"`
	Size           *PositiveSize2D `xml:"a:ext,omitempty"`
	FlipHorizontal bool            `xml:"flipH,attr,omitempty"`
	FlipVertical   bool            `xml:"flipV,attr,omitempty"`
	Rotation       ml.PropertyInt  `xml:"rot,attr,omitempty"`
}

func (t *Transform2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(transform2D(*t), start)
}