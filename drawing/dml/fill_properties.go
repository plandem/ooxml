// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"github.com/plandem/ooxml/ml"
)

//FillProperties is a direct mapping of XSD EG_FillProperties
type fillProperties struct {
	Blip     *BlipFillProperties `xml:"blipFill,omitempty"`
	No       *ml.Reserved        `xml:"noFill,omitempty"`
	Solid    *ml.Reserved        `xml:"solidFill,omitempty"`
	Gradient *ml.Reserved        `xml:"gradFill,omitempty"`
	Pattern  *ml.Reserved        `xml:"pattFill,omitempty"`
}
