// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/ml/common"
	"github.com/plandem/ooxml/ml"
)

//Space is a direct mapping of XSD CT_ChartSpace
type Space struct {
	XMLName         xml.Name                `xml:"http://schemas.openxmlformats.org/drawingml/2006/chart chartSpace"`
	DrawingName     ml.DrawingName          `xml:",attr"`
	Lang            ml.Property             `xml:"lang,omitempty"`
	Style           ml.PropertyInt          `xml:"style,omitempty"`
	Date1904        *ml.PropertyBool        `xml:"date1904,omitempty"`       //default true
	RoundedCorners  *ml.PropertyBool        `xml:"roundedCorners,omitempty"` //default true
	ColorMapping    *common.ColorMapping    `xml:"clrMapOvr,omitempty"`
	ShapeProperties *common.ShapeProperties `xml:"spPr,omitempty"`
	TextBody        *common.TextBody        `xml:"txPr,omitempty"`
	Chart           *Chart                  `xml:"chart"`
	PivotSource     *ml.Reserved            `xml:"pivotSource,omitempty"`
	Protection      *ml.Reserved            `xml:"protection,omitempty"`
	ExternalData    *ml.Reserved            `xml:"externalData,omitempty"`
	PrintSettings   *ml.Reserved            `xml:"printSettings,omitempty"`
	UserShapes      *ml.Reserved            `xml:"userShapes,omitempty"`
	ExtLst          *ml.Reserved            `xml:"extLst,omitempty"`
}

//Chart is a direct mapping of XSD CT_Chart
type Chart struct {
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

//Title is a direct mapping of XSD CT_Title
type Title = ml.Reserved

//PlotArea is a direct mapping of XSD CT_PlotArea
type PlotArea = ml.Reserved

//Legend is a direct mapping of XSD CT_Legend
type Legend = ml.Reserved
