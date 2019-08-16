// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import "github.com/plandem/ooxml/ml"

//Geometry is a direct mapping of XSD EG_Geometry
type Geometry struct {
	Custom *CustomGeometry2D `xml:"custGeom,omitempty"`
	Preset *PresetGeometry2D `xml:"prstGeom,omitempty"`
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