package vml

import (
	"encoding"
	"encoding/xml"
	"fmt"
	_ "golang.org/x/net/webdav"
	"io"
	"reflect"
	"strconv"
)

// VML is outdated and deprecated format with broken XML rules here and there. So the main purpose of that package:
// 1) unmarshal content into structures and provide access attributes, nested nodes
// 2) marshal these structures as is to keep unrelated or untouched information in original state
// 3) during marshaling respect replaced objects and marshal it as required
//
// Recap: unmarshal vml - only 'Reserved' structs, marshal vml - respecting elements/attributes that were replaced with custom types

//Name type used to encode VML namespace
type Name string

//OfficeName type used to encode Office namespace
type OfficeName string

//ExcelName type used to encode Excel namespace
type ExcelName string

//WordName type used to encode Word namespace
type WordName string

//Word10Name type used to encode Word10 namespace
type Word10Name string

//PowerPointName type used to encode PowerPoint namespace
type PowerPointName string

//Basic support of Office VML.
type Office struct {
	XMLName     xml.Name     `xml:"xml"`
	Name        Name         `xml:",attr"`
	OfficeName  OfficeName   `xml:",attr"`
	ShapeLayout *ShapeLayout `xml:"shapelayout,omitempty"`
	ShapeType   []*ShapeType `xml:"shapetype,omitempty"`
	Shape       []*Shape     `xml:"shape,omitempty"`
}

//Basic support for Excel specific VML
type Excel struct {
	Office
	ExcelName `xml:",attr"`
}

//Basic support for Word specific VML
type Word struct {
	Office
	WordName   `xml:",attr"`
	Word10Name `xml:",attr"`
}

//Basic support for PowerPoint specific VML
type PowerPoint struct {
	Office
	PowerPointName `xml:",attr"`
}

//Reserved is universal type that hold information as is with access to attributes and nested nodes. It's a much slower than ml.Reserved
type Reserved struct {
	Name     xml.Name
	Attrs    map[string]interface{}
	Nested   []interface{}
	InnerXML interface{}
}

//ShapeLayout is alias for CT_ShapeLayout
type ShapeLayout = Reserved

const (
	NamespaceVML        = "urn:schemas-microsoft-com:vml"
	NamespaceOffice     = "urn:schemas-microsoft-com:office:office"
	NamespaceExcel      = "urn:schemas-microsoft-com:office:excel"
	NamespaceWord       = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"
	NamespaceWord10     = "urn:schemas-microsoft-com:office:word"
	NamespacePowerPoint = "urn:schemas-microsoft-com:office:powerpoint"
)

// Next returns the next token, ignore comments, processing instructions and directives.
func next(d *xml.Decoder) (xml.Token, error) {
	for {
		t, err := d.Token()
		if err != nil {
			return t, err
		}
		switch t.(type) {
		case xml.Comment, xml.Directive, xml.ProcInst:
			continue
		default:
			return t, nil
		}
	}
}

//MarshalXMLAttr marshals VML namespace
func (r *Name) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: NamespaceVML}
	return attr, nil
}

//MarshalXMLAttr marshals OfficeName namespace
func (r *OfficeName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: NamespaceOffice}
	return attr, nil
}

//MarshalXMLAttr marshals ExcelName namespace
func (r *ExcelName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: NamespaceExcel}
	return attr, nil
}

//MarshalXMLAttr marshals WordName namespace
func (r *WordName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: NamespaceWord}
	return attr, nil
}

//MarshalXMLAttr marshals Word10Name namespace
func (r *Word10Name) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:w10"}, Value: NamespaceWord10}
	return attr, nil
}

//MarshalXMLAttr marshals PowerPoint namespace
func (r *PowerPoint) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:pvml"}, Value: NamespacePowerPoint}
	return attr, nil
}

