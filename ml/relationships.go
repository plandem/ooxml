package ml

import (
	"encoding/xml"
)

type TargetMode byte
type RelationType string
type RID string
type RIDName string

type Relation struct {
	ID         string `xml:"Id,attr"`
	Target     string `xml:",attr"`
	Type       RelationType `xml:",attr"`
	TargetMode TargetMode `xml:",attr,omitempty"`
}

type Relationships struct {
	XMLName       Name   `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
	Relationships []Relation `xml:"Relationship"`
}

const (
	TargetModeInternal TargetMode = iota
	TargetModeExternal
)

func (r *TargetMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	switch *r {
	case TargetModeInternal:
		attr = xml.Attr{}
	case TargetModeExternal:
		attr.Value = "External"
	}

	return attr, nil
}

func (r *TargetMode) UnmarshalXMLAttr(attr xml.Attr) (error) {
	switch attr.Value {
	case "External":
		*r = TargetModeExternal
	case "":
		*r = TargetModeInternal
	}

	return nil
}

func (r *RID) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "r:id"}, Value: string(*r)}
	return attr, nil
}

func (r *RIDName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: NamespaceRelationships}
	return attr, nil
}
