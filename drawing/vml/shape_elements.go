// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import "github.com/plandem/ooxml/ml"

//N.B.:
// Microsoft Office extended VML for more elements, but in reality only limited subset is used,
// so most used elements are exposed for better usability and rest elements considered as reserved to capture if required

type shapeElements struct {
	PathSettings *Path       `xml:"path,omitempty"`
	Formulas     *Formulas   `xml:"formulas,omitempty"`
	Handles      *Handles    `xml:"handles,omitempty"`
	Fill         *Fill       `xml:"fill,omitempty"`
	Stroke       *Stroke     `xml:"stroke,omitempty"`
	Shadow       *Shadow     `xml:"shadow,omitempty"`
	TextBox      *TextBox    `xml:"textbox,omitempty"`
	TextPath     *TextPath   `xml:"textpath,omitempty"`
	ImageData    *ImageData  `xml:"imagedata,omitempty"`
	ClientData   *ClientData `xml:"ClientData,omitempty"`

	ml.ReservedElements
}
