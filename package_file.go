package ooxml

import (
	"archive/zip"
	"fmt"
)

//PackageFileLoadFn is a callback that will be called right after loading related zipped file, if there is any
type PackageFileLoadFn func()

//PackageFile is helper object that implements common functionality for any file of package. E.g. lazy loading, marking as updated.
type PackageFile struct {
	fileName string
	zipFile  *zip.File
	target   interface{}
	source   interface{}
	pkg      *PackageInfo
	isNew    bool
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

	if zf, ok := f.(*zip.File); ok && zf != nil {
		pkgFile.fileName = zf.Name
		pkgFile.zipFile = zf
		pkgFile.isNew = false
	} else if fileName, ok := f.(string); ok {
		pkgFile.fileName = fileName
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

//LoadIfRequired lazy loads content of file into target and call required callback if there is any
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

//MarkAsUpdated marks file as updated, so content will be replaced with source's content during packing document
func (pf *PackageFile) MarkAsUpdated() {
	pf.pkg.Add(pf.fileName, pf.source)
}
