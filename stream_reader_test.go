package ooxml_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml"
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
