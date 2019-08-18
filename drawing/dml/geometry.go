// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//geometry is a direct mapping of XSD EG_Geometry
type geometry struct {
	CustomGeometry *CustomGeometry2D `xml:"custGeom,omitempty"`
	PresetGeometry *PresetGeometry2D `xml:"prstGeom,omitempty"`
}

//CustomGeometry2D is a direct mapping of XSD CT_CustomGeometry2D
type CustomGeometry2D struct {
	ml.ReservedElements
}

//PresetGeometry2D is a direct mapping of XSD CT_PresetGeometry2D
type PresetGeometry2D struct {
	Type TextShapeType `xml:"prst,attr,omitempty"`
	ml.ReservedElements
}

func (n *CustomGeometry2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDML, start.Name)})
}

func (n *PresetGeometry2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDML, start.Name)})
}
