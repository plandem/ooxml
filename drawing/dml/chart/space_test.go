// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package chart_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/drawing/dml/chart"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSpace(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "").Replace(`
<c:chartSpace xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main">
	<c:lang val="en-US"></c:lang>
	<c:date1904 val="false"></c:date1904>
	<c:roundedCorners val="false"></c:roundedCorners>
	<c:chart>
		<c:autoTitleDeleted val="false"></c:autoTitleDeleted>
	</c:chart>
	<c:spPr>
		<a:effectLst></a:effectLst>
	</c:spPr>
	<c:txPr>
		<a:bodyPr></a:bodyPr>
	</c:txPr>
	<c:printSettings>
		<c:headerFooter></c:headerFooter>
		<c:pageMargins b="0.75" l="0.7" r="0.7" t="0.75" header="0.3" footer="0.3"></c:pageMargins>
		<c:pageSetup></c:pageSetup>
	</c:printSettings>
</c:chartSpace>
`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &chart.Space{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	falseProperty := ml.PropertyBool(false)

	object := &chart.Space{
		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/drawingml/2006/chart",
			Local: "chartSpace",
		},
		Lang:           "en-US",
		Date1904:       &falseProperty,
		RoundedCorners: &falseProperty,
		Chart: &chart.Chart{
			AutoTitleIsDeleted: &falseProperty,
		},
		Shape: &dml.Shape{
			ReservedElements: ml.ReservedElements{
				Nodes: []ml.Reserved{
					{
						XMLName: xml.Name{
							Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
							Local: "effectLst",
						},
					},
				},
			},
		},
		TextBody: &dml.TextBody{
			ReservedElements: ml.ReservedElements{
				Nodes: []ml.Reserved{
					{
						XMLName: xml.Name{
							Space: "http://schemas.openxmlformats.org/drawingml/2006/main",
							Local: "bodyPr",
						},
					},
				},
			},
		},
		//ReservedAttributes: ml.ReservedAttributes{
		//	Attrs: []xml.Attr{
		//		{ Name: xml.Name{Space: "xmlns", Local:"c"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/chart"},
		//		{ Name: xml.Name{Space: "xmlns", Local:"a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"},
		//	},
		//	},
		ReservedElements: ml.ReservedElements{
			Nodes: []ml.Reserved{
				{
					XMLName: xml.Name{
						Space: "http://schemas.openxmlformats.org/drawingml/2006/chart",
						Local: "printSettings",
					},
					InnerXML: `<c:headerFooter></c:headerFooter><c:pageMargins b="0.75" l="0.7" r="0.7" t="0.75" header="0.3" footer="0.3"></c:pageMargins><c:pageSetup></c:pageSetup>`,
				},
			},
		},
	}

	require.Equal(t, object, entity)

	//encode data should be same as original
	encode, err := xml.Marshal(entity)
	require.Nil(t, err)

	//we can't compare directly, because 1) reserved elements should keep namespace and 2) we also add few more namespaces
	data = strings.NewReplacer("\t", "", "\n", "").Replace(`
<chartSpace xmlns="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<lang val="en-US"></lang>
	<date1904 val="false"></date1904>
	<roundedCorners val="false"></roundedCorners>
	<chart>
		<autoTitleDeleted val="false"></autoTitleDeleted>
	</chart>
	<spPr>
		<a:effectLst></a:effectLst>
	</spPr>
	<txPr>
		<a:bodyPr></a:bodyPr>
	</txPr>
	<c:printSettings>
		<c:headerFooter></c:headerFooter>
		<c:pageMargins b="0.75" l="0.7" r="0.7" t="0.75" header="0.3" footer="0.3"></c:pageMargins>
		<c:pageSetup></c:pageSetup>
	</c:printSettings>
</chartSpace>
`)
	require.Equal(t, strings.NewReplacer("\t", "", "\n", "").Replace(data), string(encode))
}
