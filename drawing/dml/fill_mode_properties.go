// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

import (
	"github.com/plandem/ooxml/ml"
)

//fillModeProperties is a direct mapping of EG_FillModeProperties
type fillModeProperties struct {
	Stretch         *ml.Reserved `xml:"stretch,omitempty"`
	Tile            *ml.Reserved `xml:"tile,omitempty"`
}