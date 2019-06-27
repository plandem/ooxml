package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//Handles is direct mapping of CT_Handles
type Handles struct {
	XMLName xml.Name  `xml:"handles" namespace:"v"`
	List    []*Handle `xml:"h" namespace:"v"`
}

//Handle is direct mapping of CT_H
type Handle struct {
	InvX        bool   `xml:"invx,attr,omitempty"`
	InvY        bool   `xml:"invy,attr,omitempty"`
	Switch      bool   `xml:"switch,attr,omitempty"`
	Position    string `xml:"position,attr,omitempty"`
	Polar       string `xml:"polar,attr,omitempty"`
	Map         string `xml:"map,attr,omitempty"`
	XRange      string `xml:"xrange,attr,omitempty"`
	YRange      string `xml:"yrange,attr,omitempty"`
	RadiusRange string `xml:"radiusrange,attr,omitempty"`
}

func (s *Handles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(s.List) > 0 {
		return internal2.Encode(s, e)
	}

	return nil
}
