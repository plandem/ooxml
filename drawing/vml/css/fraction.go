// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package css

import (
	"encoding/xml"
	"regexp"
	"strconv"
)

//Fraction is helper type to encode VgFraction type, that can be from 0.0 to 1.0 or in percentage, e.g. 50%
//N.B.: Fraction always transforms percentage to float, e.g. 50% -> 0.5
type Fraction float32

var (
	regExpFraction = regexp.MustCompile("^([0-9.-]+)(%)?$")
)

//UnmarshalXMLAttr unmarshal Fraction
func (f *Fraction) UnmarshalXMLAttr(attr xml.Attr) error {
	parsed := regExpFraction.FindStringSubmatch(attr.Value)
	if parsed != nil {
		if v, err := strconv.ParseFloat(parsed[1], 10); err != nil {
			return err
		} else {
			if parsed[2] == "%" {
				*f = Fraction(v / 100)
			} else {
				*f = Fraction(v)
			}
		}
	}

	return nil
}
