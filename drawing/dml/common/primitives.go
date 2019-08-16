package common

import "github.com/plandem/ooxml/ml"

//Coordinate is a direct mapping of XSD ST_Coordinate
type Coordinate string //ST_CoordinateUnqualified s:ST_UniversalMeasure

//Point2D is a direct mapping of XSD CT_Point2D
type Point2D struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}

//PositiveSize2D is a direct mapping of XSD CT_PositiveSize2D
type PositiveSize2D struct {
	Cx uint `xml:"cx,attr"`
	Cy uint `xml:"cy,attr"`
}

//Transform2D is a direct mapping of XSD CT_Transform2D
type Transform2D struct {
	Offset         *Point2D        `xml:"off,omitempty"`
	Extent         *PositiveSize2D `xml:"ext,omitempty"`
	FlipHorizontal bool            `xml:"flipH,attr,omitempty"`
	FlipVertical   bool            `xml:"flipV,attr,omitempty"`
	Rotation       ml.PropertyInt  `xml:"rot,attr,omitempty"`
}

//func (n *Point2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
//	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
//}
//
//func (n *PositiveSize2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
//	return e.EncodeElement(*n, xml.StartElement{Name: ml.ApplyNamespacePrefix(ml.NamespaceDrawingExcel, start.Name)})
//}
