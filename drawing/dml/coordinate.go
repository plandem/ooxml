// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"encoding/xml"
	"strconv"
	"strings"
)

//Coordinate is a direct mapping of XSD ST_Coordinate
//Office will read either a length followed by a unit or EMUs with no unit present, but will write only EMUs when no units are present.
type Coordinate int

//CoordinateWithUnit converts coordinate with units into EMU
func CoordinateWithUnit(c string) (Coordinate, error) {
	c = strings.ToLower(c)

	//TODO: add parser of units and convert to EMU
	//...

	i, err := strconv.Atoi(c)
	return Coordinate(i), err
}

//UnmarshalXMLAttr unmarshal text into Coordinate in EMU
func (c *Coordinate) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	*c, err = CoordinateWithUnit(attr.Value)
	return
}
