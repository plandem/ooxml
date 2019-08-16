// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//fillModeProperties is a direct mapping of EG_FillModeProperties
type fillModeProperties struct {
	Stretch         *ml.Reserved `xml:"stretch,omitempty"`
	Tile            *ml.Reserved `xml:"tile,omitempty"`
}

////Go1.12 has limited support of namespace prefixes, so use special type with hardcoded prefixes for marshalling
type _fillModeProperties struct {
	Stretch         *ml.Reserved `xml:"a:stretch,omitempty"`
	Tile            *ml.Reserved `xml:"a:tile,omitempty"`

}

func (t *fillModeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(_fillModeProperties(*t), start)
}