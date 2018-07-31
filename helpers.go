package ooxml

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"regexp"
	"strconv"
)

//MarshalPreparation is interface that must be implemented by complex data if it requires some preparation steps before marshaling
type MarshalPreparation interface {
	BeforeMarshalXML() interface{}
}

//MarshalFixation is interface that must be implemented by complex data if it requires some after marshal fixation
type MarshalFixation interface {
	AfterMarshalXML(content []byte) []byte
}

//GetLettersFn is a strings.Map walker to return [a-zA-Z] runes from string
func GetLettersFn(rune rune) rune {
	switch {
	case 'A' <= rune && rune <= 'Z':
		return rune
	case 'a' <= rune && rune <= 'z':
		return rune - 32
	}
	return -1
}

//GetNumbersFn is a strings.Map walker to return [0-9] runes from string
func GetNumbersFn(rune rune) rune {
	if rune >= 48 && rune < 58 {
		return rune
	}
	return -1
}

//UnmarshalZipFile unpacks a zip file into target object
func UnmarshalZipFile(f *zip.File, target interface{}) error {
	xmlReader, err := f.Open()
	if err != nil {
		return err
	}

	defer xmlReader.Close()

	decoder := xml.NewDecoder(xmlReader)
	return decoder.Decode(target)
}

//MarshalZipFile add a file with content of marshaled source to zip
func MarshalZipFile(fileName string, source interface{}, to *zip.Writer) error {
	if prep, ok := source.(MarshalPreparation); ok {
		source = prep.BeforeMarshalXML()

		//if BeforeMarshalXML returns nil, then consider it like file must not be added to package
		if source == nil {
			return nil
		}
	}

	writer, err := to.Create(fileName)
	if err != nil {
		return err
	}

	content, err := xml.Marshal(source)

	if fix, ok := source.(MarshalFixation); ok {
		content = fix.AfterMarshalXML(content)
	}

	if err == nil {
		_, err = writer.Write([]byte(xml.Header + string(content)))
		if err != nil {
			return err
		}
	}

	return err
}

//CopyZipFile copies a zip file as is to a new zip
func CopyZipFile(from *zip.File, to *zip.Writer) error {
	writer, err := to.Create(from.Name)
	if err != nil {
		return err
	}

	reader, err := from.Open()
	if err != nil {
		return err
	}

	buff := bytes.NewBuffer(nil)
	_, err = io.Copy(buff, reader)
	if err != nil {
		return err
	}

	_, err = writer.Write(buff.Bytes())
	if err != nil {
		return err
	}

	return nil
}

//UniqueName returns next valid unique name with a valid length
func UniqueName(name string, names []string, nameLimit int) string {
	sanityChecked := 0
	regTitle := regexp.MustCompile(`[\d]+$`)

	for i := 1; sanityChecked < 2; i++ {
		sanityChecked++

		if len(name) > nameLimit {
			name = name[:nameLimit]
			sanityChecked = 0
		}

		for _, nextName := range names {
			//non unique name?
			if name == nextName {
				sanityChecked = 0
				suffix := strconv.Itoa(i)
				title := regTitle.ReplaceAllString(name, "")

				//next name is too large?
				if len(title+suffix) > nameLimit {
					title = title[:nameLimit-len(suffix)]
				}

				name = title + suffix
				break
			}
		}
	}

	return name
}
