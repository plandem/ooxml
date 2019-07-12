// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package css

import (
	"encoding/xml"
)

type Visibility byte

const (
	VisibilityInherit Visibility = iota
	VisibilityHidden
	VisibilityVisible
	VisibilityCollapse
)

var (
	toVisibility   map[string]Visibility
	fromVisibility map[Visibility]string
)

func init() {
	fromVisibility = map[Visibility]string{
		VisibilityInherit:  "inherit",
		VisibilityHidden:   "hidden",
		VisibilityVisible:  "visible",
		VisibilityCollapse: "collapse",
	}

	toVisibility = make(map[string]Visibility, len(fromVisibility))
	for k, v := range fromVisibility {
		toVisibility[v] = k
	}
}

//String returns string presentation of Visibility
func (t Visibility) String() string {
	return fromVisibility[t]
}

//MarshalXMLAttr marshal Visibility
func (t Visibility) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromVisibility[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal Visibility
func (t *Visibility) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toVisibility[attr.Value]; ok {
		*t = v
	}

	return nil
}
