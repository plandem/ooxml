// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Ref is a direct mapping of XSD CT_RelId
type Ref struct {
	RID ml.RID `xml:"id,attr"`
}

func (n *Ref) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = ml.ApplyNamespacePrefix(ml.NamespaceDMLChart, start.Name)
	start.Attr = append(start.Attr, ml.Namespaces(
		ml.NamespaceDMLChart,
		ml.NamespaceRelationships,
	)...)

	return e.EncodeElement(*n, start)
}
