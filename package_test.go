// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ooxml_test

import (
	"github.com/plandem/ooxml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestOpening(t *testing.T) {
	type OOXmlDoc struct {
		*ooxml.PackageInfo
	}

	factory := func(pkg *ooxml.PackageInfo) (interface{}, error) {
		return &OOXmlDoc{pkg}, nil
	}

	//can't open
	doc, err := ooxml.Open("./test_files/unknown_file.xlsx", factory)
	assert.Nil(t, doc)
	assert.NotNil(t, err)

	//non zip
	doc, err = ooxml.Open("./package.go", factory)
	assert.Nil(t, doc)
	assert.NotNil(t, err)

	//ok for fileName
	doc, err = ooxml.Open("./test_files/example_simple.xlsx", factory)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
	assert.IsType(t, &OOXmlDoc{}, doc)
	assert.Implements(t, (*ooxml.Package)(nil), doc)

	//non zip
	nonZipFile, err := os.Open("./package.go")
	assert.Nil(t, err)

	doc, err = ooxml.Open(nonZipFile, factory)
	assert.NotNil(t, err)
	assert.Nil(t, doc)

	//zip file
	zipFile, err := os.Open("./test_files/example_simple.xlsx")
	assert.NotNil(t, zipFile)
	assert.Nil(t, err)

	//ok for zipReader
	doc, err = ooxml.Open(zipFile, factory)
	assert.NotNil(t, doc)
	assert.Nil(t, err)
	assert.IsType(t, &OOXmlDoc{}, doc)
	assert.Implements(t, (*ooxml.Package)(nil), doc)
}

func TestClose(t *testing.T) {
	type OOXmlDoc struct {
		*ooxml.PackageInfo
	}

	factory := func(pkg *ooxml.PackageInfo) (interface{}, error) {
		return &OOXmlDoc{pkg}, nil
	}

	//ok for zip.ReadCloser
	doc, _ := ooxml.Open("./test_files/example_simple.xlsx", factory)
	require.NotNil(t, doc)
	require.Implements(t, (*ooxml.Package)(nil), doc)

	odoc, _ := doc.(*OOXmlDoc)
	odoc.Close()

	//zip file
	zipFile, _ := os.Open("./test_files/example_simple.xlsx")
	require.NotNil(t, zipFile)

	//ok for zip.Reader
	doc, _ = ooxml.Open(zipFile, factory)
	require.NotNil(t, doc)
	require.Implements(t, (*ooxml.Package)(nil), doc)

	odoc, _ = doc.(*OOXmlDoc)
	odoc.Close()
}

func TestSaveAs(t *testing.T) {
	type OOXmlDoc struct {
		*ooxml.PackageInfo
	}

	factory := func(pkg *ooxml.PackageInfo) (interface{}, error) {
		return &OOXmlDoc{pkg}, nil
	}

	doc, _ := ooxml.Open("./test_files/example_simple.xlsx", factory)
	require.NotNil(t, doc)
	require.Implements(t, (*ooxml.Package)(nil), doc)

	odoc, _ := doc.(*OOXmlDoc)
	err := odoc.SaveAs("./test_files/example_simple_saved.xlsx")
	assert.Nil(t, err)

	err = odoc.Close()
	assert.Nil(t, err)
}
