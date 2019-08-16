// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dml

//Coordinate is a direct mapping of XSD ST_Coordinate
//Office will read either a length followed by a unit or EMUs with no unit present, but will write only EMUs when no units are present.
type Coordinate int

