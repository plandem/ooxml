package vml

type predefinedShapes struct {
	Arc       []*arc       `xml:"arc,omitempty"`
	Curve     []*curve     `xml:"curve,omitempty"`
	Image     []*image     `xml:"image,omitempty"`
	Line      []*line      `xml:"line,omitempty"`
	Oval      []*oval      `xml:"oval,omitempty"`
	PolyLine  []*polyLine  `xml:"polyLine,omitempty"`
	Rect      []*rect      `xml:"rect,omitempty"`
	RoundRect []*roundRect `xml:"roundrect,omitempty"`
}
