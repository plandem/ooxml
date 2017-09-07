package ooxml

import (
	"archive/zip"
)

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

//MarkAsUpdated marks file as updated, so content will be replaced with source's content during packing document.
//Works only with new files or files that where fully loaded (via LoadIfRequired).
//Files that were manually opened via Open can't be marked as updated
func (pf *PackageFile) MarkAsUpdated() {
	if pf.zipFile == nil {
		pf.pkg.Add(pf.fileName, pf.source)
	}
}

//LoadIfRequired lazy loads whole content of file into target and call required callback if there is any
func (pf *PackageFile) LoadIfRequired(callback func()) {
	if !pf.isNew {
		//is file already opened as read stream?
		if _, ok := pf.source.(*StreamFileReader); ok {
			panic("Can't load file that already opened as stream.")
		}

		//first time request?
		if pf.zipFile != nil {
			if err := UnmarshalZipFile(pf.zipFile, pf.target); err != nil {
				panic(err)
			}

			pf.zipFile = nil

			if callback != nil {
				callback()
			}
		}
	}
}

//ReadStream opens a zip file for manual reading as stream and return *StreamReader for it
//Files that were opened as stream can't be marked as updated via MarkAsUpdated and will be saved as is
//File can be opened as stream only once, any further requests will return previously opened stream
func (pf *PackageFile) ReadStream() *StreamReader {
	if pf.isNew {
		panic("Can't open a new file for reading.")
	}

	//if file was already opened as stream, then return it
	if s, ok := pf.source.(*StreamFileReader); ok {
		return s.StreamReader
	}

	if pf.zipFile == nil {
		panic("Can't open as stream file that was already fully loaded.")
	}

	stream, err := NewStreamFileReader(pf.zipFile)
	if err != nil {
		panic(err)
	}

	pf.source = stream
	pf.pkg.Add(pf.fileName, pf.source)
	return stream.StreamReader
}

//WriteStream creates a zip file for manual writing as stream and return StreamWriter for it
//File can be created as stream only once, any further requests will return previously created stream
func (pf *PackageFile) WriteStream(finalize StreamWriterCallback) *StreamWriter {
	if !pf.isNew {
		panic("Can't overwrite already existing file.")
	}

	//is stream already created, then return it
	if s, ok := pf.source.(*StreamFileWriter); ok {
		return s.StreamWriter
	}

	stream := NewStreamFileWriter(pf.fileName, finalize)
	pf.source = stream
	pf.pkg.Add(pf.fileName, pf.source)

	return stream.StreamWriter
}
