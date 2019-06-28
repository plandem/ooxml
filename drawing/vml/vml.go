package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

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

//Group is alias for CT_Group
type Group = Reserved

//Basic support of Office VML
type Office struct {
	XMLName     xml.Name     `xml:"xml"`
	Name        Name         `xml:",attr"`
	RIDName     ml.RIDName   `xml:",attr"`
	OfficeName  OfficeName   `xml:",attr"`
	ShapeLayout *ShapeLayout `xml:"shapelayout,omitempty"`
	ShapeType   []*ShapeType `xml:"shapetype,omitempty"`
	Shape       []*Shape     `xml:"shape,omitempty"`
	Group       []*Group     `xml:"group,omitempty"`
	Diagram     []*Diagram   `xml:"diagram,omitempty"`
	predefinedShapes
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

//Reserved is special type that catches all inner content AS IS to save original information - used to mark 'non implemented' elements. Supports namespace prefixes.
type Reserved ml.Reserved

const (
	NamespaceVML        = "urn:schemas-microsoft-com:vml"
	NamespaceOffice     = "urn:schemas-microsoft-com:office:office"
	NamespaceExcel      = "urn:schemas-microsoft-com:office:excel"
	NamespaceWord       = "urn:schemas-microsoft-com:office:word"
	NamespacePowerPoint = "urn:schemas-microsoft-com:office:powerpoint"

	NamespaceVMLPrefix        = "v"
	NamespaceOfficePrefix     = "o"
	NamespaceExcelPrefix      = "e"
	NamespaceWordPrefix       = "w"
	NamespacePowerPointPrefix = "p"
)

//resolveName tries to resolve namespace and apply prefix for it
func resolveName(a xml.Name) xml.Name {
	switch a.Space {
	case NamespaceVML:
		return xml.Name{Local: NamespaceVMLPrefix + ":" + a.Local}
	case NamespaceOffice:
		return xml.Name{Local: NamespaceOfficePrefix + ":" + a.Local}
	case NamespaceExcel:
		return xml.Name{Local: NamespaceExcelPrefix + ":" + a.Local}
	case NamespaceWord:
		return xml.Name{Local: NamespaceWordPrefix + ":" + a.Local}
	case NamespacePowerPoint:
		return xml.Name{Local: NamespacePowerPointPrefix + ":" + a.Local}
	}

	return a
}

//resolveAttributesName tries to resolve namespace and apply prefix for it for all attributes
func resolveAttributesName(attrs []xml.Attr) {
	for i, attr := range attrs {
		attrs[i].Name = resolveName(attr.Name)
	}
}

//MarshalXMLAttr marshals VML namespace
func (r *Name) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespaceVMLPrefix}, Value: NamespaceVML}
	return attr, nil
}

//MarshalXMLAttr marshals OfficeName namespace
func (r *OfficeName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespaceOfficePrefix}, Value: NamespaceOffice}
	return attr, nil
}

//MarshalXMLAttr marshals ExcelName namespace
func (r *ExcelName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespaceExcelPrefix}, Value: NamespaceExcel}
	return attr, nil
}

//MarshalXMLAttr marshals WordName namespace
func (r *WordName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespaceWordPrefix}, Value: NamespaceWord}
	return attr, nil
}

//MarshalXMLAttr marshals PowerPoint namespace
func (r *PowerPoint) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespacePowerPointPrefix}, Value: NamespacePowerPoint}
	return attr, nil
}

//MarshalXML marshal Reserved
func (r *Reserved) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	r.Name = resolveName(r.Name)
	r.Attrs = start.Attr
	resolveAttributesName(r.Attrs)

	mr := ml.Reserved(*r)
	return e.Encode(&mr)
}

//UnmarshalXML unmarshal Reserved
func (r *Reserved) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var mr ml.Reserved
	if err := d.DecodeElement(&mr, &start); err != nil {
		return err
	}

	*r = Reserved(mr)
	return nil
}
