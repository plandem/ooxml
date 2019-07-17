// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ooxml_test

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
	"regexp"
)

func ExampleStreamFileReader() {
	type OOXmlDoc struct {
		*ooxml.PackageInfo
	}

	//Cell is a direct mapping of XSD CT_Cell
	type Cell struct {
		Formula   *ml.Reserved     `xml:"f,omitempty"`
		Value     string           `xml:"v,omitempty"`
		InlineStr *ml.Reserved     `xml:"is,omitempty"`
		ExtLst    *ml.Reserved     `xml:"extLst,omitempty"`
		Ref       string           `xml:"r,attr"`
		Style     int              `xml:"s,attr,omitempty"`
		Type      string           `xml:"t,attr,omitempty"`
		Cm        ml.OptionalIndex `xml:"cm,attr,omitempty"`
		Vm        ml.OptionalIndex `xml:"vm,attr,omitempty"`
		Ph        bool             `xml:"ph,attr,omitempty"`
	}

	type Row struct {
		Cells        []*Cell      `xml:"c"`
		ExtLst       *ml.Reserved `xml:"extLst,omitempty"`
		Ref          int          `xml:"r,attr,omitempty"` //1-based index
		Spans        string       `xml:"spans,attr,omitempty"`
		Style        int          `xml:"s,attr,omitempty"`
		CustomFormat bool         `xml:"customFormat,attr,omitempty"`
		Height       float32      `xml:"ht,attr,omitempty"`
		Hidden       bool         `xml:"hidden,attr,omitempty"`
		CustomHeight bool         `xml:"customHeight,attr,omitempty"`
		OutlineLevel uint8        `xml:"outlineLevel,attr,omitempty"`
		Collapsed    bool         `xml:"collapsed,attr,omitempty"`
		ThickTop     bool         `xml:"thickTop,attr,omitempty"`
		ThickBot     bool         `xml:"thickBot,attr,omitempty"`
		Phonetic     bool         `xml:"ph,attr,omitempty"`
	}

	factory := func(pkg *ooxml.PackageInfo) (interface{}, error) {
		return &OOXmlDoc{pkg}, nil
	}

	//ok for zip.ReadCloser
	doc, _ := ooxml.Open("./test_files/example_simple.xlsx", factory)
	odoc, _ := doc.(*OOXmlDoc)
	defer odoc.Close()

	var sheetStream *ooxml.StreamFileReader

	reSheet := regexp.MustCompile(`xl/worksheets/[[:alpha:]]+[\d]+\.xml`)

	files := odoc.PackageInfo.Files(nil)
	for _, file := range files {
		if f, ok := file.(*zip.File); ok {
			if reSheet.MatchString(f.Name) {
				sheetStream, _ = ooxml.NewStreamFileReader(f)
				break
			}
		}
	}

	if sheetStream != nil {
		for {
			// Read tokens from the XML document in a stream.
			t, _ := sheetStream.Token()
			if t == nil {
				break
			}

			switch se := t.(type) {
			case xml.StartElement:
				if se.Name.Local == "row" {
					var row Row
					if sheetStream.DecodeElement(&row, &se) == nil {
						for _, c := range row.Cells {
							if c.Type == "s" || c.Value == "" {
								continue
							}

							fmt.Printf("%+v\n", c.Value)
						}
					}
				}
			}
		}

		_ = sheetStream.Close()
	}

	//Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
	// 12
	// 13
	// 14
	// 15
	// 16
	// 17
	// 18
	// 19
	// 20
}
