// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//ImageAspect is direct mapping of ST_ImageAspect
type ImageAspect byte

var (
	toImageAspect   map[string]ImageAspect
	fromImageAspect map[ImageAspect]string
)

//List of all possible values for ImageAspect
const (
	_ ImageAspect = iota
	ImageAspectIgnore
	ImageAspectAtMost
	ImageAspectAtLeast
)

func init() {
	fromImageAspect = map[ImageAspect]string{
		ImageAspectIgnore:  "ignore",
		ImageAspectAtMost:  "atMost",
		ImageAspectAtLeast: "atLeast",
	}

	toImageAspect = make(map[string]ImageAspect, len(fromImageAspect))
	for k, v := range fromImageAspect {
		toImageAspect[v] = k
	}
}

//String returns string presentation of ImageAspect
func (t ImageAspect) String() string {
	return fromImageAspect[t]
}

//MarshalXMLAttr marshal ImageAspect
func (t ImageAspect) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromImageAspect[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ImageAspect
func (t *ImageAspect) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toImageAspect[attr.Value]; ok {
		*t = v
	}

	return nil
}
