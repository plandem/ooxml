// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

type officeDrawing struct {
	ShapeLayout *ShapeLayout `xml:"shapelayout,omitempty"`
	ShapeType   []*ShapeType `xml:"shapetype,omitempty"`
	Shape       []*Shape     `xml:"shape,omitempty"`
	predefinedShapes
	ml.ReservedElements
}

//Excel is type for Excel VML Drawings
type Excel officeDrawing

//Word is type for Word VML Drawings
type Word officeDrawing

//PowerPoint is type for PowerPoint VML Drawings
type PowerPoint officeDrawing

//MarshalXML marshals Excel Drawings
func (o *Excel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "xml"}
	start.Attr = append(start.Attr, ml.Namespaces(
		ml.NamespaceVML,
		ml.NamespaceVMLOffice,
		ml.NamespaceVMLExcel,
		ml.NamespaceRelationships,
	)...)
	o.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*o, start)
}

//MarshalXML marshals Word Drawings
func (o *Word) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "xml"}
	start.Attr = append(start.Attr, ml.Namespaces(
		ml.NamespaceVML,
		ml.NamespaceVMLOffice,
		ml.NamespaceVMLWord,
		ml.NamespaceRelationships,
	)...)
	o.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*o, start)
}

//MarshalXML marshals PowerPoint Drawings
func (o *PowerPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "xml"}
	start.Attr = append(start.Attr, ml.Namespaces(
		ml.NamespaceVML,
		ml.NamespaceVMLOffice,
		ml.NamespaceVMLPowerPoint,
		ml.NamespaceRelationships,
	)...)
	o.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*o, start)
}
