package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
)

//ConnectType is direct mapping of ST_ConnectType
type ConnectType byte

var (
	toConnectType   map[string]ConnectType
	fromConnectType map[ConnectType]string
)

//List of all possible values for ConnectType
const (
	ConnectTypeNone ConnectType = iota
	ConnectTypeRect
	ConnectTypeSegments
	ConnectTypeCustom
)

func init() {
	fromConnectType = map[ConnectType]string{
		ConnectTypeNone:     "none",
		ConnectTypeRect:     "rect",
		ConnectTypeSegments: "segments",
		ConnectTypeCustom:   "custom",
	}

	toConnectType = make(map[string]ConnectType, len(fromConnectType))
	for k, v := range fromConnectType {
		toConnectType[v] = k
	}
}

//String returns string presentation of ConnectType
func (t ConnectType) String() string {
	return fromConnectType[t]
}

//MarshalXMLAttr marshal ConnectType
func (t ConnectType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: ooxml.ApplyNamespacePrefix(NamespaceOfficePrefix, name)}

	if v, ok := fromConnectType[t]; ok && t != ConnectTypeNone {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ConnectType
func (t *ConnectType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toConnectType[attr.Value]; ok {
		*t = v
	}

	return nil
}
