// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml/chart"
	"github.com/plandem/ooxml/ml"
)

//Graphic is a direct mapping of XSD CT_GraphicalObject
type Graphic struct {
	Data *GraphicData `xml:"graphicData"`
}

//GraphicData is a direct mapping of XSD CT_GraphicalObjectData
type GraphicData struct {
	Uri   string     `xml:"uri,attr"`
	Chart *chart.Ref `xml:"chart,omitempty"`
	ml.ReservedElements
}

func (n *Graphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawing, start.Name)})
}

func (n *GraphicData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()

	if n.Chart != nil {
		n.Uri = ml.NamespaceDrawingChart
	}

	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawing, start.Name)})
}
