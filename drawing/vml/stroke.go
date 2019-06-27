package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//StrokeDashStyle
type StrokeDashStyle string //enum

//StrokeArrowType is direct mapping of ST_StrokeArrowType
type StrokeArrowType string //enum

//StrokeArrowWidth is direct mapping of ST_StrokeArrowWidth
type StrokeArrowWidth string //enum

//StrokeArrowLength is direct mapping of ST_StrokeArrowLength
type StrokeArrowLength string //enum

//StrokeJoinStyle is direct mapping of ST_StrokeJoinStyle
type StrokeJoinStyle string //enum

//StrokeEndCap is direct mapping of ST_StrokeEndCap
type StrokeEndCap string //enum

//StrokeLineStyle is direct mapping of ST_StrokeLineStyle
type StrokeLineStyle string //enum

//StrokeAttributes is direct mapping of AG_StrokeAttributes
type StrokeAttributes struct {
	AltHRef          string            `xml:"althref,attr,omitempty" namespace:"o"`
	Color            string            `xml:"color,attr,omitempty"`
	Color2           string            `xml:"color2,attr,omitempty"`
	DashStyle        StrokeDashStyle   `xml:"dashstyle,attr,omitempty"`
	EndArrow         StrokeArrowType   `xml:"endarrow,attr,omitempty"`
	EndArrowLength   StrokeArrowLength `xml:"endarrowlength,attr,omitempty"`
	EndArrowWidth    StrokeArrowWidth  `xml:"endarrowwidth,attr,omitempty"`
	EndCap           StrokeEndCap      `xml:"endcap,attr,omitempty"`
	FillType         FillType          `xml:"filltype,attr,omitempty"`
	ForceDash        *bool             `xml:"forcedash,attr,omitempty" namespace:"o"`
	HRef             string            `xml:"href,attr,omitempty" namespace:"o"`
	ImageAlignShape  *bool             `xml:"imagealignshape,attr,omitempty"`
	ImageAspect      ImageAspect       `xml:"imageaspect,attr,omitempty"`
	ImageSize        string            `xml:"imagesize,attr,omitempty"`
	InsetPen         string            `xml:"insetpen,attr,omitempty"`
	JoinStyle        StrokeJoinStyle   `xml:"joinstyle,attr,omitempty"`
	LineStyle        StrokeLineStyle   `xml:"linestyle,attr,omitempty"`
	MiterLimit       *int              `xml:"miterlimit,attr,omitempty"`
	On               *bool             `xml:"on,attr,omitempty"`
	Opacity          string            `xml:"opacity,attr,omitempty"`
	RelID            string            `xml:"relid,attr,omitempty" namespace:"o"`
	Src              string            `xml:"src,attr,omitempty"`
	StartArrow       StrokeArrowType   `xml:"startarrow,attr,omitempty"`
	StartArrowLength StrokeArrowLength `xml:"startarrowlength,attr,omitempty"`
	StartArrowWidth  StrokeArrowWidth  `xml:"startarrowwidth,attr,omitempty"`
	Title            string            `xml:"title,attr,omitempty" namespace:"o"`
	Weight           string            `xml:"weight,attr,omitempty"`
	Relations
}

//Stroke is direct mapping of CT_Stroke
type Stroke struct {
	XMLName xml.Name          `xml:"stroke" namespace:"v"`
	Left    *StrokeAttributes `xml:"left" namespace:"o"`
	Top     *StrokeAttributes `xml:"top" namespace:"o"`
	Right   *StrokeAttributes `xml:"right" namespace:"o"`
	Bottom  *StrokeAttributes `xml:"bottom" namespace:"o"`
	Column  *StrokeAttributes `xml:"column" namespace:"o"`
	ID      string            `xml:"id,attr,omitempty"`
	StrokeAttributes
}

func (s *Stroke) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
