// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

type Charts struct {
	Area2D    *ml.Reserved `xml:"areaChart,omitempty"`
	Area3D    *ml.Reserved `xml:"area3DChart,omitempty"`
	Line2D    *ml.Reserved `xml:"lineChart,omitempty"`
	Line3D    *ml.Reserved `xml:"line3DChart,omitempty"`
	Stock     *ml.Reserved `xml:"stockChart,omitempty"`
	Radar     *ml.Reserved `xml:"radarChart,omitempty"`
	Scatter   *ml.Reserved `xml:"scatterChart,omitempty"`
	Pie2D     *ml.Reserved `xml:"pieChart,omitempty"`
	Pie3D     *ml.Reserved `xml:"pie3DChart,omitempty"`
	Doughnut  *ml.Reserved `xml:"doughnutChart,omitempty"`
	Bar2D     *ml.Reserved `xml:"barChart,omitempty"`
	Bar3D     *ml.Reserved `xml:"bar3DChart,omitempty"`
	OffPie    *ml.Reserved `xml:"ofPieChart,omitempty"`
	Surface2D *ml.Reserved `xml:"surfaceChart,omitempty"`
	Surface3D *ml.Reserved `xml:"surface3DChart,omitempty"`
	Bubble    *ml.Reserved `xml:"bubbleChart,omitempty"`
}

type Axes struct {
	Date     *ml.Reserved `xml:"dateAx,omitempty"`
	Value    *ml.Reserved `xml:"valAx,omitempty"`
	Serial   *ml.Reserved `xml:"serAx,omitempty"`
	Category *ml.Reserved `xml:"catAx,omitempty"`
}

//PlotArea is a direct mapping of XSD CT_PlotArea
type PlotArea struct {
	Axes
	Charts
	ml.ReservedElements
}

func (n *PlotArea) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.Encode(*n)
}
