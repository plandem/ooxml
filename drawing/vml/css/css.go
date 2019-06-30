package css

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//Style is helper type for AG_Style
type Style struct {
	Position     position   `css:"position"`
	Left         Number     `css:"left"`
	MarginLeft   Number     `css:"margin-left"`
	Top          Number     `css:"top"`
	MarginTop    Number     `css:"margin-top"`
	Right        Number     `css:"right"`
	MarginRight  Number     `css:"margin-right"`
	Bottom       Number     `css:"bottom"`
	MarginBottom Number     `css:"margin-bottom"`
	Width        Number     `css:"width"`
	Height       Number     `css:"height"`
	ZIndex       int        `css:"z-index"`
	Visible      visibility `css:"visibility"`

	//TODO: add MSO-attributes
}

var (
	regExpCss = regexp.MustCompile("(?P<key>[a-zA-z-]+):(?P<value>[0-9a-z.]+)+")
)

//Decode decodes VML CSS string into Style type
func Decode(s string) Style {
	parsed := regExpCss.FindAllStringSubmatch(s, -1)
	mapped := make(map[string]string)
	for _, p := range parsed {
		mapped[p[1]] = p[2]
	}

	style := Style{}
	v := reflect.ValueOf(&style).Elem()
	vt := reflect.TypeOf(style)

	for i := 0; i < vt.NumField(); i++ {
		field := v.Field(i)
		tags := vt.Field(i).Tag

		if cssName, ok := tags.Lookup("css"); ok && cssName != "" {
			if value, ok := mapped[cssName]; ok {
				switch field.Interface().(type) {
				case position:
					field.Set(reflect.ValueOf(toPosition[value]))
				case visibility:
					field.Set(reflect.ValueOf(toVisibility[value]))
				case int:
					if i, ok := strconv.ParseInt(value, 10, 64); ok == nil {
						field.SetInt(int64(i))
					}
				default:
					field.Set(reflect.ValueOf(toNumber(value)))
				}
			}
		}
	}

	return style
}

//String is alias for Encode that return string version of styles
func (s Style) String() string {
	return s.Encode()
}

//Encode encodes Style type into VML CSS string
func (s Style) Encode() string {
	var result []string

	v := reflect.ValueOf(&s).Elem()
	vt := reflect.TypeOf(s)

	for i := 0; i < vt.NumField(); i++ {
		tags := vt.Field(i).Tag
		field := v.Field(i)
		if cssName, ok := tags.Lookup("css"); ok && cssName != "" && field.IsValid() && !ooxml.IsEmptyValue(field) {
			switch field.Kind() {
			case reflect.Interface:
				result = append(result, fmt.Sprintf("%s:%s", cssName, fromNumber(field.Interface())))
			default:
				result = append(result, fmt.Sprintf("%s:%+v", cssName, field.Interface()))
			}
		}
	}

	return strings.Join(result, ";")
}

//MarshalXMLAttr marshal Style
func (s *Style) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if s != nil {
		if v := s.Encode(); len(v) > 0 {
			return xml.Attr{Name: name, Value: v}, nil
		}
	}

	return xml.Attr{}, nil
}

//UnmarshalXMLAttr unmarshal Style
func (s *Style) UnmarshalXMLAttr(attr xml.Attr) error {
	*s = Decode(attr.Value)
	return nil
}
