package ml

import (
	"encoding/xml"
	"strconv"
)

//Name is alias for xml.Name to decrease number of imports
type Name xml.Name

//Attr is alias for xml.Attr to decrease number of imports
type Attr xml.Attr

//CharData is alias for xml.CharData to decrease number of imports
type CharData xml.CharData

//OptionalIndex is custom type to allow encode/decode optional 0-based indexes
type OptionalIndex *int

//Property is common type of property for strings. E.g.: <propName val="abcdef"/>
type Property string

//PropertyBool is special type of property for booleans. E.g.: <propName val="true"/>, <propName/>, <propName val="false"/>
type PropertyBool bool

//PropertyInt is special type of property for integers. E.g.: <propName val="123"/>
type PropertyInt int

//PropertyDouble is special type of property for doubles. E.g.: <propName val="123.456"/>
type PropertyDouble float64

//AttrPreserveSpace is common attr to preserve space
var AttrPreserveSpace = xml.Attr{
	Name:  xml.Name{Space: NamespaceXML, Local: "space"},
	Value: "preserve",
}

func (p *Property) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"}, Value: string(*p)})
	return e.EncodeElement(struct{}{}, start)
}

func (p *Property) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if len(start.Attr) > 0 {
		*p = Property(start.Attr[0].Value)
	}

	return d.Skip()
}

func (p *PropertyBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"}, Value: strconv.FormatBool(bool(*p))})
	return e.EncodeElement(struct{}{}, start)
}

func (p *PropertyBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if len(start.Attr) > 0 {
		if b, err := strconv.ParseBool(start.Attr[0].Value); err == nil {
			*p = PropertyBool(b)
		}
	} else {
		*p = PropertyBool(true)
	}

	return d.Skip()
}

func (p *PropertyInt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"}, Value: strconv.FormatInt(int64(*p), 10)})
	return e.EncodeElement(struct{}{}, start)
}

func (p *PropertyInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if len(start.Attr) > 0 {
		if i, err := strconv.ParseInt(start.Attr[0].Value, 10, 0); err == nil {
			*p = PropertyInt(i)
		}
	}

	return d.Skip()
}

func (p *PropertyDouble) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"}, Value: strconv.FormatFloat(float64(*p), 'f', -1, 64)})
	return e.EncodeElement(struct{}{}, start)
}

func (p *PropertyDouble) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if len(start.Attr) > 0 {
		if f, err := strconv.ParseFloat(start.Attr[0].Value, 64); err == nil {
			*p = PropertyDouble(f)
		}
	}

	return d.Skip()
}
