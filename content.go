package ooxml

import (
	"github.com/plandem/ooxml/ml"
)

//ContentTypes is helper object that implements some functionality for content types that is a required part of any OOXML document
type ContentTypes struct {
	ml   ml.ContentTypes
	pkg  *PackageInfo
	file *PackageFile
}

//newContentTypes creates and returns content types information
func newContentTypes(f interface{}, pkg *PackageInfo) *ContentTypes {
	content := &ContentTypes{
		pkg: pkg,
	}

	content.file = NewPackageFile(pkg, f, &content.ml, nil)
	content.file.LoadIfRequired(nil)
	return content
}

//RegisterType adds information about a new type of content if there is no such type already
func (ct *ContentTypes) RegisterType(extension string, contentType ml.ContentType) {
	//check if there is type with such extension already, and if it's here then ignore
	for _, def := range ct.ml.Defaults {
		if def.Extension == extension {
			return
		}
	}

	ct.ml.Defaults = append(ct.ml.Defaults, &ml.TypeDefault{
		Extension:   extension,
		ContentType: contentType,
	})

	ct.file.MarkAsUpdated()
}

//RegisterContent adds information about a new content with fileName of contentType
func (ct *ContentTypes) RegisterContent(fileName string, contentType ml.ContentType) {
	if fileName[0] != '/' {
		fileName = "/" + fileName
	}

	ct.ml.Overrides = append(ct.ml.Overrides, &ml.TypeOverride{
		PartName:    fileName,
		ContentType: contentType,
	})

	ct.file.MarkAsUpdated()
}

//RemoveContent removes information about a content with fileName
func (ct *ContentTypes) RemoveContent(fileName string) {
	for i, part := range ct.ml.Overrides {
		if part.PartName == fileName {
			ct.ml.Overrides = append(ct.ml.Overrides[:i], ct.ml.Overrides[i+1:]...)
			ct.file.MarkAsUpdated()
			return
		}
	}
}
