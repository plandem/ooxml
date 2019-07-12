// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml

import (
	"encoding/xml"
	"strconv"
	"unsafe"
)

//TriStateType is helper type to encode optional booleans 'true', 'false', 'default' to make it possible to omit 'default' or encode 'false' value if required
type TriStateType byte

//TriStateBlankTrue is subtype of TriStateType - during decoding process convert blank value to true value
type TriStateBlankTrue TriStateType

//TriStateBlankFalse is subtype of TriStateType - during decoding process convert blank value to false value
type TriStateBlankFalse TriStateType

//List of all possible values for TriStateType
const (
	TriStateBlank TriStateType = iota
	TriStateTrue
	TriStateFalse
)

//TriState is helper function that return TriState for boolean type
func TriState(state interface{}) TriStateType {
	if s, ok := state.(string); ok {
		if v, err := stringToTriStateType(s, TriStateFalse); err == nil {
			return v
		}
	} else {
		if b, ok := state.(bool); ok && b {
			return TriStateTrue
		}
	}

	return TriStateFalse
}

func stringToTriStateType(s string, blank TriStateType) (TriStateType, error) {
	if s != "" {
		if b, err := strconv.ParseBool(s); err != nil {
			return TriStateBlank, err
		} else {
			return TriState(b), nil
		}
	}

	return blank, nil
}

//String returns string presentation of TriStateType
func (t TriStateType) String() string {
	switch t {
	case TriStateTrue:
		return strconv.FormatBool(true)
	case TriStateFalse:
		return strconv.FormatBool(false)
	default:
		return ""
	}
}

func (t *TriStateType) defaultUnmarshalXMLAttr(blank TriStateType, attr xml.Attr) error {
	if ts, err := stringToTriStateType(attr.Value, blank); err != nil {
		return err
	} else {
		*t = ts
	}

	return nil
}

func (t *TriStateType) defaultUnmarshalXML(blank TriStateType, d *xml.Decoder, start xml.StartElement) error {
	s := ""
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if ts, err := stringToTriStateType(s, blank); err != nil {
		return err
	} else {
		*t = ts
	}

	return nil
}

//MarshalXMLAttr marshal TriStateType as attribute
func (t TriStateType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: t.String()}, nil
}

//UnmarshalXMLAttr unmarshal TriStateType as attribute
func (t *TriStateType) UnmarshalXMLAttr(attr xml.Attr) error {
	return t.defaultUnmarshalXMLAttr(TriStateBlank, attr)
}

//MarshalXML marshal TriStateType as element
func (t TriStateType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.String(), start)
}

//UnmarshalXML unmarshal TriStateType as element
func (t *TriStateType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return t.defaultUnmarshalXML(TriStateBlank, d, start)
}

//MarshalXMLAttr marshal TriStateBlankFalse as attribute
func (t TriStateBlankFalse) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return TriStateType(t).MarshalXMLAttr(name)
}

//MarshalXML marshal TriStateBlankFalse as element
func (t TriStateBlankFalse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return TriStateType(t).MarshalXML(e, start)
}

//UnmarshalXMLAttr unmarshal TriStateType as attribute
func (t *TriStateBlankFalse) UnmarshalXMLAttr(attr xml.Attr) error {
	tp := (*TriStateType)(unsafe.Pointer(t))
	return tp.defaultUnmarshalXMLAttr(TriStateFalse, attr)
}

//UnmarshalXML unmarshal TriStateType as element
func (t *TriStateBlankFalse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tp := (*TriStateType)(unsafe.Pointer(t))
	return tp.defaultUnmarshalXML(TriStateFalse, d, start)
}

//MarshalXMLAttr marshal TriStateBlankTrue as attribute
func (t TriStateBlankTrue) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return TriStateType(t).MarshalXMLAttr(name)
}

//MarshalXML marshal TriStateBlankTrue as element
func (t TriStateBlankTrue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return TriStateType(t).MarshalXML(e, start)
}

//UnmarshalXMLAttr unmarshal TriStateBlankTrue as attribute
func (t *TriStateBlankTrue) UnmarshalXMLAttr(attr xml.Attr) error {
	tp := (*TriStateType)(unsafe.Pointer(t))
	return tp.defaultUnmarshalXMLAttr(TriStateTrue, attr)
}

//UnmarshalXML unmarshal TriStateBlankTrue as element
func (t *TriStateBlankTrue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tp := (*TriStateType)(unsafe.Pointer(t))
	return tp.defaultUnmarshalXML(TriStateTrue, d, start)
}
