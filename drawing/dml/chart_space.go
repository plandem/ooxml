// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml/chart"
	"github.com/plandem/ooxml/ml"
)

//ChartSpace is a direct mapping of XSD CT_ChartSpace
type ChartSpace struct {
	XMLName         xml.Name         `xml:"http://schemas.openxmlformats.org/drawingml/2006/chart chartSpace"`
	DMLName         ml.DMLName       `xml:",attr"`
	Lang            ml.Property      `xml:"lang,omitempty"`
	Style           ml.PropertyInt   `xml:"style,omitempty"`
	Date1904        *ml.PropertyBool `xml:"date1904,omitempty"`       //default true
	RoundedCorners  *ml.PropertyBool `xml:"roundedCorners,omitempty"` //default true
	ColorMapping    *ColorMapping    `xml:"clrMapOvr,omitempty"`
	ShapeProperties *ShapeProperties `xml:"spPr,omitempty"`
	TextBody        *TextBody        `xml:"txPr,omitempty"`
	Chart           *chart.Chart     `xml:"chart"`
	PivotSource     *ml.Reserved     `xml:"pivotSource,omitempty"`
	Protection      *ml.Reserved     `xml:"protection,omitempty"`
	ExternalData    *ml.Reserved     `xml:"externalData,omitempty"`
	PrintSettings   *ml.Reserved     `xml:"printSettings,omitempty"`
	UserShapes      *ml.Reserved     `xml:"userShapes,omitempty"`
	ExtLst          *ml.Reserved     `xml:"extLst,omitempty"`
}
