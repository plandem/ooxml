// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//ExtType is a type to encode ST_Ext
type ExtType byte

var (
	toExtType   map[string]ExtType
	fromExtType map[ExtType]string
)

//List of all possible values for ExtType
const (
	_ ExtType = iota
	ExtTypeEdit
	ExtTypeView
	ExtTypeBackwardCompatible
)

func init() {
	fromExtType = map[ExtType]string{
		ExtTypeEdit:               "edit",
		ExtTypeView:               "view",
		ExtTypeBackwardCompatible: "backwardCompatible",
	}

	toExtType = make(map[string]ExtType, len(fromExtType))
	for k, v := range fromExtType {
		toExtType[v] = k
	}
}

//String returns string presentation of ExtType
func (t ExtType) String() string {
	return fromExtType[t]
}

//MarshalXMLAttr marshal ExtType
func (t ExtType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: ml.ApplyNamespacePrefix(ml.NamespaceVML, name)}

	if v, ok := fromExtType[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ExtType
func (t *ExtType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toExtType[attr.Value]; ok {
		*t = v
	}

	return nil
}
