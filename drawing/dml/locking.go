// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

//Locking is a direct mapping of XSD AG_Locking
type Locking struct {
	NoGroup            bool `xml:"noGrp,attr,omitempty"`
	NoSelect           bool `xml:"noSelect,attr,omitempty"`
	NoRotate           bool `xml:"noRot,attr,omitempty"`
	NoChangeAspect     bool `xml:"noChangeAspect,attr,omitempty"`
	NoMove             bool `xml:"noMove,attr,omitempty"`
	NoResize           bool `xml:"noResize,attr,omitempty"`
	NoEditPoints       bool `xml:"noEditPoints,attr,omitempty"`
	NoAdjustHandles    bool `xml:"noAdjustHandles,attr,omitempty"`
	NoChangeArrowheads bool `xml:"noChangeArrowheads,attr,omitempty"`
	NoChangeShapeType  bool `xml:"noChangeShapeType,attr,omitempty"`
}
