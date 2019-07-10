// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Path is direct mapping of CT_Path
type Path struct {
	XMLName         xml.Name        `xml:"path"`
	Value           string          `xml:"v,attr,omitempty"`
	ConnectType     ConnectType     `xml:"connecttype,attr,omitempty"`
	FillOK          ml.TriStateType `xml:"fillok,attr,omitempty"`
	StrokeOK        ml.TriStateType `xml:"strokeok,attr,omitempty"`
	ShadowOK        ml.TriStateType `xml:"shadowok,attr,omitempty"`
	ArrowOK         ml.TriStateType `xml:"arrowok,attr,omitempty"`
	GradientShapeOK ml.TriStateType `xml:"gradientshapeok,attr,omitempty"`
	ml.ReservedAttributes
}

func (s *Path) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
