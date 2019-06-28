package vml

//shapeAttributes is direct mapping of AG_AllShapeAttributes
type shapeAttributes struct {
	ChromaKey    string `xml:"chromakey,attr,omitempty"`
	Filled       bool   `xml:"filled,attr,omitempty"`
	FillColor    string `xml:"fillcolor,attr,omitempty"`
	Opacity      string `xml:"opacity,attr,omitempty"` //E.g.: 50%
	Stroked      bool   `xml:"stroked,attr,omitempty"`
	StrokeColor  string `xml:"strokecolor,attr,omitempty"`
	StrokeWeight string `xml:"strokeweight,attr,omitempty"` //E.g.: 1pt
	InsetPen     bool   `xml:"insetpen,attr,omitempty"`
	shapeOfficeAttributes
}
