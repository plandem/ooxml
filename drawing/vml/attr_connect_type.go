// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//ConnectType is direct mapping of ST_ConnectType
type ConnectType byte

var (
	toConnectType   map[string]ConnectType
	fromConnectType map[ConnectType]string
)

//List of all possible values for ConnectType
const (
	_ ConnectType = iota
	ConnectTypeNone
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
	attr := xml.Attr{Name: ml.ApplyNamespacePrefix(ml.NamespaceVMLOffice, name)}

	if v, ok := fromConnectType[t]; ok {
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
