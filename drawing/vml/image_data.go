package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//ImageAttributes is direct mapping of AG_ImageAttributes
type ImageAttributes struct {
	Src           string  `xml:"src,attr,omitempty"`
	CropLeft      float64 `xml:"cropleft,attr,omitempty"`
	CropTop       float64 `xml:"croptop,attr,omitempty"`
	CropRight     float64 `xml:"cropright,attr,omitempty"`
	CropBottom    float64 `xml:"cropbottom,attr,omitempty"`
	Gain          float64 `xml:"gain,attr,omitempty"`
	BlackLevel    float64 `xml:"blacklevel,attr,omitempty"`
	Gamma         float64 `xml:"gamma,attr,omitempty"`
	GrayScale     bool    `xml:"grayscale,attr,omitempty"`
	BlackAndWhite bool    `xml:"bilevel,attr,omitempty"`
}

//ImageData is direct mapping of CT_ImageData
type ImageData struct {
	XMLName          xml.Name `xml:"imagedata" namespace:"v"`
	ID               string   `xml:"id,attr,omitempty"`
	AltHRef          string   `xml:"althref,attr,omitempty" namespace:"o"`
	HRef             string   `xml:"href,attr,omitempty" namespace:"o"`
	Title            string   `xml:"title,attr,omitempty" namespace:"o"`
	DetectMouseClick *bool    `xml:"detectmouseclick,attr,omitempty" namespace:"o"`
	Movie            float64  `xml:"movie,attr,omitempty" namespace:"o"`
	RelID            string   `xml:"relid,attr,omitempty" namespace:"o"`
	ChromaKey        string   `xml:"chromakey,attr,omitempty"`
	EmbossColor      string   `xml:"embosscolor,attr,omitempty"`
	ReColorTarget    string   `xml:"recolortarget,attr,omitempty"`
	ImageAttributes
	Relations
}

func (s *ImageData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
