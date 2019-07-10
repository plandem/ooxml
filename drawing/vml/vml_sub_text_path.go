// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//TextPath is direct mapping of CT_TextPath
type TextPath struct {
	XMLName  xml.Name        `xml:"textpath"`
	Style    string          `xml:"style,attr,omitempty"`
	Text     string          `xml:"string,attr,omitempty"`
	On       ml.TriStateType `xml:"on,attr,omitempty"`
	FitShape ml.TriStateType `xml:"fitshape,attr,omitempty"`
	FitPath  ml.TriStateType `xml:"fitpath,attr,omitempty"`
	Trim     ml.TriStateType `xml:"trim,attr,omitempty"`
	XScale   ml.TriStateType `xml:"xscale,attr,omitempty"`
	ml.ReservedAttributes
}

func (s *TextPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
