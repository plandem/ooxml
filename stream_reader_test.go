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
)

func ExampleStreamFileReader() {
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
		Ref          int          `xml:"r,attr,omitempty"`
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

	openSheet := func(cb func(f *zip.File)) {
		z, err := zip.OpenReader(`./test_files/example_simple.xlsx`)
		if err != nil {
			panic(err)
		}

		defer z.Close()

		for _, f := range z.File {
			if f.Name == `xl/worksheets/sheet1.xml` {
				cb(f)
				break
			}
		}
	}

	openSheet(func(f *zip.File) {
		stream, _ := ooxml.NewStreamFileReader(f)

		for {
			t, _ := stream.Token()
			if t == nil {
				break
			}

			switch start := t.(type) {
			case xml.StartElement:
				if start.Name.Local == "row" {
					var row Row
					if stream.DecodeElement(&row, &start) == nil {
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

		_ = stream.Close()
	})

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