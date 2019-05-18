package vml

import (
	"encoding"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
)

// VML is outdated and deprecated format with broken XML rules here and there. So the main purpose of that package:
// 1) unmarshal content into structures and provide access attributes, nested nodes
// 2) marshal these structures as is to keep unrelated or untouched information in original state

//Name type used to encode VML namespace
type Name string

//OfficeName type used to encode Office namespace
type OfficeName string

//ExcelName type used to encode Excel namespace
type ExcelName string

//WordName type used to encode Word namespace
type WordName string

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
	WordName `xml:",attr"`
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
	Nested   []*Reserved
	InnerXML interface{}
}

//ShapeLayout is alias for CT_ShapeLayout
type ShapeLayout = Reserved

//Shape is alias for CT_Shape
type Shape = Reserved

//ShapeType is alias for CT_ShapeType
type ShapeType = Shape

//MarshalXMLAttr marshals VML namespace
func (r *Name) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"}
	return attr, nil
}

//MarshalXMLAttr marshals OfficeName namespace
func (r *OfficeName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"}
	return attr, nil
}

//MarshalXMLAttr marshals ExcelName namespace
func (r *ExcelName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "urn:schemas-microsoft-com:office:excel"}
	return attr, nil
}

//MarshalXMLAttr marshals WordName namespace
func (r *WordName) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "urn:schemas-microsoft-com:office:word"}
	return attr, nil
}

//MarshalXMLAttr marshals PowerPoint namespace
func (r *PowerPoint) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "urn:schemas-microsoft-com:office:powerpoint"}
	return attr, nil
}

func resolveName(a xml.Name) xml.Name {
	switch a.Space {
	case "urn:schemas-microsoft-com:vml":
		return xml.Name{Local: "v:" + a.Local}
	case "urn:schemas-microsoft-com:office:office":
		return xml.Name{Local: "o:" + a.Local}
	case "urn:schemas-microsoft-com:office:excel":
		return xml.Name{Local: "x:" + a.Local}
	case "urn:schemas-microsoft-com:office:word":
		return xml.Name{Local: "w:" + a.Local}
	case "urn:schemas-microsoft-com:office:powerpoint":
		return xml.Name{Local: "p:" + a.Local}
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

//MarshalXML marshals Reserved
func (r *Reserved) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = r.Name.Local
	start.Attr = make([]xml.Attr, 0, len(r.Attrs))

	var err error

	for k, v := range r.Attrs {
		attr := xml.Attr{ Name: xml.Name{Local: k} }

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

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, nested := range r.Nested {
		if err := e.EncodeElement(nested, start); err != nil {
			return err
		}
	}

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
