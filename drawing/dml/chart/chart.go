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
	Title                     *Title           `xml:"title,omitempty"`
	PlotArea                  *PlotArea        `xml:"plotArea,omitempty"`
	Legend                    *Legend          `xml:"legend,omitempty"`
	DisplayBlanksAs           *ml.Property     `xml:"dispBlanksAs,omitempty"`
	AutoTitleIsDeleted        *ml.PropertyBool `xml:"autoTitleDeleted,omitempty"` //default true
	PlotVisibleOnly           *ml.PropertyBool `xml:"plotVisOnly,omitempty"`      //default true
	ShowDataLabelsOverMaximum *ml.PropertyBool `xml:"showDLblsOverMax,omitempty"` //default true
	ml.ReservedElements
}

func (n *Chart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.Encode(*n)
}
