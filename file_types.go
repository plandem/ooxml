// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ooxml

import "github.com/plandem/ooxml/ml"

//List of all supported ContentType for any office file - excel, word, powerpoint
const (
	ContentTypeVmlDrawing    ml.ContentType = "application/vnd.openxmlformats-officedocument.vmlDrawing"
	ContentTypeRelationships ml.ContentType = "application/vnd.openxmlformats-package.relationships+xml"
)
