// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//IdMap is direct mapping of CT_IdMap
type IdMap struct {
	XMLName xml.Name `xml:"idmap"`
	Ext     ExtType  `xml:"ext,attr,omitempty"`
	Data    string   `xml:"data,attr,omitempty"`
}

func (s *IdMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVMLOffice, start.Name)})
}
