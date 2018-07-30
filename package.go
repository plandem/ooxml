package ooxml

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"reflect"
)

//Package is interface to expose some of PackageInfo methods via embedded struct
type Package interface {
	io.Closer
	Save() error
	SaveAs(fileName string) error
}

//DocumentFactoryFn is factory to in a OOXML type specific type
type DocumentFactoryFn func(pkg *PackageInfo) (interface{}, error)

//DocumentValidatorFn is callback to validate OOXML document. Using right before saving
type DocumentValidatorFn func() error

//PackageInfo holds all required information for OOXML package
type PackageInfo struct {
	Validator     DocumentValidatorFn
	reader        interface{}
	contentTypes  *ContentTypes
	relationships *Relationships
	files         map[string]interface{}
	fileName      string
}

var _ Package = (*PackageInfo)(nil)

//ErrorUnknownPackage returns a common error if DocumentFactoryFn returned invalid result
func ErrorUnknownPackage(p interface{}) error {
	return fmt.Errorf("unknown type of document, expects: %s", reflect.Indirect(reflect.ValueOf(p)).Type().Name())
}

//NewPackage returns a new package with zip reader if there is any
func NewPackage(reader interface{}) *PackageInfo {
	pkg := &PackageInfo{}

	var zipReader *zip.Reader

	switch rt := reader.(type) {
	case *zip.Reader:
		pkg.reader = rt
		zipReader = rt
	case *zip.ReadCloser:
		pkg.reader = rt
		zipReader = &rt.Reader
	}

	pkg.files = make(map[string]interface{})

	if pkg.reader != nil {
		//if there is a reader, than populate files with minimal information
		for _, f := range zipReader.File {
			pkg.files[f.Name] = f

			switch {
			case f.Name == "_rels/.rels":
				pkg.relationships = NewRelationships(f, pkg)
			case f.Name == "[Content_Types].xml":
				pkg.contentTypes = newContentTypes(f, pkg)
			}
		}
	} else {
		//if it's a new package, then add minimal required information for any OOXML document
		pkg.initPackage()
	}

	return pkg
}

//Open opens a file with fileName or io.Reader and returns an instance of document
func Open(f interface{}, docFactory DocumentFactoryFn) (interface{}, error) {
	var pkg *PackageInfo

	switch ft := f.(type) {
	case string:
		//f is name of file to open
		zipFile, err := zip.OpenReader(ft)
		if err != nil {
			return nil, err
		}

		pkg = NewPackage(zipFile)
		pkg.fileName = ft
	case io.Reader:
		//f is reader to read from
		b, err := ioutil.ReadAll(ft)
		if err != nil {
			panic(err)
		}

		readerAt := bytes.NewReader(b)
		zipReader, err := zip.NewReader(readerAt, int64(readerAt.Len()))
		if err != nil {
			return nil, err
		}

		pkg = NewPackage(zipReader)
	default:
		return nil, errors.New("unsupported type of f. It must be name of file or io.Reader")
	}

	return docFactory(pkg)
}

//IsNew returns true if package is a new one or false in other case
func (pkg *PackageInfo) IsNew() bool {
	return pkg.reader == nil
}

//Close closes current OOXML package
func (pkg *PackageInfo) Close() error {
	//close all opened reading streams
	for _, content := range pkg.files {
		if sr, ok := content.(*StreamFileReader); ok {
			sr.Close()
		}
	}

	//close zip reader
	if closer, ok := pkg.reader.(*zip.ReadCloser); ok {
		return closer.Close()
	}

	return nil
}

//Save saves current OOXML package
func (pkg *PackageInfo) Save() error {
	if pkg.fileName == "" {
		return fmt.Errorf("no filename defined for file. Try to use SaveAs")
	}

	//create file with a temp name
	tmpFile, err := ioutil.TempFile(path.Dir(pkg.fileName), path.Base(pkg.fileName))
	if err != nil {
		return err
	}

	//save content
	err = pkg.SavePackage(tmpFile)
	if err != nil {
		os.Remove(tmpFile.Name())
		return err
	}

	tmpFile.Close()

	//rename temp file into original name
	return os.Rename(tmpFile.Name(), pkg.fileName)
}

