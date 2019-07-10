// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
	"strings"
)

//StrokeDashStyle is mapping for stroke dash type
type StrokeDashStyle byte

var (
	toStrokeDashStyle   map[string]StrokeDashStyle
	fromStrokeDashStyle map[StrokeDashStyle]string
)

//List of all possible values for StrokeDashStyle
const (
	_ StrokeDashStyle = iota
	StrokeDashStyleSolid
	StrokeDashStyleShortDash
	StrokeDashStyleShortDot
	StrokeDashStyleShortDashDot
	StrokeDashStyleShortDashDotDot
	StrokeDashStyleDot
	StrokeDashStyleDash
	StrokeDashStyleLongDash
	StrokeDashStyleDashDot
	StrokeDashStyleLongDashDot
	StrokeDashStyleLongDashDotDot
)

func init() {
	fromStrokeDashStyle = map[StrokeDashStyle]string{
		StrokeDashStyleSolid:           "solid",
		StrokeDashStyleShortDash:       "shortdash",
		StrokeDashStyleShortDot:        "shortdot",
		StrokeDashStyleShortDashDot:    "shortdashdot",
		StrokeDashStyleShortDashDotDot: "shortdashdotdot",
		StrokeDashStyleDot:             "dot",
		StrokeDashStyleDash:            "dash",
		StrokeDashStyleLongDash:        "longdash",
		StrokeDashStyleDashDot:         "dashdot",
		StrokeDashStyleLongDashDot:     "longdashdot",
		StrokeDashStyleLongDashDotDot:  "longdashdotdot",
	}

	toStrokeDashStyle = make(map[string]StrokeDashStyle, len(fromStrokeDashStyle))
	for k, v := range fromStrokeDashStyle {
		toStrokeDashStyle[v] = k
	}
}

//String returns string presentation of StrokeDashStyle
func (t StrokeDashStyle) String() string {
	return fromStrokeDashStyle[t]
}

//MarshalXMLAttr marshal StrokeDashStyle
func (t StrokeDashStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromStrokeDashStyle[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal StrokeDashStyle
func (t *StrokeDashStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toStrokeDashStyle[strings.ToLower(attr.Value)]; ok {
		*t = v
	}

	return nil
}
