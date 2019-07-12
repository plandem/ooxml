// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Formulas is direct mapping of CT_Formulas
type Formulas struct {
	XMLName xml.Name  `xml:"formulas"`
	List    []Formula `xml:"f"`
}

//Formula is direct mapping of CT_F
type Formula string

type formula struct {
	Eqn string `xml:"eqn,attr"`
}

func (s *Formulas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(s.List) > 0 {
		return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
	}

	return nil
}

func (s Formula) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(formula{Eqn: string(s)}, xml.StartElement{
		Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name),
	})
}

func (s *Formula) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if len(start.Attr) > 0 {
		*s = Formula(start.Attr[0].Value)
	}

	return d.Skip()
}
