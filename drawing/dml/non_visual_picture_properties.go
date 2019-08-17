// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//NonVisualPictureProperties is a direct mapping of XSD CT_NonVisualPictureProperties
type NonVisualPictureProperties struct {
	Locking              *PictureLocking `xml:"picLocks,omitempty"`
	PreferRelativeResize *bool           `xml:"preferRelativeResize,attr,omitempty"`
	ml.ReservedElements
}

//PictureLocking is a direct mapping of XSD CT_PictureLocking
type PictureLocking struct {
	NoCrop bool `xml:"noCrop,attr,omitempty"`
	ml.ReservedElements
	locking
}

func (n *NonVisualPictureProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, start)
}

func (n *PictureLocking) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, start)
}