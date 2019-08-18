// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Group is a direct mapping of XSD CT_GroupShapeProperties
type Group struct {
	Transform *GroupTransform2D `xml:"xfrm,omitempty"`
	ml.ReservedElements
	fill
}

//Go1.12 has limited support of namespace prefixes, so use special type with hardcoded prefixes for marshalling
type group struct {
	Transform *GroupTransform2D `xml:"a:xfrm,omitempty"`
	ml.ReservedElements
	fill
}

func (n *Group) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(group(*n), start)
}
