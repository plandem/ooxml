// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

type DMLName string

//Space is a direct mapping of XSD CT_ChartSpace
type Space struct {
	XMLName        xml.Name          `xml:"http://schemas.openxmlformats.org/drawingml/2006/chart chartSpace"`
	DMLChart       DMLName           `xml:",attr"`
	DMLName        dml.Name          `xml:",attr"`
	RIDName        ml.RIDName        `xml:",attr"`
	Lang           ml.Property       `xml:"lang,omitempty"`
	Style          ml.PropertyInt    `xml:"style,omitempty"`
	Date1904       *ml.PropertyBool  `xml:"date1904,omitempty"`       //default true
	RoundedCorners *ml.PropertyBool  `xml:"roundedCorners,omitempty"` //default true
	ColorMapping   *dml.ColorMapping `xml:"clrMapOvr,omitempty"`
	Chart          *Chart            `xml:"chart"`
	Shape          *dml.Shape        `xml:"spPr,omitempty"`
	TextBody       *dml.TextBody     `xml:"txPr,omitempty"`
	ml.ReservedElements
}

func (n DMLName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if prefix, ok := ml.ResolveNamespacePrefix(ml.NamespaceDMLChart); ok {
		return xml.Attr{Name: xml.Name{Local: "xmlns:" + prefix}, Value: ml.NamespaceDMLChart}, nil
	}

	return xml.Attr{}, ml.ErrorNamespace(ml.NamespaceDMLChart)
}

func (n *Space) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.Encode(*n)
}
