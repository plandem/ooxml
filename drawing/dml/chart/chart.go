// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Chart is a direct mapping of XSD CT_Chart
type Chart struct {
	XMLName                   xml.Name         `xml:"chart"`
	Title                     *Title           `xml:"title,omitempty"`
	PlotArea                  *PlotArea        `xml:"plotArea,omitempty"`
	Legend                    *Legend          `xml:"legend,omitempty"`
	DisplayBlanksAs           *ml.Property     `xml:"dispBlanksAs,omitempty"`
	AutoTitleIsDeleted        *ml.PropertyBool `xml:"autoTitleDeleted,omitempty"` //default true
	PlotVisibleOnly           *ml.PropertyBool `xml:"plotVisOnly,omitempty"`      //default true
	ShowDataLabelsOverMaximum *ml.PropertyBool `xml:"showDLblsOverMax,omitempty"` //default true
	Floor                     *ml.Reserved     `xml:"floor,omitempty"`
	SideWall                  *ml.Reserved     `xml:"sideWall,omitempty"`
	BackWall                  *ml.Reserved     `xml:"backWall,omitempty"`
	PivotFmts                 *ml.Reserved     `xml:"pivotFmts,omitempty"`
	View3D                    *ml.Reserved     `xml:"view3D,omitempty"`
	ExtLst                    *ml.Reserved     `xml:"extLst,omitempty"`
}
