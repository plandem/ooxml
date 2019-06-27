package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//Formulas is direct mapping of CT_Formulas
type Formulas struct {
	XMLName xml.Name   `xml:"formulas" namespace:"v"`
	List    []*Formula `xml:"f" namespace:"v"`
}

//Formula is direct mapping of CT_F
type Formula struct {
	Eqn string `xml:"eqn,attr"`
}

func (s *Formulas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(s.List) > 0 {
		return internal2.Encode(s, e)
	}

	return nil
}
