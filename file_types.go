package ooxml

import "github.com/plandem/ooxml/ml"

//List of all supported ContentType for any office file - excel, word, powerpoint
const (
	ContentTypeVmlDrawing    ml.ContentType = "application/vnd.openxmlformats-officedocument.vmlDrawing"
	ContentTypeRelationships ml.ContentType = "application/vnd.openxmlformats-package.relationships+xml"
)
