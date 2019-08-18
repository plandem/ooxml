// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//NonVisualShapeProperties is a direct mapping of XSD CT_NonVisualDrawingShapeProps
type NonVisualShapeProperties struct {
	ml.ReservedAttributes
	ml.ReservedElements
}

func (n *NonVisualShapeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	n.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, start)
}
