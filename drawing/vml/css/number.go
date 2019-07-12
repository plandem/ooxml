// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package css

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
)

//Number is helper type which allow to encode numbers with units. To simplify, integer as value in pixels and float as value in points. Eg. 10 => 10px, 10.5 => 10.5pt
type Number struct {
	val  interface{}
	unit numberUnit
}

type numberUnit byte

const (
	unitUnknown    numberUnit = iota
	UnitPx                    //used to encode 'px' numbers
	UnitCm                    //used to encode 'cm' numbers
	UnitMm                    //used to encode 'mm' numbers
	UnitIn                    //used to encode 'in' numbers
	UnitPt                    //used to encode 'pt' numbers
	UnitPc                    //used to encode 'pc' numbers
	UnitPercentage            //used to encode '%' numbers
)

var (
	regExpNumber = regexp.MustCompile("^([0-9.]+)(cm|mm|in|pt|pc|px|%)?$")
)

//NewNumber returns a Number type for provided value
func NewNumber(n interface{}, o ...numberUnit) Number {
	if s, ok := n.(string); ok {
		return fromString(s)
	}

	var u numberUnit
	if len(o) > 0 {
		u = o[0]
	} else {
		u = unitUnknown
	}

	//for numeric types we just need to resolve type of unit
	switch n.(type) {
	case float32, float64:
		if u == unitUnknown || u == UnitPx {
			u = UnitPt
		}

	case byte, uint, uint16, uint32, uint64, int, int8, int16, int32, int64:
		if u == unitUnknown {
			u = UnitPx
		}

	default:
		n = 0
		u = UnitPx
	}

	return Number{n, u}
}

//String returns string presentation of Number
func (t Number) String() string {
	switch t.unit {
	case UnitCm:
		return fmt.Sprintf("%vcm", t.val)
	case UnitMm:
		return fmt.Sprintf("%vmm", t.val)
	case UnitIn:
		return fmt.Sprintf("%vin", t.val)
	case UnitPt:
		return fmt.Sprintf("%vpt", t.val)
	case UnitPc:
		return fmt.Sprintf("%vpc", t.val)
	case UnitPx:
		return fmt.Sprintf("%vpx", t.val)
	case UnitPercentage:
		return fmt.Sprintf("%v%%", t.val)
	}

	return ""
}

//MarshalXMLAttr marshal Number
func (t Number) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t.unit == unitUnknown {
		return xml.Attr{}, nil
	}

	return xml.Attr{Name: name, Value: t.String()}, nil
}

//UnmarshalXMLAttr unmarshal Number
func (t *Number) UnmarshalXMLAttr(attr xml.Attr) error {
	*t = NewNumber(attr.Value)
	return nil
}

//convert string into Number
func fromString(n string) Number {
	parsed := regExpNumber.FindStringSubmatch(n)
	if parsed != nil {
		switch parsed[2] {
		case "cm":
			if cm, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return Number{cm, UnitCm}
			}
		case "mm":
			if mm, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return Number{mm, UnitMm}
			}
		case "in":
			if in, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return Number{in, UnitIn}
			}
		case "pt":
			if pt, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return Number{pt, UnitPt}
			}
		case "pc":
			if pc, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return Number{pc, UnitPc}
			}
		case "%":
			if num, err := strconv.ParseInt(parsed[1], 10, 64); err == nil {
				return Number{int(num), UnitPercentage}
			} else {
				if num, err := strconv.ParseFloat(parsed[1], 10); err == nil {
					return Number{num, UnitPercentage}
				}
			}
		case "px":
			fallthrough
		default:
			if num, err := strconv.ParseInt(parsed[1], 10, 64); err == nil {
				return Number{int(num), UnitPx}
			} else {
				if num, err := strconv.ParseFloat(parsed[1], 10); err == nil {
					return Number{num, UnitPt}
				}
			}
		}
	}

	return Number{0, UnitPx}
}
