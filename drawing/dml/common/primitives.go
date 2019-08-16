package common

import "github.com/plandem/ooxml/ml"

//Coordinate is a direct mapping of XSD ST_Coordinate
//Office will read either a length followed by a unit or EMUs with no unit present, but will write only EMUs when no units are present.
type Coordinate int

//Point2D is a direct mapping of XSD CT_Point2D
type Point2D struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}

//PositiveSize2D is a direct mapping of XSD CT_PositiveSize2D
type PositiveSize2D struct {
	Height uint `xml:"cx,attr"` //cx - height in EMU
	Width  uint `xml:"cy,attr"` //cy - width in EMU
}

//Transform2D is a direct mapping of XSD CT_Transform2D
type Transform2D struct {
	Offset         *Point2D        `xml:"off,omitempty"`
	Size           *PositiveSize2D `xml:"ext,omitempty"`
	FlipHorizontal bool            `xml:"flipH,attr,omitempty"`
	FlipVertical   bool            `xml:"flipV,attr,omitempty"`
	Rotation       ml.PropertyInt  `xml:"rot,attr,omitempty"`
}
