// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package css

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//Style is helper type for AG_Style
type Style struct {
	Position          Position        `css:"position"`
	Left              Number          `css:"left"`
	MarginLeft        Number          `css:"margin-left"`
	Top               Number          `css:"top"`
	MarginTop         Number          `css:"margin-top"`
	Right             Number          `css:"right"`
	MarginRight       Number          `css:"margin-right"`
	Bottom            Number          `css:"bottom"`
	MarginBottom      Number          `css:"margin-bottom"`
	Width             Number          `css:"width"`
	Height            Number          `css:"height"`
	ZIndex            int             `css:"z-index"`
	Visible           Visibility      `css:"visibility"`
	Flip              string          `css:"flip"`
	Font              string          `css:"font"`
	TextDecoration    string          `css:"text-decoration"`
	Trim              ml.TriStateType `css:"trim"`
	XScale            ml.TriStateType `css:"xscale"`
	MSOFitShapeToText ml.TriStateType `css:"mso-fit-shape-to-text"`
	MSOFitTextToShape ml.TriStateType `css:"mso-fit-text-to-shape"`
	MSOTextShadow     ml.TriStateType `css:"mso-text-shadow"`
	MSODirectionAlt   string          `css:"mso-direction-alt"`
	MSOLayoutFlowAlt  string          `css:"mso-layout-flow-alt"`
	MSONextTextbox    string          `css:"mso-next-textbox"`
	MSORotate         float64         `css:"mso-rotate"`
	MSOTextScale      float64         `css:"mso-text-scale"`
}

var (
	regExpCss = regexp.MustCompile("(?P<key>[a-zA-z-]+):(?P<value>[0-9a-zA-Z.%-- ]+)+")
)

//NewStyle decodes VML CSS string into Style type
func NewStyle(s string) *Style {
	parsed := regExpCss.FindAllStringSubmatch(s, -1)
	mapped := make(map[string]string)
	for _, p := range parsed {
		mapped[p[1]] = strings.TrimSpace(p[2])
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
				case Position:
					field.Set(reflect.ValueOf(toPosition[value]))
				case Visibility:
					field.Set(reflect.ValueOf(toVisibility[value]))
				case Number:
					field.Set(reflect.ValueOf(NewNumber(value)))
				case ml.TriStateType:
					field.Set(reflect.ValueOf(ml.TriState(value)))
				case string:
					field.SetString(value)
				case float32, float64:
					if f, ok := strconv.ParseFloat(value, 10); ok == nil {
						field.SetFloat(f)
					}
				case byte, uint, uint16, uint32, uint64, int, int8, int16, int32, int64:
					if i, ok := strconv.ParseInt(value, 10, 64); ok == nil {
						field.SetInt(int64(i))
					}
				}
			}
		}
	}

	return &style
}

//String is alias for Encode that return string version of styles
func (s Style) String() string {
	return s.encode()
}

func (s Style) encode() string {
	var result []string

	v := reflect.ValueOf(&s).Elem()
	vt := reflect.TypeOf(s)

	for i := 0; i < vt.NumField(); i++ {
		tags := vt.Field(i).Tag
		field := v.Field(i)
		if cssName, ok := tags.Lookup("css"); ok && cssName != "" && field.IsValid() && !ooxml.IsEmptyValue(field) {
			switch t := field.Interface().(type) {
			case ml.TriStateType:
				if t != ml.TriStateBlank {
					result = append(result, fmt.Sprintf("%s:%s", cssName, t.String()))
				}
			case Number:
				if t.unit != unitUnknown {
					result = append(result, fmt.Sprintf("%s:%s", cssName, t.String()))
				}
			default:
				result = append(result, fmt.Sprintf("%s:%v", cssName, field.Interface()))
			}
		}
	}

	return strings.Join(result, ";")
}

//MarshalXMLAttr marshal Style
func (s *Style) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if v := s.String(); len(v) > 0 {
		return xml.Attr{Name: name, Value: v}, nil
	}

	return xml.Attr{}, nil
}

//UnmarshalXMLAttr unmarshal Style
func (s *Style) UnmarshalXMLAttr(attr xml.Attr) error {
	*s = *NewStyle(attr.Value)
	return nil
}
