// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//polyLine is direct mapping of CT_PolyLine
type polyLine struct {
	XMLName xml.Name `xml:"polyline"`
	Points  string   `xml:"points,attr,omitempty"`
	shapeAttributes
	shapeElements
}

//PolyLine creates a new object with default values
func PolyLine() *polyLine {
	return &polyLine{
		Points: "0,0 10,10",
	}
}

func (s *polyLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	s.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
