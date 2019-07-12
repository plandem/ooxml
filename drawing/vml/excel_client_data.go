// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//ClientData is direct mapping of CT_ClientData
type ClientData struct {
	XMLName       xml.Name             `xml:"ClientData"`
	Type          ObjectType           `xml:"ObjectType,attr"`
	MoveWithCells ml.TriStateBlankTrue `xml:"MoveWithCells"`
	SizeWithCells ml.TriStateBlankTrue `xml:"SizeWithCells"`
	AutoFill      ml.TriStateBlankTrue `xml:"AutoFill"`
	AutoLine      ml.TriStateBlankTrue `xml:"AutoLine"`
	AutoPict      ml.TriStateBlankTrue `xml:"AutoPict"`
	RowHidden     ml.TriStateBlankTrue `xml:"RowHidden"`
	ColumnHidden  ml.TriStateBlankTrue `xml:"ColHidden"`
	MultiLine     ml.TriStateBlankTrue `xml:"MultiLine"`
	Visible       ml.TriStateBlankTrue `xml:"Visible"`
	Row           int                  `xml:"Row"`
	Column        int                  `xml:"Column"`
	Anchor        string               `xml:"Anchor"`
	ml.ReservedElements
}

//Go1.12 doesn't support namespace prefixes, so clientDataEncode is copy of original ClientData type, but with hardcoded namespace to simplify process of encoding
type clientDataEncode struct {
	XMLName       xml.Name
	Type          ObjectType           `xml:"ObjectType,attr"`
	MoveWithCells ml.TriStateBlankTrue `xml:"x:MoveWithCells,omitempty"`
	SizeWithCells ml.TriStateBlankTrue `xml:"x:SizeWithCells,omitempty"`
	AutoFill      ml.TriStateBlankTrue `xml:"x:AutoFill,omitempty"`
	AutoLine      ml.TriStateBlankTrue `xml:"x:AutoLine,omitempty"`
	AutoPict      ml.TriStateBlankTrue `xml:"x:AutoPict,omitempty"`
	RowHidden     ml.TriStateBlankTrue `xml:"x:RowHidden,omitempty"`
	ColumnHidden  ml.TriStateBlankTrue `xml:"x:ColHidden,omitempty"`
	MultiLine     ml.TriStateBlankTrue `xml:"x:MultiLine,omitempty"`
	Visible       ml.TriStateBlankTrue `xml:"x:Visible,omitempty"`
	Row           int                  `xml:"x:Row"`
	Column        int                  `xml:"x:Column"`
	Anchor        string               `xml:"x:Anchor,omitempty"`
	ml.ReservedElements
}

func (s *ClientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start = xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVMLExcel, start.Name)}
	s.ReservedElements.ResolveNamespacePrefixes()
	namespacedClientData := clientDataEncode(*s)
	return e.EncodeElement(namespacedClientData, start)
}
