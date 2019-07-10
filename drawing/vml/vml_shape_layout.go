// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//ShapeLayout is alias for CT_ShapeLayout
type ShapeLayout struct {
	XMLName xml.Name `xml:"shapelayout"`
	Ext     ExtType  `xml:"ext,attr,omitempty"`
	IdMap   *IdMap   `xml:"idmap,omitempty"`
	ml.ReservedElements
}

func (s *ShapeLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVMLOffice, start.Name)})
}
