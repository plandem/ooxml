// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//StrokeJoinStyle is direct mapping of ST_StrokeJoinStyle
type StrokeJoinStyle byte

var (
	toStrokeJoinStyle   map[string]StrokeJoinStyle
	fromStrokeJoinStyle map[StrokeJoinStyle]string
)

//List of all possible values for StrokeJoinStyle
const (
	_ StrokeJoinStyle = iota
	StrokeJoinStyleRound
	StrokeJoinStyleBevel
	StrokeJoinStyleMiter
)

func init() {
	fromStrokeJoinStyle = map[StrokeJoinStyle]string{
		StrokeJoinStyleRound: "round",
		StrokeJoinStyleBevel: "bevel",
		StrokeJoinStyleMiter: "miter",
	}

	toStrokeJoinStyle = make(map[string]StrokeJoinStyle, len(fromStrokeJoinStyle))
	for k, v := range fromStrokeJoinStyle {
		toStrokeJoinStyle[v] = k
	}
}

//String returns string presentation of StrokeJoinStyle
func (t StrokeJoinStyle) String() string {
	return fromStrokeJoinStyle[t]
}

//MarshalXMLAttr marshal StrokeJoinStyle
func (t StrokeJoinStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromStrokeJoinStyle[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal StrokeJoinStyle
func (t *StrokeJoinStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toStrokeJoinStyle[attr.Value]; ok {
		*t = v
	}

	return nil
}
