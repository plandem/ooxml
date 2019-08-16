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
	SourceRect      *ml.Reserved `xml:"srcRect,omitempty"`
	Dpi             int          `xml:"dpi,attr,omitempty"`
	RotateWithShape bool         `xml:"rotWithShape,attr,omitempty"`
	*FillModeProperties
}

//Go1.12 has limited support of namespace prefixes, so use special type with hardcoded prefixes for marshalling
type blipFillProperties struct {
	Blip            *Blip        `xml:"a:blip,omitempty"`
	SourceRect      *ml.Reserved `xml:"a:srcRect,omitempty"`
	Dpi             int          `xml:"dpi,attr,omitempty"`
	RotateWithShape bool         `xml:"rotWithShape,attr,omitempty"`
	*FillModeProperties
}

func (t *BlipFillProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(blipFillProperties(*t), start)
}