func resolveName(a xml.Name) xml.Name {
	switch a.Space {
	case NamespaceVML:
		return xml.Name{Local: "v:" + a.Local}
	case NamespaceOffice:
		return xml.Name{Local: "o:" + a.Local}
	case NamespaceExcel:
		return xml.Name{Local: "x:" + a.Local}
	case NamespaceWord:
		return xml.Name{Local: "w:" + a.Local}
	case NamespaceWord10:
		return xml.Name{Local: "w10:" + a.Local}
	case NamespacePowerPoint:
		return xml.Name{Local: "pvml:" + a.Local}
	}

	if len(a.Space) > 0 {
		return xml.Name{Local: a.Space + ":" + a.Local}
	}

	return a
}

func toString(v interface{}) (string, error) {
	switch vv := v.(type) {
	case string:
		return vv, nil
	case float32:
		return strconv.FormatFloat(float64(vv), 'g', -1, 64), nil
	case float64:
		return strconv.FormatFloat(vv, 'g', -1, 64), nil
	case uint:
		return strconv.FormatUint(uint64(vv), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(vv), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(vv), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(vv), 10), nil
	case uint64:
		return strconv.FormatUint(vv, 10), nil
	case int:
		return strconv.FormatInt(int64(vv), 10), nil
	case int8:
		return strconv.FormatInt(int64(vv), 10), nil
	case int16:
		return strconv.FormatInt(int64(vv), 10), nil
	case int32:
		return strconv.FormatInt(int64(vv), 10), nil
	case int64:
		return strconv.FormatInt(vv, 10), nil
	case bool:
		return strconv.FormatBool(vv), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("can't convert value of type=%T", v)
	}
}

//try to encode pointer to element, because marshaller uses by pointer to value, not value
func marshalElement(elem interface{}, e *xml.Encoder) error {
	val := reflect.ValueOf(elem)

	// Drill into interfaces and pointers.
	// This can turn into an infinite loop given a cyclic chain,
	// but it matches the Go 1 behavior.
	for val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	//convert into pointer to value
	pv := reflect.New(val.Type())
	pv.Elem().Set(val)

	return e.Encode(pv.Interface())
}

//MarshalXML marshals Reserved
func (r *Reserved) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = r.Name.Local
	start.Attr = make([]xml.Attr, 0, len(r.Attrs))

	var err error

	//normalize attributes
	for k, v := range r.Attrs {
		attr := xml.Attr{Name: xml.Name{Local: k}}

		switch vt := v.(type) {
		case xml.MarshalerAttr:
			if attr, err = vt.MarshalXMLAttr(xml.Name{Local: k}); err != nil {
				return err
			}
		case encoding.TextMarshaler:
			if text, err := vt.MarshalText(); err != nil {
				return err
			} else {
				attr.Value = string(text)
			}
		case fmt.Stringer:
			attr.Value = vt.String()
		default:
			if value, err := toString(v); err != nil {
				return err
			} else {
				attr.Value = value
			}
		}

		if attr.Name.Local != "" {
			start.Attr = append(start.Attr, attr)
		}
	}

	//encode start token with attributes
	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	//encode nested elements
	for _, nested := range r.Nested {
		if err := marshalElement(nested, e); err != nil {
			return err
		}
	}

	//encode inner xml
	if r.InnerXML != nil {
		var value []byte
		if marshaler, ok := r.InnerXML.(encoding.TextMarshaler); ok {
			if value, err = marshaler.MarshalText(); err != nil {
				return err
			}
		} else {
			if s, err := toString(r.InnerXML); err != nil {
				return err
			} else {
				value = []byte(s)
			}
		}

		if err := e.EncodeToken(xml.CharData([]byte(value))); err != nil {
			return err
		}
	}

	//encode end token
	return e.EncodeToken(start.End())
}

//UnmarshalXML unmarshal Reserved
func (r *Reserved) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	r.Attrs = make(map[string]interface{}, len(start.Attr))
	for i := range start.Attr {
		a := &start.Attr[i]
		a.Name = resolveName(a.Name)
		r.Attrs[a.Name.Local] = a.Value
	}

	r.Name = resolveName(start.Name)
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		switch nextToken := token.(type) {
		case xml.StartElement:
			var child Reserved
			if err := d.DecodeElement(&child, &nextToken); err != nil {
				return err
			}

			r.Nested = append(r.Nested, &child)
		case xml.CharData:
			r.InnerXML = string(nextToken)
		}
	}

	return nil
}
