package vml

import (
	"encoding/xml"
	internal2 "github.com/plandem/ooxml/drawing/vml/internal"
)

//SignatureLine is direct mapping of CT_SignatureLine
type SignatureLine struct {
	XMLName                xml.Name `xml:"signatureline,omitempty" namespace:"o"`
	ID                     string   `xml:"id,attr,omitempty"`
	ProvID                 string   `xml:"provid,attr,omitempty"`
	IsSignatureLine        *bool    `xml:"issignatureline,attr,omitempty"`
	SigningInstructionsSet *bool    `xml:"signinginstructionsset,attr,omitempty"`
	AllowComments          *bool    `xml:"allowcomments,attr,omitempty"`
	ShowSignDate           *bool    `xml:"showsigndate,attr,omitempty"`
	SuggestedSigner        string   `xml:"suggestedsigner,attr,omitempty"`
	SuggestedSigner2       string   `xml:"suggestedsigner2,attr,omitempty"`
	SuggestedSignerEmail   string   `xml:"suggestedsigneremail,attr,omitempty"`
	SigningInstructions    string   `xml:"signinginstructions,attr,omitempty"`
	AddLXml                string   `xml:"addlxml,attr,omitempty"`
	SigProvUrl             string   `xml:"sigprovurl,attr,omitempty"`
	Ext                    Ext      `xml:"ext,attr,omitempty" namespace:"v"`
}

func (s *SignatureLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal2.Encode(s, e)
}
