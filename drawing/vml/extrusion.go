package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//ExtrusionType is direct mapping of ST_ExtrusionType
type ExtrusionType string //enum

//ExtrusionRender is direct mapping of ST_ExtrusionRender
type ExtrusionRender string //enum

//ExtrusionPlane is direct mapping of ST_ExtrusionPlane
type ExtrusionPlane string //enum

//ExtrusionColorMode is direct mapping of ST_ColorMode
type ExtrusionColorMode string //enum

//Extrusion is direct mapping of CT_Extrusion
type Extrusion struct {
	XMLName        xml.Name `xml:"extrusion,omitempty" namespace:"o"`
	On                 bool               `xml:"on,attr,omitempty"`
	Type               ExtrusionType      `xml:"type,attr,omitempty"`
	Render             ExtrusionRender    `xml:"render,attr,omitempty"`
	Plane              ExtrusionPlane     `xml:"plane,attr,omitempty"`
	ViewpointOrigin    string             `xml:"viewpointorigin,attr,omitempty"`
	Viewpoint          string             `xml:"viewpoint,attr,omitempty"`
	SkewAngle          float64            `xml:"skewangle,attr,omitempty"`
	SkewAmt            string             `xml:"skewamt,attr,omitempty"`
	ForeDepth          string             `xml:"foredepth,attr,omitempty"`
	BackDepth          string             `xml:"backdepth,attr,omitempty"`
	Orientation        string             `xml:"orientation,attr,omitempty"`
	OrientationAngle   float64            `xml:"orientationangle,attr,omitempty"`
	LockRotationCenter *bool              `xml:"lockrotationcenter,attr,omitempty"`
	AutoRotationCenter bool               `xml:"autorotationcenter,attr,omitempty"`
	RotationCenter     string             `xml:"rotationcenter,attr,omitempty"`
	RotationAngle      string             `xml:"rotationangle,attr,omitempty"`
	ColorMode          ExtrusionColorMode `xml:"colormode,attr,omitempty"`
	Color              string             `xml:"color,attr,omitempty"`
	Shininess          float64            `xml:"shininess,attr,omitempty"`
	Specularity        float64            `xml:"specularity,attr,omitempty"`
	Diffusity          float64            `xml:"diffusity,attr,omitempty"`
	Metal              bool               `xml:"metal,attr,omitempty"`
	Edge               float64            `xml:"edge,attr,omitempty"`
	Facet              float64            `xml:"facet,attr,omitempty"`
	LightFace          *bool              `xml:"lightface,attr,omitempty"`
	LightHarsh     *bool    `xml:"lightharsh,attr,omitempty"`
	LightHarsh2    *bool    `xml:"lightharsh2,attr,omitempty"`
	LightLevel     float64  `xml:"lightlevel,attr,omitempty"`
	LightLevel2    float64  `xml:"lightlevel2,attr,omitempty"`
	LightPosition  string   `xml:"lightposition,attr,omitempty"`
	LightPosition2 string   `xml:"lightposition2,attr,omitempty"`
	Brightness     float64  `xml:"brightness,attr,omitempty"`
	Ext            Ext      `xml:"ext,attr,omitempty" namespace:"v"`
}

func (s *Extrusion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
