// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Legend is a direct mapping of XSD CT_Legend
type Legend struct {
	Position ml.Property      `xml:"legendPos,omitempty"`
	Overlay  *ml.PropertyBool `xml:"overlay,omitempty"`
	ml.ReservedElements
}

func (n *Legend) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.Encode(*n)
}
