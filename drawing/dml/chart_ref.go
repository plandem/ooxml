// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//ChartRef is a direct mapping of XSD CT_RelId
type ChartRef struct {
	RID ml.RID `xml:"id,attr"`
}

func (n *ChartRef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = ml.ApplyNamespacePrefix(ml.NamespaceDMLChart, start.Name)
	start.Attr = append(start.Attr, ml.Namespaces(
		ml.NamespaceDMLChart,
		ml.NamespaceRelationships,
	)...)

	return e.EncodeElement(*n, start)
}
