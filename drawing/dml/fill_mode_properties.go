// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//FillModeProperties is a direct mapping of EG_FillModeProperties
type FillModeProperties struct {
	Stretch         *ml.Reserved `xml:"stretch,omitempty"`
	Tile            *ml.Reserved `xml:"tile,omitempty"`
}

//Go1.12 has limited support of namespace prefixes, so use special type with hardcoded prefixes for marshalling
type fillModeProperties struct {
	Stretch         *ml.Reserved `xml:"a:stretch,omitempty"`
	Tile            *ml.Reserved `xml:"a:tile,omitempty"`

}

func (t *FillModeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(fillModeProperties(*t), start)
}