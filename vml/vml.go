package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

type Name string
type OfficeName string
type ExcelName string
type WordName string
type PowerPointName string

//Basic support of Office VML
type Office struct {
	XMLName     xml.Name     `xml:"xml"`
	Name        Name         `xml:",attr"`
	OfficeName  OfficeName   `xml:",attr"`
	ShapeLayout *ShapeLayout `xml:"shapelayout,omitempty"`
	ShapeType   []*ShapeType `xml:"shapetype,omitempty"`
	Shape       []*Shape     `xml:"shape,omitempty"`
}

//Basic support for Excel VML
type Excel struct {
	Office
	ExcelName `xml:",attr"`
}

//Basic support for Word VML
type Word struct {
	Office
	WordName `xml:",attr"`
}

//Basic support for PowerPoint VML
type PowerPoint struct {
	Office
	PowerPointName `xml:",attr"`
}

//Reserved
type Reserved struct {
	ml.Reserved
	id string
	t  string
}

//CT_ShapeType
type ShapeType = Reserved

//CT_ShapeLayout
type ShapeLayout = Reserved

//CT_Shape
type Shape = Reserved

func (r *Name) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"}
	return attr, nil
}

func (r *OfficeName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"}
	return attr, nil
}

func (r *ExcelName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "urn:schemas-microsoft-com:office:excel"}
	return attr, nil
}

func (r *WordName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "urn:schemas-microsoft-com:office:word"}
	return attr, nil
}

func (r *PowerPoint) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "urn:schemas-microsoft-com:office:powerpoint"}
	return attr, nil
}

func resolveName(a xml.Name) xml.Name {
	switch a.Space {
	case "urn:schemas-microsoft-com:vml":
		return xml.Name{Local: "v:" + a.Local}
	case "urn:schemas-microsoft-com:office:office":
		return xml.Name{Local: "o:" + a.Local}
	case "urn:schemas-microsoft-com:office:excel":
		return xml.Name{Local: "x:" + a.Local}
	case "urn:schemas-microsoft-com:office:word":
		return xml.Name{Local: "w:" + a.Local}
	case "urn:schemas-microsoft-com:office:powerpoint":
		return xml.Name{Local: "p:" + a.Local}
	}

	return a
}

func (r *Reserved) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = r.Name
	start.Attr = r.Attrs
	return e.EncodeElement(r.InnerXML, start)
}

func (r *Reserved) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for i := range start.Attr {
		a := &start.Attr[i]
		a.Name = resolveName(a.Name)

		//cache some attributes to optimize access
		switch a.Name.Local {
		case "id":
			r.id = a.Value
		case "type":
			r.t = a.Value
		}
	}

	name := resolveName(start.Name)
	r.Name = name.Local
	r.Attrs = start.Attr
	return d.DecodeElement(&r.InnerXML, &start)
}

func (r *Reserved) Attr(name string) string {
	for _, a := range r.Attrs {
		if a.Name.Local == name {
			return a.Value
		}
	}

	return ""
}

func (r *Reserved) ID() string {
	return r.id
}

func (r *Reserved) Type() string {
	return r.t
}