//SaveAs saves current OOXML package with fileName
func (pkg *PackageInfo) SaveAs(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer f.Close()
	return pkg.SavePackage(f)
}

//FileName is a private method that returns filename of opened file
func (pkg *PackageInfo) FileName() string {
	return pkg.fileName
}

//Add is a private method that adds a file to a package
func (pkg *PackageInfo) Add(fileName string, content interface{}) {
	pkg.files[fileName] = content
}

//File is a private method that returns file with required name
func (pkg *PackageInfo) File(fileName string) interface{} {
	for pkgFileName, content := range pkg.files {
		if pkgFileName == fileName {
			return content
		}
	}

	return nil
}

//Remove is a private method that removes file from a package
func (pkg *PackageInfo) Remove(fileName string) {
	if _, ok := pkg.files[fileName]; ok {
		//remove content
		delete(pkg.files, fileName)

		//remove info about type of content
		pkg.ContentTypes().RemoveContent(fileName)
	}
}

//Files is a private method to get list of all files inside of package
func (pkg *PackageInfo) Files() map[string]interface{} {
	return pkg.files
}

//SavePackage is private method with implementation of saving OOXML document to file
func (pkg *PackageInfo) SavePackage(f io.Writer) error {
	//If there is a validator, then validate and exit if there is any error and package can't be saved
	if pkg.Validator != nil {
		if err := pkg.Validator(); err != nil {
			return err
		}
	}

	//files holds differ kind of information:
	// 1) pointers to original files (*zip.File) that where not changed and must be coped as is
	// 2) pointers to write only files (*StreamFileWriter) that require callback execution to save information
	// 3) pointers to objects that must be marshaled to get content for a new file

	var err error

	buf := bytes.NewBuffer(nil)
	zipper := zip.NewWriter(buf)

	//add files to zip
	for fileName, content := range pkg.files {
		switch ft := content.(type) {
		case *zip.File:
			//file was not updated, so lets copy it as is
			err = CopyZipFile(ft, zipper)
		case *StreamFileWriter:
			//file was created as write stream, so need an additional callback execution to finalize postponed updates
			err = ft.Save(zipper)
		default:
			//file was probably updated, so let's marshal it and save with a new content
			err = MarshalZipFile(fileName, content, zipper)
		}
	}

	//looks like zip file was successfully created without any errors
	err = zipper.Close()
	if err != nil {
		return err
	}

	//physically save file
	_, err = buf.WriteTo(f)
	return err
}

//Relationships is a getter that returns top-level relationships of package
func (pkg *PackageInfo) Relationships() *Relationships {
	return pkg.relationships
}

//ContentTypes is a getter that returns content types of package
func (pkg *PackageInfo) ContentTypes() *ContentTypes {
	return pkg.contentTypes
}

//initPackage populates package with minimal required information for any OOXML document
func (pkg *PackageInfo) initPackage() {
	//content types must be initialized first of all - other types will use it
	pkg.contentTypes = newContentTypes("[Content_Types].xml", pkg)

	//register top-level relations
	pkg.relationships = NewRelationships("_rels/.rels", pkg)

	//add some default types
	pkg.contentTypes.RegisterType("rels", "application/vnd.openxmlformats-package.relationships+xml")
	pkg.contentTypes.RegisterType("vml", "application/vnd.openxmlformats-officedocument.vmlDrawing")
	pkg.contentTypes.RegisterType("png", "image/png")
	pkg.contentTypes.RegisterType("jpeg", "image/jpeg")
	pkg.contentTypes.RegisterType("jpg", "image/jpeg")
	pkg.contentTypes.RegisterType("gif", "image/gif")
	pkg.contentTypes.RegisterType("xml", "application/xml")
}
