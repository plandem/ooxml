// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"encoding/xml"
)

//ObjectType is direct mapping of ST_ObjectType
type ObjectType byte

var (
	toObjectType   map[string]ObjectType
	fromObjectType map[ObjectType]string
)

//List of all possible values for ObjectType
const (
	_ ObjectType = iota
	ObjectTypeButton
	ObjectTypeCheckbox
	ObjectTypeDialog
	ObjectTypeDrop
	ObjectTypeEdit
	ObjectTypeGBox
	ObjectTypeLabel
	ObjectTypeLineA
	ObjectTypeList
	ObjectTypeMovie
	ObjectTypeNote
	ObjectTypePict
	ObjectTypeRadio
	ObjectTypeRectA
	ObjectTypeScroll
	ObjectTypeSpin
	ObjectTypeShape
	ObjectTypeGroup
	ObjectTypeRect
)

func init() {
	fromObjectType = map[ObjectType]string{
		ObjectTypeButton:   "Button",
		ObjectTypeCheckbox: "Checkbox",
		ObjectTypeDialog:   "Dialog",
		ObjectTypeDrop:     "Drop",
		ObjectTypeEdit:     "Edit",
		ObjectTypeGBox:     "GBox",
		ObjectTypeLabel:    "Label",
		ObjectTypeLineA:    "LineA",
		ObjectTypeList:     "List",
		ObjectTypeMovie:    "Movie",
		ObjectTypeNote:     "Note",
		ObjectTypePict:     "Pict",
		ObjectTypeRadio:    "Radio",
		ObjectTypeRectA:    "RectA",
		ObjectTypeScroll:   "Scroll",
		ObjectTypeSpin:     "Spin",
		ObjectTypeShape:    "Shape",
		ObjectTypeGroup:    "Group",
		ObjectTypeRect:     "Rect",
	}

	toObjectType = make(map[string]ObjectType, len(fromObjectType))
	for k, v := range fromObjectType {
		toObjectType[v] = k
	}
}

//String returns string presentation of ObjectType
func (t ObjectType) String() string {
	return fromObjectType[t]
}

//MarshalXMLAttr marshal ObjectType
func (t ObjectType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := fromObjectType[t]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal ObjectType
func (t *ObjectType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := toObjectType[attr.Value]; ok {
		*t = v
	}

	return nil
}
