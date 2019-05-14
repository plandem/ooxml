package vml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/vml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestVML(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
		<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel">
			<o:shapelayout v:ext="edit">
				<o:idmap v:ext="edit" data="1"/>
			</o:shapelayout>
			<v:shapetype id="_x0000_t202" coordsize="21600,21600" o:spt="202" path="m,l,21600r21600,l21600,xe" strokecolor="#81835a" o:insetmode="auto">
				<v:fill color="red"/>
				<v:shadow on="t" color="silver" opacity="1" obscured="true" />
				<v:stroke joinstyle="miter"/>
				<v:path gradientshapeok="t" o:connecttype="rect"/>
				<v:path o:connecttype="none"/>
			</v:shapetype>
			<v:shape id="_x0000_s1025" type="#_x0000_t202" style="position:absolute;margin-left:59.25pt;margin-top:1.5pt;width:96pt;height:55.5pt;z-index:1;visibility:hidden" o:insetmode="auto">
				<v:fill color2="#ffffe1"/>
				<v:shadow on="t" color="black" obscured="t"/>
				<v:path o:connecttype="none"/>
				<v:textbox style="mso-direction-alt:auto">
					<div style="text-align:left"/>
				</v:textbox>
				<x:ClientData ObjectType="Note">
					<x:MoveWithCells/>
					<x:SizeWithCells/>
					<x:Anchor>1, 15, 0, 2, 3, 15, 3, 16</x:Anchor>
					<x:AutoFill>False</x:AutoFill>
					<x:Row>0</x:Row>
					<x:Column>0</x:Column>
				</x:ClientData>
			</v:shape>
			<v:shape id="_x0000_s1026" type="#_x0000_t202" style="position:absolute;margin-left:59.25pt;margin-top:1.5pt;width:96pt;height:55.5pt;z-index:1;visibility:hidden">
				<x:ClientData ObjectType="Note">
					<x:MoveWithCells/>
					<x:SizeWithCells/>
					<x:Anchor>1, 15, 0, 2, 3, 15, 3, 16</x:Anchor>
					<x:AutoFill>False</x:AutoFill>
					<x:Row>0</x:Row>
					<x:Column>1</x:Column>
				</x:ClientData>
			</v:shape>
		</xml>
	`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &vml.Excel{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, "_x0000_t202", entity.ShapeType[0].ID())
	require.Equal(t, "#81835a", entity.ShapeType[0].Attr("strokecolor"))
	require.Equal(t, "", entity.ShapeType[0].Type())
	require.Equal(t, entity.ShapeType[0].ID(), entity.ShapeType[0].Attr("id"))
	require.Equal(t, entity.ShapeType[0].Type(), entity.ShapeType[0].Attr("type"))

	require.Equal(t, "_x0000_s1025", entity.Shape[0].ID())
	require.Equal(t, "#_x0000_t202", entity.Shape[0].Type())
	require.Equal(t, "auto", entity.Shape[0].Attr("o:insetmode"))
	require.Equal(t, entity.Shape[0].ID(), entity.Shape[0].Attr("id"))
	require.Equal(t, entity.Shape[0].Type(), entity.Shape[0].Attr("type"))

	require.Equal(t, "_x0000_s1026", entity.Shape[1].ID())
	require.Equal(t, "#_x0000_t202", entity.Shape[1].Type())
	require.Equal(t, entity.Shape[1].ID(), entity.Shape[1].Attr("id"))
	require.Equal(t, entity.Shape[1].Type(), entity.Shape[1].Attr("type"))
	require.Equal(t, data, string(encoded))
}