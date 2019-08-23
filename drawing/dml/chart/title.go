// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Title is a direct mapping of XSD CT_Title
type Title struct {
	Overlay *ml.PropertyBool `xml:"overlay,omitempty"`
	Text    *Text            `xml:"tx,omitempty"`
	ml.ReservedElements
}

func (n *Title) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, start)
}
