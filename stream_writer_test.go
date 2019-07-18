// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ooxml_test

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStreamFileWriter(t *testing.T) {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}

	type Address struct {
		City, State string
	}

	type Person struct {
		Name   string `xml:"FullName"`
		Phone  string
		Email  []Email
		Groups []string `xml:"Group>Value"`
		Address
	}

	fileName := `xl/worksheets/sheet1.xml`

	tests := []struct {
		name   string
		memory bool
	}{
		{"tempFile", false},
		{"memory", true},
	}

	streamClosed := false
	for _, info := range tests {
		t.Run(info.name, func(tt *testing.T) {
			sheetStream, err := ooxml.NewStreamFileWriter(fileName, info.memory, func() error {
				streamClosed = true
				return nil
			})
			require.IsType(tt, &ooxml.StreamFileWriter{}, sheetStream)
			require.Nil(tt, err)

			persons := []*Person{
				{
					Name: "John Doe",
					Email: []Email{
						{Where: "home", Addr: "john@example.com"},
						{Where: "work", Addr: "john@work.com"},
					},
					Groups: []string{"husband", "father"},
					Address: Address{
						City:  "Denver",
						State: "Colorado",
					},
				},
				{
					Name: "Jane Dow",
					Email: []Email{
						{Where: "home", Addr: "jane@example.com"},
						{Where: "work", Addr: "jane@work.com"},
					},
					Groups: []string{"wife", "mother"},
					Address: Address{
						City:  "Washington",
						State: "Washington",
					},
				},
			}

			err = sheetStream.Encode(persons[0])
			require.Nil(tt, err)

			err = sheetStream.Encode(persons[1])
			require.Nil(tt, err)

			//prepare target zip
			buf := bytes.NewBuffer(nil)
			zipper := zip.NewWriter(buf)

			//add stream file into the target
			err = sheetStream.Save(zipper)
			require.Equal(tt, true, streamClosed)
			require.Nil(tt, err)

			err = zipper.Close()
			require.Nil(tt, err)

			//open file
			readerAt := bytes.NewReader(buf.Bytes())
			reader, err := zip.NewReader(readerAt, int64(readerAt.Len()))
			require.Nil(tt, err)

			//file has same name
			require.Equal(tt, fileName, reader.File[0].Name)
			zipFile, err := reader.File[0].Open()
			require.Nil(tt, err)

			//content is same
			decoder := xml.NewDecoder(zipFile)
			result := make([]*Person, 0)
			err = decoder.Decode(&result)
			require.Nil(tt, err)
			err = decoder.Decode(&result)
			require.Nil(tt, err)
			require.Equal(tt, persons, result)
			err = zipFile.Close()
			require.Nil(tt, err)
		})
	}
}
