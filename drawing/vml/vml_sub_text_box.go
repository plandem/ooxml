// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//TextBox is direct mapping of CT_TextBox
type TextBox struct {
	XMLName   xml.Name      `xml:"textbox"`
	Style     string        `xml:"style,attr,omitempty"`
	Inset     string        `xml:"inset,attr,omitempty"`
	InsetMode InsetModeType `xml:"insetmode,attr,omitempty"`
	Text      string        `xml:",innerxml"`
	ml.ReservedAttributes
}

func (s *TextBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
