package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//InsetModeType is a type to encode ST_InsetMode
type InsetModeType byte

var (
	toInsetMode   map[string]InsetModeType
	fromInsetMode map[InsetModeType]string
)

//List of all possible values for InsetMode
const (
	_ InsetModeType = iota
	InsetModeCustom
	InsetModeAuto
)

func init() {
	fromInsetMode = map[InsetModeType]string{
		InsetModeAuto:   "auto",
		InsetModeCustom: "custom",
	}

	toInsetMode = make(map[string]InsetModeType, len(fromInsetMode))
	for k, v := range fromInsetMode {
		toInsetMode[v] = k
	}
}

//String returns string presentation of InsetModeType
func (t InsetModeType) String() string {
	return fromInsetMode[t]
}

//MarshalXMLAttr marshal InsetModeType
func (t InsetModeType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: ooxml.ApplyNamespacePrefix(NamespaceOfficePrefix, name)}

	if v, ok := fromInsetMode[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal InsetModeType
func (t *InsetModeType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toInsetMode[attr.Value]; ok {
		*t = v
	}

	return nil
}
