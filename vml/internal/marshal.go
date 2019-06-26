package internal

import (
	"encoding"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
)

var (
	attrType          = reflect.TypeOf(xml.Attr{})
	marshalerAttrType = reflect.TypeOf((*xml.MarshalerAttr)(nil)).Elem()
	textMarshalerType = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	stringerType      = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

func marshalSimple(typ reflect.Type, val reflect.Value) (string, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(val.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'g', -1, val.Type().Bits()), nil
	case reflect.String:
		return val.String(), nil
	case reflect.Bool:
		return strconv.FormatBool(val.Bool()), nil
	}

	return "", &xml.UnsupportedTypeError{Type: typ}
}

func marshalAttr(start *xml.StartElement, name xml.Name, val reflect.Value) error {
	//xml.MarshalerAttr
	if val.CanInterface() && val.Type().Implements(marshalerAttrType) {
		attr, err := val.Interface().(xml.MarshalerAttr).MarshalXMLAttr(name)
		if err != nil {
			return err
		}

		if attr.Name.Local != "" {
			start.Attr = append(start.Attr, attr)
		}

		return nil
	}

	if val.CanAddr() {
		pv := val.Addr()

		if pv.CanInterface() && pv.Type().Implements(marshalerAttrType) {
			attr, err := pv.Interface().(xml.MarshalerAttr).MarshalXMLAttr(name)

			if err != nil {
				return err
			}

			if attr.Name.Local != "" {
				start.Attr = append(start.Attr, attr)
			}

			return nil
		}
	}

	//encoding.TextMarshaler
	if val.CanInterface() && val.Type().Implements(textMarshalerType) {
		text, err := val.Interface().(encoding.TextMarshaler).MarshalText()

		if err != nil {
			return err

		}

		start.Attr = append(start.Attr, xml.Attr{Name: name, Value: string(text)})
		return nil
	}

	if val.CanAddr() {
		pv := val.Addr()

		if pv.CanInterface() && pv.Type().Implements(textMarshalerType) {
			text, err := pv.Interface().(encoding.TextMarshaler).MarshalText()

			if err != nil {
				return err
			}

			start.Attr = append(start.Attr, xml.Attr{Name: name, Value: string(text)})
			return nil
		}
	}

	//fmt.Stringer
	if val.CanInterface() && val.Type().Implements(stringerType) {
		text := val.Interface().(fmt.Stringer).String()
		start.Attr = append(start.Attr, xml.Attr{Name: name, Value: string(text)})
		return nil
	}

	if val.CanAddr() {
		pv := val.Addr()
		if pv.CanInterface() && pv.Type().Implements(stringerType) {
			text := pv.Interface().(fmt.Stringer).String()
			start.Attr = append(start.Attr, xml.Attr{Name: name, Value: string(text)})
			return nil
		}
	}

	// Dereference or skip nil pointer, interface values.
	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return nil
		}

		val = val.Elem()
	}

	if val.Type() == attrType {
		start.Attr = append(start.Attr, val.Interface().(xml.Attr))
		return nil
	}

	s, err := marshalSimple(val.Type(), val)
	if err != nil {
		return err
	}

	start.Attr = append(start.Attr, xml.Attr{Name: name, Value: s})
	return nil
}

func Encode(value interface{}, e *xml.Encoder) error {
	val := reflect.ValueOf(value)
	if !val.IsValid() {
		return nil
	}

	// Drill into interfaces and pointers.
	// This can turn into an infinite loop given a cyclic chain,
	// but it matches the Go 1 behavior.
	for val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	typ := val.Type()

	tinfo, err := getTypeInfo(typ)
	if err != nil {
		return err
	}

	//get start element
	var start xml.StartElement
	if tinfo.xmlname != nil {
		xmlname := tinfo.xmlname

		if xmlname.name != "" {
			start.Name.Space, start.Name.Local = xmlname.xmlns, xmlname.name
		} else if v, ok := xmlname.value(val).Interface().(xml.Name); ok && v.Local != "" {
			start.Name = v
		}
	}

	if start.Name.Local == "" {
		name := typ.Name()

		if name == "" {
			return &xml.UnsupportedTypeError{typ}
		}

		start.Name.Local = name
	}

	// marshal attributes
	for i := range tinfo.fields {
		finfo := &tinfo.fields[i]
		if finfo.flags&fAttr == 0 {
			continue
		}

		fv := finfo.value(val)

		if finfo.flags&fOmitEmpty != 0 && isEmptyValue(fv) {
			continue
		}

		if fv.Kind() == reflect.Interface && fv.IsNil() {
			continue
		}

		//encode attribute
		xmlName := xml.Name{Space: finfo.xmlns, Local: finfo.name}
		if err := marshalAttr(&start, xmlName, fv); err != nil {
			return err
		}
	}

	// open start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	//encode nested elements
	for i := range tinfo.fields {
		finfo := &tinfo.fields[i]

		if finfo.flags&fElement == 0 {
			continue
		}

		inner := xml.StartElement{
			Name: xml.Name{
				Local: finfo.name,
				Space: finfo.xmlns,
			},
		}

		if err := e.EncodeElement(finfo.value(val).Interface(), inner); err != nil {
			return err
		}
	}

	// close start element
	if start.Name.Local != "" {
		return e.EncodeToken(start.End())
	}

	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
