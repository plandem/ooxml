// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//BlipCompression is a direct mapping of XSD ST_BlipCompression
type BlipCompression string //enum

//Blip is a direct mapping of XSD CT_Blip
type Blip struct {
	RIDName     ml.RIDName      `xml:",attr"`
	Embed       ml.RID          `xml:"embed,attr,omitempty"`
	Link        ml.RID          `xml:"link,attr,omitempty"`
	Compression BlipCompression `xml:"cstate,attr,omitempty"`
	ml.ReservedElements
}

//BlipFillProperties is a direct mapping of XSD CT_BlipFillProperties
type BlipFillProperties struct {
	Blip            *Blip        `xml:"blip,omitempty"`
	Dpi             int          `xml:"dpi,attr,omitempty"`
	RotateWithShape bool         `xml:"rotWithShape,attr,omitempty"`
	ml.ReservedElements
}

func (n *Blip) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDML, start.Name)})
}

func (n *BlipFillProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, start)
}
