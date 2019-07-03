package vml

import (
	"fmt"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
	"regexp"
	"strconv"
)

//N.B.:
// VML has tons of attributes, but in reality only limited subset is used by Microsoft Office,
// so most used attributes are exposed for better usability and rest attributes considered as reserved to capture if required

//FIXME:
//Go1.12:
// can't capture namespaced and common attribute same time (e.g.: id and r:id), so theoretically it's possible to have collision.
type shapeAttributes struct {
	InsetMode    InsetModeType   `xml:"insetmode,attr,omitempty"`
	resolvedID   int
	Spt          int             `xml:"spt,attr,omitempty"`
	Opacity      css.Fraction    `xml:"opacity,attr,omitempty"`
	StrokeWeight css.Number      `xml:"strokeweight,attr,omitempty"`
	Style        string          `xml:"style,attr,omitempty"`
	ID           string          `xml:"id,attr,omitempty"`
	CoordSize    string          `xml:"coordsize,attr,omitempty"`
	FillColor    string          `xml:"fillcolor,attr,omitempty"`
	StrokeColor  string          `xml:"strokecolor,attr,omitempty"`
	Filled       ml.TriStateType `xml:"filled,attr,omitempty"`
	Stroked      ml.TriStateType `xml:"stroked,attr,omitempty"`
	InsetPen     ml.TriStateType `xml:"insetpen,attr,omitempty"`

	ml.ReservedAttributes
}

var (
	regexpShapeID  = regexp.MustCompile(`_x0000_s([\d]+)`)
)

//ResolvedID() returns ID as integer
func (s shapeAttributes) ResolvedID() int {
	if s.resolvedID == 0 {
		if matched := regexpShapeID.FindSubmatch([]byte(s.ID)); len(matched) > 0 {
			if id, err := strconv.Atoi(string(matched[1])); err != nil {
				panic(fmt.Errorf("can't get ID of shape: %s", s.ID))
			} else {
				s.resolvedID = id
			}
		}
	}

	return s.resolvedID
}
