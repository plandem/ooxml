// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//GraphicFrame is a direct mapping of XSD CT_GraphicalObject
type GraphicFrame struct {
	Data *GraphicFrameData `xml:"graphicData"`
}

func (n *GraphicFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDML, start.Name)})
}
