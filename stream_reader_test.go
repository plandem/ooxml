package ooxml_test

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
	"regexp"
)

func ExampleStreamReader() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}

	type Address struct {
		City, State string
	}

	type Result struct {
		XMLName xml.Name `xml:"personNamespace Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}

	data := `
  		<Person>
  			<FullName>Grace R. Emlin</FullName>
  			<Company>Example Inc.</Company>
  			<Email where="home">
  				<Addr>gre@example.com</Addr>
  			</Email>
  			<Email where='work'>
  				<Addr>gre@work.com</Addr>
  			</Email>
  			<Group>
  				<Value>Friends</Value>
  				<Value>Squash</Value>
  			</Group>
  			<City>Hanga Roa</City>
  			<State>Easter Island</State>
  		</Person>
  	`

	rs := ooxml.StreamReader{Decoder: xml.NewDecoder(bytes.NewReader([]byte(data)))}

	for next, hasNext := rs.StartIterator(nil); hasNext; {
		hasNext = next(func(decoder *xml.Decoder, start *xml.StartElement) bool {
			fmt.Println(start.Name.Local)
			return true
		})
	}

	//Output:
	// Person
	// FullName
	// Company
	// Email
	// Addr
	// Email
	// Addr
	// Group
	// Value
	// Value
	// City
	// State
}

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

	files := odoc.PackageInfo.Files()
	for _, file := range files {
		if f, ok := file.(*zip.File); ok {
			if reSheet.MatchString(f.Name) {
				sheetStream, _ = ooxml.NewStreamFileReader(f)
				break
			}
		}
	}

	var rowIterator ooxml.StreamReaderIterator
	var hasNextRow bool
	for next, hasNext := sheetStream.StartIterator(nil); hasNext; {
		hasNext = next(func(decoder *xml.Decoder, start *xml.StartElement) bool {
			switch start.Name.Local {
			case "row":
				rowIterator, hasNextRow = sheetStream.StartIterator(start)
				return false
			}

			return true
		})
	}

	for hasNextRow {
		hasNextRow = rowIterator(func(decoder *xml.Decoder, start *xml.StartElement) bool {
			if start != nil && start.Name.Local == "row" {
				var row Row
				if decoder.DecodeElement(&row, start) == nil {
					for _, c := range row.Cells {
						if c.Type == "s"  || c.Value == "" {
							continue
						}

						fmt.Printf("%+v\n", c.Value)
					}

					return true
				}
			}

			return false
		})
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
