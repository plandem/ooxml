// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//StrokeArrowWidth is direct mapping of ST_StrokeArrowWidth
type StrokeArrowWidth byte

var (
	toStrokeArrowWidth   map[string]StrokeArrowWidth
	fromStrokeArrowWidth map[StrokeArrowWidth]string
)

//List of all possible values for StrokeArrowWidth
const (
	_ StrokeArrowWidth = iota
	StrokeArrowWidthNarrow
	StrokeArrowWidthMedium
	StrokeArrowWidthWide
)

func init() {
	fromStrokeArrowWidth = map[StrokeArrowWidth]string{
		StrokeArrowWidthNarrow: "narrow",
		StrokeArrowWidthMedium: "medium",
		StrokeArrowWidthWide:   "wide",
	}

	toStrokeArrowWidth = make(map[string]StrokeArrowWidth, len(fromStrokeArrowWidth))
	for k, v := range fromStrokeArrowWidth {
		toStrokeArrowWidth[v] = k
	}
}

//String returns string presentation of StrokeArrowWidth
func (t StrokeArrowWidth) String() string {
	return fromStrokeArrowWidth[t]
}

//MarshalXMLAttr marshal StrokeArrowWidth
func (t StrokeArrowWidth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromStrokeArrowWidth[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal StrokeArrowWidth
func (t *StrokeArrowWidth) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toStrokeArrowWidth[attr.Value]; ok {
		*t = v
	}

	return nil
}
