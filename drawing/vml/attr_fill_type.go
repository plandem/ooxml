// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//FillType is direct mapping of ST_FillType
type FillType byte

var (
	toFillType   map[string]FillType
	fromFillType map[FillType]string
)

//List of all possible values for FillType
const (
	_ FillType = iota
	FillTypeSolid
	FillTypeGradient
	FillTypeGradientRadial
	FillTypeTile
	FillTypePattern
	FillTypeFrame
)

func init() {
	fromFillType = map[FillType]string{
		FillTypeSolid:          "solid",
		FillTypeGradient:       "gradient",
		FillTypeGradientRadial: "gradientRadial",
		FillTypeTile:           "tile",
		FillTypePattern:        "pattern",
		FillTypeFrame:          "frame",
	}

	toFillType = make(map[string]FillType, len(fromFillType))
	for k, v := range fromFillType {
		toFillType[v] = k
	}
}

//String returns string presentation of FillType
func (t FillType) String() string {
	return fromFillType[t]
}

//MarshalXMLAttr marshal FillType
func (t FillType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromFillType[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal FillType
func (t *FillType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toFillType[attr.Value]; ok {
		*t = v
	}

	return nil
}
