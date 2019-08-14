// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

type Chart1 struct {
	RID ml.RID `xml:"id,attr"`
}

type Entity1 struct {
	XMLName xml.Name
	RIDName ml.RIDName `xml:",attr"`
	Chart   *Chart1    `xml:"chart,omitempty"`
}

func (c *Chart1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = ml.ApplyNamespacePrefix(ml.NamespaceDrawingChart, start.Name)
	start.Attr = append(start.Attr, ml.Namespaces(
		ml.NamespaceDrawingChart,
	)...)

	return e.EncodeElement(*c, start)
}

func TestRelationships_marshalCustom(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<entity xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<c:chart xmlns:c="http://schemas.openxmlformats.org/drawingml/2006/chart" r:id="rId1"></c:chart>
</entity>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity1{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, &Chart1{
		RID: "rId1",
	}, entity.Chart)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}

func TestRelationships_marshalDefault(t *testing.T) {
	type Chart2 struct {
		RID ml.RID `xml:"id,attr"`
	}

	type Entity2 struct {
		XMLName xml.Name
		RIDName ml.RIDName `xml:",attr"`
		Chart   *Chart2    `xml:"chart,omitempty"`
	}

	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<entity xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<chart r:id="rId1"></chart>
</entity>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity2{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, &Chart2{
		RID: "rId1",
	}, entity.Chart)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
