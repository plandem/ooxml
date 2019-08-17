// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//GroupTransform2D is a direct mapping of XSD CT_GroupTransform2D
type GroupTransform2D struct {
	Offset         *Point2D        `xml:"off,omitempty"`
	Size           *PositiveSize2D `xml:"ext,omitempty"`
	ChildOffset    *Point2D        `xml:"chOff,omitempty"`
	ChildSize      *PositiveSize2D `xml:"chExt,omitempty"`
	FlipHorizontal bool            `xml:"flipH,attr,omitempty"`
	FlipVertical   bool            `xml:"flipV,attr,omitempty"`
	Rotation       ml.PropertyInt  `xml:"rot,attr,omitempty"`
}

//Go1.12 has limited support of namespace prefixes, so use special type with hardcoded prefixes for marshalling
type groupTransform2D struct {
	Offset         *Point2D        `xml:"a:off,omitempty"`
	Size           *PositiveSize2D `xml:"a:ext,omitempty"`
	ChildOffset    *Point2D        `xml:"a:chOff,omitempty"`
	ChildSize      *PositiveSize2D `xml:"a:chExt,omitempty"`
	FlipHorizontal bool            `xml:"flipH,attr,omitempty"`
	FlipVertical   bool            `xml:"flipV,attr,omitempty"`
	Rotation       ml.PropertyInt  `xml:"rot,attr,omitempty"`
}

func (t *GroupTransform2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(groupTransform2D(*t), start)
}