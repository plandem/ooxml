package css

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//Number is helper type which allow to encode integer as value in pixels and float as value in points. Eg. 10 => 10px, 10.5 => 10.5pt
type Number interface{}

type position byte
type visibility byte

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
}

const (
	PositionStatic position = iota
	PositionAbsolute
	PositionRelative

	VisibilityInherit visibility = iota
	VisibilityHidden
	VisibilityVisible
	VisibilityCollapse
)

var (
	toVisibility   map[string]visibility
	fromVisibility map[visibility]string

	toPosition   map[string]position
	fromPosition map[position]string

	regExpCss    = regexp.MustCompile("(?P<key>[a-zA-z-]+):(?P<value>[0-9a-z.]+)+")
	regExpNumber = regexp.MustCompile("^([0-9.]+)(cm|mm|in|pt|pc|px)?$")
)

func init() {
	//setup visibility
	fromVisibility = map[visibility]string{
		VisibilityInherit:  "inherit",
		VisibilityHidden:   "hidden",
		VisibilityVisible:  "visible",
		VisibilityCollapse: "collapse",
	}

	toVisibility = make(map[string]visibility, len(fromVisibility))
	for k, v := range fromVisibility {
		toVisibility[v] = k
	}

	//setup position
	fromPosition = map[position]string{
		PositionStatic:   "static",
		PositionAbsolute: "absolute",
		PositionRelative: "relative",
	}

	toPosition = make(map[string]position, len(fromPosition))
	for k, v := range fromPosition {
		toPosition[v] = k
	}
}

func (e visibility) String() string {
	return fromVisibility[e]
}

func (e position) String() string {
	return fromPosition[e]
}

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

//Encode encodes Style type into VML CSS string
func (s Style) Encode() string {
	var result []string

	v := reflect.ValueOf(&s).Elem()
	vt := reflect.TypeOf(s)

	for i := 0; i < vt.NumField(); i++ {
		tags := vt.Field(i).Tag
		field := v.Field(i)
		if cssName, ok := tags.Lookup("css"); ok && cssName != "" && field.IsValid() && !isEmptyValue(field) {
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

func fromNumber(n Number) string {
	switch v := n.(type) {
	case float32:
		return fmt.Sprintf("%.2fpt", v)
	case float64:
		return fmt.Sprintf("%.2fpt", v)
	case uint:
		return fmt.Sprintf("%dpx", v)
	case uint8:
		return fmt.Sprintf("%dpx", v)
	case uint16:
		return fmt.Sprintf("%dpx", v)
	case uint32:
		return fmt.Sprintf("%dpx", v)
	case uint64:
		return fmt.Sprintf("%dpx", v)
	case int:
		return fmt.Sprintf("%dpx", v)
	case int8:
		return fmt.Sprintf("%dpx", v)
	case int16:
		return fmt.Sprintf("%dpx", v)
	case int32:
		return fmt.Sprintf("%dpx", v)
	case int64:
		return fmt.Sprintf("%dpx", v)
	}

	return ""
}

func toNumber(n string) Number {
	parsed := regExpNumber.FindStringSubmatch(n)
	if parsed != nil {
		if parsed[2] == "" || parsed[2] == "px" {
			if pxNum, err := strconv.ParseInt(parsed[1], 10, 64); err == nil {
				return int64(pxNum)
			}
		}

		if pxFloat, err := strconv.ParseFloat(parsed[1], 10); err == nil {
			return float64(pxFloat)
		}
	}

	return 0
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
