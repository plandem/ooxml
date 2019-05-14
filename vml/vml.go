package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

// VML is outdated and deprecated format with broken XML rules here and there. So to the main purpose of that package:
// 1) unmarshal content into structures and provide access to few core attributes to analyze it if required
// 2) marshal these structures as is to keep unrelated or untouched information in original state

//Name type used to encode VML namespace
type Name string

//OfficeName type used to encode Office namespace
type OfficeName string

//ExcelName type used to encode Excel namespace
type ExcelName string

//WordName type used to encode Word namespace
type WordName string

//PowerPointName type used to encode PowerPoint namespace
type PowerPointName string

//Basic support of Office VML.
type Office struct {
	XMLName     xml.Name     `xml:"xml"`
	Name        Name         `xml:",attr"`
	OfficeName  OfficeName   `xml:",attr"`
	ShapeLayout *ShapeLayout `xml:"shapelayout,omitempty"`
	ShapeType   []*ShapeType `xml:"shapetype,omitempty"`
	Shape       []*Shape     `xml:"shape,omitempty"`
}

//Basic support for Excel specific VML
type Excel struct {
	Office
	ExcelName `xml:",attr"`
}

//Basic support for Word specific VML
type Word struct {
	Office
	WordName `xml:",attr"`
}

//Basic support for PowerPoint specific VML
type PowerPoint struct {
	Office
	PowerPointName `xml:",attr"`
}

//Reserved is universal type that hold information as is with cached few core attributes for faster access
type Reserved struct {
	ml.Reserved
	id string
	t  string
}

//ShapeType is alias for CT_ShapeType
type ShapeType = Reserved

//ShapeLayout is alias for CT_ShapeLayout
type ShapeLayout = Reserved

//Shape is alias for CT_Shape
type Shape = Reserved

//MarshalXMLAttr marshals VML namespace
func (r *Name) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"}
	return attr, nil
}

//MarshalXMLAttr marshals OfficeName namespace
func (r *OfficeName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"}
	return attr, nil
}

//MarshalXMLAttr marshals ExcelName namespace
func (r *ExcelName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "urn:schemas-microsoft-com:office:excel"}
	return attr, nil
}

//MarshalXMLAttr marshals WordName namespace
func (r *WordName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "urn:schemas-microsoft-com:office:word"}
	return attr, nil
}

//MarshalXMLAttr marshals PowerPoint namespace
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

//MarshalXML marshals Reserved
func (r *Reserved) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = r.Name
	start.Attr = r.Attrs
	return e.EncodeElement(r.InnerXML, start)
}

//UnmarshalXML unmarshal Reserved
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

//Attr returns value for attribute with name, for namespaced attribute provide full name, e.g.: "o:insetmode"
func (r *Reserved) Attr(name string) string {
	for _, a := range r.Attrs {
		if a.Name.Local == name {
			return a.Value
		}
	}

	return ""
}

//ID returns cached value of id attribute
func (r *Reserved) ID() string {
	return r.id
}

//Type returns cached value of type attribute
func (r *Reserved) Type() string {
	return r.t
}
