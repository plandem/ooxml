// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import "github.com/plandem/ooxml/ml"

//Blip is a direct mapping of XSD CT_BlipFillProperties
type Blip struct {
	RIDName ml.RIDName `xml:",attr"`
	Embed   ml.RID     `xml:"embed,attr,omitempty"`
	Link    ml.RID     `xml:"link,attr,omitempty"`
	ml.ReservedAttributes
	ml.ReservedElements
}
