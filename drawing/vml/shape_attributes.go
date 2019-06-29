package vml

import (
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
)

//N.B.:
// VML has tons of attributes, but in reality only limited subset is used by Microsoft Office,
// so most used attributes are exposed for better usability and rest attributes considered as reserved to capture if required

//FIXME:
//Go1.12:
// can't capture namespaced and common attribute same time (e.g.: id and r:id), so theoretically it's possible to have collision.
type shapeAttributes struct {
	InsetMode    InsetModeType `xml:"insetmode,attr,omitempty"`
	Spt          string        `xml:"spt,attr,omitempty" namespace:"o"`
	ID           string        `xml:"id,attr,omitempty"`
	Style        string        `xml:"style,attr,omitempty"`
	CoordSize    string        `xml:"coordsize,attr,omitempty"`
	Filled       bool          `xml:"filled,attr,omitempty"`
	FillColor    string        `xml:"fillcolor,attr,omitempty"`
	Opacity      css.Fraction  `xml:"opacity,attr,omitempty"`
	Stroked      bool          `xml:"stroked,attr,omitempty"`
	StrokeColor  string        `xml:"strokecolor,attr,omitempty"`
	StrokeWeight css.Number    `xml:"strokeweight,attr,omitempty"`
	InsetPen     bool          `xml:"insetpen,attr,omitempty"`

	ml.ReservedAttributes
}
