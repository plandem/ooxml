// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//imageAttributes is direct mapping of AG_ImageAttributes
type imageAttributes struct {
	Src           string          `xml:"src,attr,omitempty"`
	CropLeft      float64         `xml:"cropleft,attr,omitempty"`
	CropTop       float64         `xml:"croptop,attr,omitempty"`
	CropRight     float64         `xml:"cropright,attr,omitempty"`
	CropBottom    float64         `xml:"cropbottom,attr,omitempty"`
	Gain          float64         `xml:"gain,attr,omitempty"`
	BlackLevel    float64         `xml:"blacklevel,attr,omitempty"`
	Gamma         float64         `xml:"gamma,attr,omitempty"`
	GrayScale     ml.TriStateType `xml:"grayscale,attr,omitempty"`
	BlackAndWhite ml.TriStateType `xml:"bilevel,attr,omitempty"`
}

//ImageData is direct mapping of CT_ImageData
type ImageData struct {
	XMLName xml.Name `xml:"imagedata"`
	ml.ReservedAttributes
	imageAttributes
}

func (s *ImageData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s.ReservedAttributes.ResolveNamespacePrefixes()
	return e.EncodeElement(*s, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, start.Name)})
}
