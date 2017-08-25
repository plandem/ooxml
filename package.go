package shared

import (
	"os"
	"io"
	"fmt"
	"bytes"
	"reflect"
	"io/ioutil"
	"archive/zip"
)

//Package is interface to expose some of PackageInfo methods via embedded struct
type Package interface {
	Close()
	Save() (error)
	SaveAs(fileName string) (error)
}

//DocumentFactoryFn is factory to in a OOXML type specific type
type DocumentFactoryFn func(pkg *PackageInfo) (interface{}, error)

type DocumentValidatorFn func() (error)

//PackageReader is a wrapper around ZIP file to unify proccessing
type PackageReader struct {
	*zip.Reader
	io.Closer
}

//PackageInfo holds all required information for OOXML package
type PackageInfo struct {
	Validator     DocumentValidatorFn
	reader        *PackageReader
	contentTypes  *ContentTypes
	relationships *Relationships
	files         map[string]interface{}
	fileName      string
}

//ErrorUnknownPackage returns a common error if DocumentFactoryFn returned invalid result
func ErrorUnknownPackage(p interface{}) (error) {
	return fmt.Errorf("Unknown type of document. Expects: %s", reflect.Indirect(reflect.ValueOf(p)).Type().Name())
}

//NewPackage returns a new package with zip reader if there is any
func NewPackage(reader interface{}) (*PackageInfo) {
	pkg := &PackageInfo{}

	if zipFile, ok := reader.(*zip.ReadCloser); ok {
		pkg = &PackageInfo{
			reader: &PackageReader{
				&zipFile.Reader,
				zipFile,
			},
		}
	} else if zipStream, ok := reader.(*zip.Reader); ok {
		pkg = &PackageInfo{
			reader: &PackageReader{
				zipStream,
				nil,
			},
		}
	}

	pkg.files = make(map[string]interface{})

	if pkg.reader != nil {
		//if there is a reader, than populate files with minimal information
		for _, f := range pkg.reader.File {
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

//OpenFile opens a file with fileName and returns an instance of document
func OpenFile(fileName string, docFactory DocumentFactoryFn) (interface{}, error) {
	zipFile, err := zip.OpenReader(fileName)
	if err != nil {
		return nil, err
	}

	pkg := NewPackage(zipFile)
	pkg.fileName = fileName
	return docFactory(pkg)
}

//OpenStream opens a zip stream and returns an instance of document
//Note: True streams can't be supported, due to zip files (can't be streamed)
func OpenStream(stream io.Reader, docFactory DocumentFactoryFn) (interface{}, error) {
	b, err := ioutil.ReadAll(stream)
	if err != nil {
		panic(err)
	}

	readerAt := bytes.NewReader(b)
	zipStream, err := zip.NewReader(readerAt, int64(readerAt.Len()))
	if err != nil {
		return nil, err
	}

	return docFactory(NewPackage(zipStream))
}

//IsNew returns true if package is a new one or false in other case
func (pkg *PackageInfo) IsNew() (bool) {
	return pkg.reader == nil
}

//Close closes current OOXML package
func (pkg *PackageInfo) Close() {
	if pkg.reader != nil && pkg.reader.Closer != nil {
		pkg.reader.Closer.Close()
	}
}

//Save saves current OOXML package
func (pkg *PackageInfo) Save() (error) {
	if pkg.fileName == "" {
		return fmt.Errorf("No filename defined for file, try to use SaveAs")
	}

	return pkg.SaveAs(pkg.fileName)
}

//SaveAs saves current OOXML package with fileName
func (pkg *PackageInfo) SaveAs(fileName string) (error) {
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

//AddFile is a private method that adds a file to a package
func (pkg *PackageInfo) Add(fileName string, content interface{}) {
	pkg.files[fileName] = content
}

//RemoveFile is a private method that removes  file from a package
func (pkg *PackageInfo) Remove(fileName string) {
	if _, ok := pkg.files[fileName]; ok {
		//remove content
		delete(pkg.files, fileName)

		//remove info about type of content
		pkg.ContentTypes().RemoveContent(fileName)
	}
}

//Files is a private method to get list of all files inside of package
func (pkg *PackageInfo) Files() (map[string]interface{}) {
	return pkg.files
}

//SavePackage is private method with implementation of saving OOXML document to file
func (pkg *PackageInfo) SavePackage(f io.Writer) (error) {
	//If there is a validator, then validate and exit if there is any error and package can't be saved
	if pkg.Validator != nil {
		if err := pkg.Validator(); err != nil {
			return err
		}
	}

	//files holds two kind of information:
	// 1) pointers to original files (*zip.File) that where not changed and must be coped as is
	// 2) pointers to structs that must be marshaled to get content for a new file

	var err error

	buf := bytes.NewBuffer(nil)
	zipper := zip.NewWriter(buf)

	//add files to zip
	for fileName, content := range pkg.files {
		if f, ok := content.(*zip.File); ok {
			//file was not updated, so lets copy it as is
			err = CopyZipFile(f, zipper)
		} else {
			//file was probably updated, so let's marshal it and save with a new content
			err = MarshalZipFile(fileName, content, zipper)
		}

		if err != nil {
			return err
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
func (pkg *PackageInfo) Relationships() (*Relationships) {
	return pkg.relationships
}

//ContentTypes is a getter that returns content types of package
func (pkg *PackageInfo) ContentTypes() (*ContentTypes) {
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
