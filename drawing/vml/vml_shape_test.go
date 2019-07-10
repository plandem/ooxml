// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestShape(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<v:shape path="m1285,251l1126,469,580,1009,,1285,25,1412,93,1547,194,1673,1017,2026,2312,2060,3209,1756,3765,1388,3278,680,3059,319,2976,,1285,251,1285,251xe" 
		style="position:relative;left:234.75pt;top:208.875pt;width:235.25pt;height:128.875pt" 
		id="_x0000_s1025" 
		coordsize="3765,2060" 
		fillcolor="#bcbcd6" stroked="false"
	>
		<v:path arrowok="true"></v:path>
		<x:ClientData ObjectType="Note">
			<x:MoveWithCells>true</x:MoveWithCells>
			<x:SizeWithCells>true</x:SizeWithCells>
			<x:AutoFill>false</x:AutoFill>
			<x:Row>1</x:Row>
			<x:Column>2</x:Column>
			<x:Anchor>1, 15, 0, 2, 3, 15, 3, 16</x:Anchor>
		</x:ClientData>
	</v:shape>
</xml>
	`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Excel{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	//check decoded shape
	require.Equal(t, "_x0000_s1025", entity.Shape[0].ID)
	require.Equal(t, &css.Style{
		Position: css.PositionRelative,
		Left:     css.NewNumber(234.75),
		Top:      css.NewNumber(208.875),
		Width:    css.NewNumber(235.25),
		Height:   css.NewNumber(128.875),
	}, css.NewStyle(entity.Shape[0].Style))
	require.Equal(t, "3765,2060", entity.Shape[0].CoordSize)
	require.Equal(t, "m1285,251l1126,469,580,1009,,1285,25,1412,93,1547,194,1673,1017,2026,2312,2060,3209,1756,3765,1388,3278,680,3059,319,2976,,1285,251,1285,251xe", entity.Shape[0].Path)
	require.Equal(t, "#bcbcd6", entity.Shape[0].FillColor)
	require.Equal(t, ml.TriStateFalse, entity.Shape[0].Stroked)
	require.Equal(t, ml.TriStateTrue, entity.Shape[0].PathSettings.ArrowOK)

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
	}, entity.Shape[0].ClientData)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
