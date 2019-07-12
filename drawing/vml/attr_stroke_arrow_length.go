// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//StrokeArrowLength is direct mapping of ST_StrokeArrowLength
type StrokeArrowLength byte

var (
	toStrokeArrowLength   map[string]StrokeArrowLength
	fromStrokeArrowLength map[StrokeArrowLength]string
)

//List of all possible values for StrokeArrowLength
const (
	_ StrokeArrowLength = iota
	StrokeArrowLengthShort
	StrokeArrowLengthMedium
	StrokeArrowLengthLong
)

func init() {
	fromStrokeArrowLength = map[StrokeArrowLength]string{
		StrokeArrowLengthShort:  "short",
		StrokeArrowLengthMedium: "medium",
		StrokeArrowLengthLong:   "long",
	}

	toStrokeArrowLength = make(map[string]StrokeArrowLength, len(fromStrokeArrowLength))
	for k, v := range fromStrokeArrowLength {
		toStrokeArrowLength[v] = k
	}
}

//String returns string presentation of StrokeArrowLength
func (t StrokeArrowLength) String() string {
	return fromStrokeArrowLength[t]
}

//MarshalXMLAttr marshal StrokeArrowLength
func (t StrokeArrowLength) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromStrokeArrowLength[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal StrokeArrowLength
func (t *StrokeArrowLength) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toStrokeArrowLength[attr.Value]; ok {
		*t = v
	}

	return nil
}
