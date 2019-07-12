// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
	"strconv"
)

//ShapeType is direct mapping of CT_ShapeType
type ShapeType struct {
	XMLName xml.Name `xml:"shapetype"`
	Path    string   `xml:"path,attr,omitempty"`
	shapeAttributes
	shapeElements
}

func (s *ShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//FIXME: Go1.12 doesnt support namespace prefixes, but 'spt' has namespace,
	//so better to manually encode it, than create a special type for it
	spt := s.Spt
	if s.Spt > 0 {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  ml.ApplyNamespacePrefix(ml.NamespaceVMLOffice, xml.Name{Local: "spt"}),
			Value: strconv.Itoa(int(s.Spt)),
		})
	}

	start.Name = ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)
	s.ReservedAttributes.ResolveNamespacePrefixes()
	s.ReservedElements.ResolveNamespacePrefixes()

	//keep original spt, but make encoder to omit empty original spt, cause we already manually encoded spt
	s.Spt = 0
	err := e.EncodeElement(*s, start)
	s.Spt = spt
	return err
}
