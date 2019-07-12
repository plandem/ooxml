// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//StrokeEndCap is direct mapping of ST_StrokeEndCap
type StrokeEndCap byte //enum

var (
	toStrokeEndCap   map[string]StrokeEndCap
	fromStrokeEndCap map[StrokeEndCap]string
)

//List of all possible values for StrokeEndCap
const (
	_ StrokeEndCap = iota
	StrokeEndCapFlat
	StrokeEndCapSquare
	StrokeEndCapRound
)

func init() {
	fromStrokeEndCap = map[StrokeEndCap]string{
		StrokeEndCapFlat:   "flat",
		StrokeEndCapSquare: "square",
		StrokeEndCapRound:  "round",
	}

	toStrokeEndCap = make(map[string]StrokeEndCap, len(fromStrokeEndCap))
	for k, v := range fromStrokeEndCap {
		toStrokeEndCap[v] = k
	}
}

//String returns string presentation of StrokeEndCap
func (t StrokeEndCap) String() string {
	return fromStrokeEndCap[t]
}

//MarshalXMLAttr marshal StrokeEndCap
func (t StrokeEndCap) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromStrokeEndCap[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal StrokeEndCap
func (t *StrokeEndCap) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toStrokeEndCap[attr.Value]; ok {
		*t = v
	}

	return nil
}
