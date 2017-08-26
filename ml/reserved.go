package ml

import "encoding/xml"

//Reserved is special type that catches all inner content AS IS to save original information - used to mark 'non implemented' elements
type Reserved struct {
	Name     string
	Attrs    []xml.Attr `xml:",attr"`
	InnerXML *struct {
		XML string `xml:",innerxml"`
	} `xml:",omitempty"`
}

func (r *Reserved) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = r.Name
	start.Attr = r.Attrs
	return e.EncodeElement(r.InnerXML, start)
}

func (r *Reserved) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	r.Name = start.Name.Local
	r.Attrs = start.Attr
	return d.DecodeElement(&r.InnerXML, &start)
}
