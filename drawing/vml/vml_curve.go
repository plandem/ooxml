package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//curve is direct mapping of CT_Curve
type curve struct {
	XMLName  xml.Name `xml:"curve"`
	From     string   `xml:"from,attr,omitempty"`
	To       string   `xml:"to,attr,omitempty"`
	Control1 string   `xml:"control1,attr,omitempty"`
	Control2 string   `xml:"control2,attr,omitempty"`
	coreAttributes
	shapeAttributes
	shapeElements
}

//Curve creates a new object with default values
func Curve() *curve {
	return &curve{
		From:     "0,0",
		To:       "30,20",
		Control1: "10,10",
		Control2: "20,0",
	}
}

func (s *curve) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
