// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Name is a helper type for marshaling namespace for DrawingML
type Name string

//Point2D is a direct mapping of XSD CT_Point2D
type Point2D struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}

//PositiveSize2D is a direct mapping of XSD CT_PositiveSize2D
type PositiveSize2D struct {
	Height uint `xml:"cx,attr"` //cx - height in EMU
	Width  uint `xml:"cy,attr"` //cy - width in EMU
}

func (n Name) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if prefix, ok := ml.ResolveNamespacePrefix(ml.NamespaceDML); ok {
		return xml.Attr{Name: xml.Name{Local: "xmlns:" + prefix}, Value: ml.NamespaceDML}, nil
	}

	return xml.Attr{}, ml.ErrorNamespace(ml.NamespaceDML)
}