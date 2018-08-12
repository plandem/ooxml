package ml

import (
	"encoding/xml"
)

//TargetMode is a type to encode XSD type
type TargetMode byte

//RelationType is a type to encode XSD type
type RelationType string

//RID is a helper type for marshaling references for relationship - r:id
type RID string

//RIDName is a helper type for marshaling namespace for relationships
type RIDName string

//Relation is a direct mapping of XSD type
type Relation struct {
	ID         string       `xml:"Id,attr"`
	Target     string       `xml:",attr"`
	Type       RelationType `xml:",attr"`
	TargetMode TargetMode   `xml:",attr,omitempty"`
}

//Relationships is a direct mapping of XSD type
type Relationships struct {
	XMLName       Name       `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
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

func (r *TargetMode) UnmarshalXMLAttr(attr xml.Attr) error {
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

//BeforeMarshalXML mark Relationships as non valid in case if there is no any relations inside
func (r *Relationships) BeforeMarshalXML() interface{} {
	if len(r.Relationships) == 0 {
		return nil
	}

	return r
}
