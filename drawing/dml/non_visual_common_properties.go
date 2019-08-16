// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import "github.com/plandem/ooxml/ml"

//NonVisualCommonProperties is a direct mapping of XSD CT_NonVisualDrawingProps
type NonVisualCommonProperties struct {
	ID          int    `xml:"id,attr"`
	Name        string `xml:"name,attr"`
	Description string `xml:"descr,attr,omitempty"`
	Title       string `xml:"title,attr,omitempty"`
	Hidden      bool   `xml:"hidden,attr,omitempty"`
	ml.ReservedElements
}
