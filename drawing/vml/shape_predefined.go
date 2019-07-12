// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

//List of VML predefined shapes
type predefinedShapes struct {
	Arc       []*arc       `xml:"arc,omitempty"`
	Curve     []*curve     `xml:"curve,omitempty"`
	Image     []*image     `xml:"image,omitempty"`
	Line      []*line      `xml:"line,omitempty"`
	Oval      []*oval      `xml:"oval,omitempty"`
	PolyLine  []*polyLine  `xml:"polyline,omitempty"`
	Rect      []*rect      `xml:"rect,omitempty"`
	RoundRect []*roundRect `xml:"roundrect,omitempty"`
}
