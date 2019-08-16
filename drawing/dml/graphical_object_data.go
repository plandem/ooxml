// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml/chart"
	"github.com/plandem/ooxml/ml"
)

//GraphicalObjectData is a direct mapping of XSD CT_GraphicalObjectData
type GraphicalObjectData struct {
	Uri   string     `xml:"uri,attr"`
	Chart *chart.Ref `xml:"chart,omitempty"`
	ml.ReservedElements
}

func (n *GraphicalObjectData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()

	if n.Chart != nil {
		n.Uri = ml.NamespaceDrawingChart
	}

	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawing, start.Name)})
}
