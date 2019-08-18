// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//GraphicFrameData is a direct mapping of XSD CT_GraphicalObjectData
type GraphicFrameData struct {
	Uri   string    `xml:"uri,attr"`
	Chart *ChartRef `xml:"chart,omitempty"`
	ml.ReservedElements
}

func (n *GraphicFrameData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()

	if n.Chart != nil {
		n.Uri = ml.NamespaceDMLChart
	}

	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDML, start.Name)})
}
