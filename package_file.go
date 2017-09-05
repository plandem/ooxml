package ooxml

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
)

//PackageFileLoadFn is a callback that will be called right after loading related zipped file, if there is any
type PackageFileLoadFn func()

//PackageFile is helper object that implements common functionality for any file of package. E.g. lazy loading, marking as updated.
type PackageFile struct {
	//name of file
	fileName string

	//pointer to *zip.File or nil in case of a new file
	zipFile *zip.File

	//pointer to target-object to unmarshal content from a file
	target interface{}

	//pointer to source-object to marshal content into a file
	source interface{}

	//package owner of this file
	pkg *PackageInfo

	//flag that indicate new file or not
	isNew bool
}

//NewPackageFile creates and returns package file that attached target via file f with source of information to save
func NewPackageFile(pkg *PackageInfo, f interface{}, target interface{}, source interface{}) *PackageFile {
	if source == nil {
		source = target
	}

	pkgFile := &PackageFile{
		pkg:    pkg,
		target: target,
		source: source,
		isNew:  true,
	}

	if f != nil {
		switch ft := f.(type) {
		case *zip.File:
			pkgFile.fileName = ft.Name
			pkgFile.zipFile = ft
			pkgFile.isNew = false
		case string:
			pkgFile.fileName = ft
		}
	}

	if len(pkgFile.fileName) == 0 {
		panic("You must provide a file to use - zip.File for existing or filename for a new one.")
	}

	return pkgFile
}

//FileName returns name of file
func (pf *PackageFile) FileName() string {
	return pf.fileName
}

//IsNew returns true if this file is a new file or false in other case
func (pf *PackageFile) IsNew() bool {
	return pf.isNew
}

//MarkAsUpdated marks file as updated, so content will be replaced with source's content during packing document
func (pf *PackageFile) MarkAsUpdated() {
	if pf.zipFile == nil {
		//only new or fully loaded (via LoadIfRequired) can be marked as updated.
		pf.pkg.Add(pf.fileName, pf.source)
	}
}

//LoadIfRequired lazy loads whole content of file into target and call required callback if there is any
func (pf *PackageFile) LoadIfRequired(callback PackageFileLoadFn) {
	if pf.zipFile != nil {
		if err := UnmarshalZipFile(pf.zipFile, pf.target); err != nil {
			panic(fmt.Sprintln("Can't load zipped data: ", err))
		}

		pf.zipFile = nil

		if callback != nil {
			callback()
		}
	}
}

//Open opens a zip file for reading and return handler for it or error in case of any issue
func (pf *PackageFile) Open() (io.ReadCloser, error) {
	if pf.zipFile != nil {
		return pf.zipFile.Open()
	}

	return nil, errors.New("can't open zip file for reading")
}
