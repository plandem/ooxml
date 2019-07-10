// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//FillMethod is direct mapping of ST_FillMethod
type FillMethod byte //enum

var (
	toFillMethod   map[string]FillMethod
	fromFillMethod map[FillMethod]string
)

//List of all possible values for FillType
const (
	_ FillMethod = iota
	FillMethodNone
	FillMethodLinear
	FillMethodSigma
	FillMethodAny
	FillMethodLinearSigma
)

func init() {
	fromFillMethod = map[FillMethod]string{
		FillMethodNone:        "none",
		FillMethodLinear:      "linear",
		FillMethodSigma:       "sigma",
		FillMethodAny:         "any",
		FillMethodLinearSigma: "linear sigma",
	}

	toFillMethod = make(map[string]FillMethod, len(fromFillMethod))
	for k, v := range fromFillMethod {
		toFillMethod[v] = k
	}
}

//String returns string presentation of FillMethod
func (t FillMethod) String() string {
	return fromFillMethod[t]
}

//MarshalXMLAttr marshal FillMethod
func (t FillMethod) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromFillMethod[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal FillMethod
func (t *FillMethod) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toFillMethod[attr.Value]; ok {
		*t = v
	}

	return nil
}
