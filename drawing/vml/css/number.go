package css

import (
	"fmt"
	"regexp"
	"strconv"
)

//Number is helper type which allow to encode integer as value in pixels and float as value in points. Eg. 10 => 10px, 10.5 => 10.5pt
type Number interface{}

//NumberCm is helper type to encode 'cm' numbers
type NumberCm float64

//NumberCm is helper type to encode 'mm' numbers
type NumberMm float64

//NumberCm is helper type to encode 'in' numbers
type NumberIn float64

//NumberCm is helper type to encode 'pt' numbers
type NumberPt float64

//NumberCm is helper type to encode 'pc' numbers
type NumberPc float64

//NumberCm is helper type to encode 'px' numbers
type NumberPx int

var (
	regExpNumber = regexp.MustCompile("^([0-9.]+)(cm|mm|in|pt|pc|px)?$")
)

func fromNumber(n Number) string {
	switch v := n.(type) {
	case NumberCm:
		return fmt.Sprintf("%.2fcm", v)
	case NumberMm:
		return fmt.Sprintf("%.2fmm", v)
	case NumberIn:
		return fmt.Sprintf("%.2fin", v)
	case NumberPt:
		return fmt.Sprintf("%.2fpt", v)
	case NumberPc:
		return fmt.Sprintf("%.2fpc", v)
	case NumberPx:
		return fmt.Sprintf("%dpx", v)
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
		switch parsed[2] {
		case "cm":
			if cm, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return NumberCm(cm)
			}
		case "mm":
			if mm, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return NumberMm(mm)
			}
		case "in":
			if in, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return NumberIn(in)
			}
		case "pt":
			if pt, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return NumberPt(pt)
			}
		case "pc":
			if pc, err := strconv.ParseFloat(parsed[1], 10); err == nil {
				return NumberPc(pc)
			}
		case "px":
			fallthrough
		default:
			if num, err := strconv.ParseInt(parsed[1], 10, 64); err == nil {
				return NumberPx(num)
			}
		}
	}

	return NumberPx(0)
}
