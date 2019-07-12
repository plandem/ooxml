// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Handles is direct mapping of CT_Handles
type Handles struct {
	XMLName xml.Name `xml:"handles"`
	List    []Handle `xml:"h"`
}

//Handle is direct mapping of CT_H
type Handle struct {
	InvX        bool   `xml:"invx,attr,omitempty"`
	InvY        bool   `xml:"invy,attr,omitempty"`
	Switch      bool   `xml:"switch,attr,omitempty"`
	Position    string `xml:"position,attr,omitempty"`
	Polar       string `xml:"polar,attr,omitempty"`
	Map         string `xml:"map,attr,omitempty"`
	XRange      string `xml:"xrange,attr,omitempty"`
	YRange      string `xml:"yrange,attr,omitempty"`
	RadiusRange string `xml:"radiusrange,attr,omitempty"`
	ml.ReservedAttributes
}

func (s *Handles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(s.List) > 0 {
		return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
	}

	return nil
}

func (s *Handle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
