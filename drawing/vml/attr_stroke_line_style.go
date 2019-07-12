// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//StrokeLineStyle is direct mapping of ST_StrokeLineStyle
type StrokeLineStyle byte

var (
	toStrokeLineStyle   map[string]StrokeLineStyle
	fromStrokeLineStyle map[StrokeLineStyle]string
)

//List of all possible values for StrokeLineStyle
const (
	_ StrokeLineStyle = iota
	StrokeLineStyleSingle
	StrokeLineStyleThinThin
	StrokeLineStyleThinThick
	StrokeLineStyleThickThin
	StrokeLineStyleThickBetweenThin
)

func init() {
	fromStrokeLineStyle = map[StrokeLineStyle]string{
		StrokeLineStyleSingle:           "single",
		StrokeLineStyleThinThin:         "thinThin",
		StrokeLineStyleThinThick:        "thinThick",
		StrokeLineStyleThickThin:        "thickThin",
		StrokeLineStyleThickBetweenThin: "thickBetweenThin",
	}

	toStrokeLineStyle = make(map[string]StrokeLineStyle, len(fromStrokeLineStyle))
	for k, v := range fromStrokeLineStyle {
		toStrokeLineStyle[v] = k
	}
}

//String returns string presentation of StrokeLineStyle
func (t StrokeLineStyle) String() string {
	return fromStrokeLineStyle[t]
}

//MarshalXMLAttr marshal StrokeLineStyle
func (t StrokeLineStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromStrokeLineStyle[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal StrokeLineStyle
func (t *StrokeLineStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toStrokeLineStyle[attr.Value]; ok {
		*t = v
	}

	return nil
}
