package ooxml_test

import (
	"github.com/plandem/ooxml"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"strings"
	"testing"
)

func TestUnzip(t *testing.T) {
	dest, err := ioutil.TempDir("", "example_simple.xlsx-")
	require.Nil(t, err)
	require.NotNil(t, dest)
	require.Equal(t, false, strings.HasSuffix(dest, "/"))

	files, err := ooxml.Unzip("./test_files/example_simple.xlsx", dest)
	require.Nil(t, err)
	require.NotNil(t, files)
	require.Equal(t, []string {
		"[Content_Types].xml",
		"_rels/.rels",
		"xl/_rels/workbook.xml.rels",
		"xl/workbook.xml",
		"xl/sharedStrings.xml",
		"xl/theme/theme1.xml",
		"xl/styles.xml",
		"xl/worksheets/sheet1.xml",
		"docProps/thumbnail.jpeg",
		"docProps/core.xml",
		"docProps/app.xml",
	}, files)
}
