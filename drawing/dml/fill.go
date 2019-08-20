// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

//Fill is a direct mapping of XSD EG_FillProperties
type Fill struct {
	Blip *BlipFill `xml:"blipFill,omitempty"`
}
