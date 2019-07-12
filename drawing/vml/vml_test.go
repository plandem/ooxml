// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"bytes"
	"encoding/xml"
	"github.com/go-test/deep"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestExcel_MarshalXML(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"></xml>
`)

	entity := &Excel{}
	encoded, err := xml.Marshal(entity)
	require.Nil(t, err)
	require.Equal(t, data, string(encoded))
}

func TestWord_MarshalXML(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:w="urn:schemas-microsoft-com:office:word" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"></xml>
`)

	encoded, err := xml.Marshal(&Word{})
	require.Nil(t, err)
	require.Equal(t, data, string(encoded))
}

func TestPowerPoint_MarshalXML(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:p="urn:schemas-microsoft-com:office:powerpoint" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"></xml>
`)

	encoded, err := xml.Marshal(&PowerPoint{})
	require.Nil(t, err)
	require.Equal(t, data, string(encoded))
}

func TestVML(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<o:shapelayout v:ext="edit">
		<o:idmap v:ext="edit" data="1"/>
	</o:shapelayout>
	<v:shapetype id="_x0000_t202" coordsize="21600,21600" o:spt="202" path="m,l,21600r21600,l21600,xe">
		<v:stroke joinstyle="miter"/>
		<v:path gradientshapeok="t" o:connecttype="rect"/>
	</v:shapetype>
	<v:shape id="_x0000_s1025"
		style='position:relative;left:234.75pt;top:208.875pt;width:235.25pt;height:128.875pt'
		coordsize="3765,2060"
		path="m1285,251l1126,469,580,1009,,1285,25,1412,93,1547,194,1673,1017,2026,2312,2060,3209,1756,3765,1388,3278,680,3059,319,2976,,1285,251,1285,251xe"
		fillcolor="#bcbcd6" stroked="f"
	>
		<v:path arrowok="t"/>
		<x:ClientData ObjectType="Note">
			<x:MoveWithCells/>
			<x:SizeWithCells/>
			<x:Anchor>1, 15, 0, 2, 3, 15, 3, 16</x:Anchor>
			<x:AutoFill>False</x:AutoFill>
			<x:Row>1</x:Row>
			<x:Column>2</x:Column>
		</x:ClientData>
	</v:shape>
</xml>
	`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	excel := &Excel{}
	err := decoder.DecodeElement(excel, nil)
	require.Nil(t, err)

	//check decoded shape layout
	require.Equal(t, ExtTypeEdit, excel.ShapeLayout.Ext)
	require.Equal(t, ExtTypeEdit, excel.ShapeLayout.IdMap.Ext)
	require.Equal(t, "1", excel.ShapeLayout.IdMap.Data)

	//check decoded shape type
	require.Equal(t, "_x0000_t202", excel.ShapeType[0].ID)
	require.Equal(t, "21600,21600", excel.ShapeType[0].CoordSize)
	require.Equal(t, 202, excel.ShapeType[0].Spt)
	require.Equal(t, "m,l,21600r21600,l21600,xe", excel.ShapeType[0].Path)
	require.Equal(t, StrokeJoinStyleMiter, excel.ShapeType[0].Stroke.JoinStyle)
	require.Equal(t, ml.TriStateTrue, excel.ShapeType[0].PathSettings.GradientShapeOK)
	require.Equal(t, ConnectTypeRect, excel.ShapeType[0].PathSettings.ConnectType)

	//check decoded shape
	require.Equal(t, "_x0000_s1025", excel.Shape[0].ID)
	require.Equal(t, &css.Style{
		Position: css.PositionRelative,
		Left:     css.NewNumber(234.75),
		Top:      css.NewNumber(208.875),
		Width:    css.NewNumber(235.25),
		Height:   css.NewNumber(128.875),
	}, css.NewStyle(excel.Shape[0].Style))
	require.Equal(t, "3765,2060", excel.Shape[0].CoordSize)
	require.Equal(t, "m1285,251l1126,469,580,1009,,1285,25,1412,93,1547,194,1673,1017,2026,2312,2060,3209,1756,3765,1388,3278,680,3059,319,2976,,1285,251,1285,251xe", excel.Shape[0].Path)
	require.Equal(t, "#bcbcd6", excel.Shape[0].FillColor)
	require.Equal(t, ml.TriStateFalse, excel.Shape[0].Stroked)
	require.Equal(t, ml.TriStateTrue, excel.Shape[0].PathSettings.ArrowOK)

	//check ClientData
	require.Equal(t, &ClientData{
		XMLName:       xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ClientData"},
		Type:          ObjectTypeNote,
		MoveWithCells: ml.TriStateBlankTrue(ml.TriStateTrue),
		SizeWithCells: ml.TriStateBlankTrue(ml.TriStateTrue),
		AutoFill:      ml.TriStateBlankTrue(ml.TriStateFalse),
		Anchor:        "1, 15, 0, 2, 3, 15, 3, 16",
		Row:           1,
		Column:        2,
	}, excel.Shape[0].ClientData)

	//check encode -> decode -> original
	encoded, err := xml.MarshalIndent(excel, "", " ")
	require.Nil(t, err)

	excel2 := &Excel{}
	decoder = xml.NewDecoder(bytes.NewReader([]byte(encoded)))
	err = decoder.DecodeElement(excel2, nil)
	require.Nil(t, err)

	//we need to encode excel2 also to resolve namespaces
	_, err = xml.Marshal(excel2)
	require.Nil(t, err)

	//both entity should be same now
	if diff := deep.Equal(excel, excel2); diff != nil {
		t.Error(diff)
	}
}
