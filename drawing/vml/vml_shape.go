// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/index"
	"github.com/plandem/ooxml/ml"
	"strconv"
	"strings"
)

//Shape is direct mapping of CT_Shape
type Shape struct {
	XMLName xml.Name `xml:"shape"`
	Type    string   `xml:"type,attr,omitempty"`
	ShapeType
}

func (s *Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}

//Hash builds hash code for all required values of Shape to use as unique index
func (s *Shape) Hash() index.Code {
	shape := s
	if shape == nil {
		shape = &Shape{}
	}

	clientData := shape.ClientData
	if clientData == nil {
		clientData = &ClientData{}
	}

	return index.Hash(strings.Join([]string{
		shape.Type,
		strconv.FormatInt(int64(shape.Spt), 10),
		strconv.FormatInt(int64(clientData.Column), 10),
		strconv.FormatInt(int64(clientData.Row), 10),
	}, ":"))
}
