// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//ShadowType is direct mapping of ST_ShadowType
type ShadowType byte //enum

var (
	toShadowType   map[string]ShadowType
	fromShadowType map[ShadowType]string
)

//List of all possible values for ShadowType
const (
	_ ShadowType = iota
	ShadowTypeSingle
	ShadowTypeDouble
	ShadowTypeEmboss
	ShadowTypePerspective
)

func init() {
	fromShadowType = map[ShadowType]string{
		ShadowTypeSingle:      "single",
		ShadowTypeDouble:      "double",
		ShadowTypeEmboss:      "emboss",
		ShadowTypePerspective: "perspective",
	}

	toShadowType = make(map[string]ShadowType, len(fromShadowType))
	for k, v := range fromShadowType {
		toShadowType[v] = k
	}
}

//String returns string presentation of ShadowType
func (t ShadowType) String() string {
	return fromShadowType[t]
}

//MarshalXMLAttr marshal ShadowType
func (t ShadowType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromShadowType[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ShadowType
func (t *ShadowType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toShadowType[attr.Value]; ok {
		*t = v
	}

	return nil
}
