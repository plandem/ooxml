// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"fmt"
	"strconv"
	"strings"
)

//ClientDataAnchor is helper class for anchor information of ClientData
type ClientDataAnchor struct {
	LeftColumn   int
	LeftOffset   int
	TopRow       int
	TopOffset    int
	RightColumn  int
	RightOffset  int
	BottomRow    int
	BottomOffset int
}

//String returns string presentation of ClientDataAnchor
func (t ClientDataAnchor) String() string {
	return fmt.Sprintf("%d, %d, %d, %d, %d, %d, %d, %d",
		t.LeftColumn,
		t.LeftOffset,
		t.TopRow,
		t.TopOffset,
		t.RightColumn,
		t.RightOffset,
		t.BottomRow,
		t.BottomOffset,
	)
}

//StringToAnchor converts string into ClientDataAnchor object
func StringToClientDataAnchor(s string) (ClientDataAnchor, error) {
	a := ClientDataAnchor{}
	numbers := strings.Split(s, ",")
	for i, s := range numbers {
		if n, err := strconv.Atoi(strings.TrimSpace(s)); err != nil {
			return ClientDataAnchor{}, err
		} else {
			switch i {
			case 0:
				a.LeftColumn = n
			case 1:
				a.LeftOffset = n
			case 2:
				a.TopRow = n
			case 3:
				a.TopOffset = n
			case 4:
				a.RightColumn = n
			case 5:
				a.RightOffset = n
			case 6:
				a.BottomRow = n
			case 7:
				a.BottomOffset = n
			}
		}
	}

	return a, nil
}
