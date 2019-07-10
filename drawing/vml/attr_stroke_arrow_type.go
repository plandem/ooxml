// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//StrokeArrowType is direct mapping of ST_StrokeArrowType
type StrokeArrowType byte //enum

var (
	toStrokeArrowType   map[string]StrokeArrowType
	fromStrokeArrowType map[StrokeArrowType]string
)

//List of all possible values for StrokeArrowType
const (
	_ StrokeArrowType = iota
	StrokeArrowTypeNone
	StrokeArrowTypeBlock
	StrokeArrowTypeClassic
	StrokeArrowTypeOval
	StrokeArrowTypeDiamond
	StrokeArrowTypeOpen
)

func init() {
	fromStrokeArrowType = map[StrokeArrowType]string{
		StrokeArrowTypeNone:    "none",
		StrokeArrowTypeBlock:   "block",
		StrokeArrowTypeClassic: "classic",
		StrokeArrowTypeOval:    "oval",
		StrokeArrowTypeDiamond: "diamond",
		StrokeArrowTypeOpen:    "open",
	}

	toStrokeArrowType = make(map[string]StrokeArrowType, len(fromStrokeArrowType))
	for k, v := range fromStrokeArrowType {
		toStrokeArrowType[v] = k
	}
}

//String returns string presentation of StrokeArrowType
func (t StrokeArrowType) String() string {
	return fromStrokeArrowType[t]
}

//MarshalXMLAttr marshal StrokeArrowType
func (t StrokeArrowType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromStrokeArrowType[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal StrokeArrowType
func (t *StrokeArrowType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toStrokeArrowType[attr.Value]; ok {
		*t = v
	}

	return nil
}
