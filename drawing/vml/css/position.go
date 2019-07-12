// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package css

import "encoding/xml"

type Position byte

const (
	PositionStatic Position = iota
	PositionAbsolute
	PositionRelative
)

var (
	toPosition   map[string]Position
	fromPosition map[Position]string
)

func init() {
	fromPosition = map[Position]string{
		PositionStatic:   "static",
		PositionAbsolute: "absolute",
		PositionRelative: "relative",
	}

	toPosition = make(map[string]Position, len(fromPosition))
	for k, v := range fromPosition {
		toPosition[v] = k
	}
}

//String returns string presentation of Position
func (t Position) String() string {
	return fromPosition[t]
}

//MarshalXMLAttr marshal Position
func (t Position) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromPosition[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal Position
func (t *Position) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toPosition[attr.Value]; ok {
		*t = v
	}

	return nil
}
