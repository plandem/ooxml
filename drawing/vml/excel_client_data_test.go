package vml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestClientData(t *testing.T) {
	type Entity struct {
		XMLName    xml.Name       `xml:"xml"`
		ClientData vml.ClientData `xml:"ClientData"`
	}
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:x="urn:schemas-microsoft-com:office:excel">
	<x:ClientData ObjectType="Note">
		<x:MoveWithCells>false</x:MoveWithCells>
		<x:SizeWithCells/>
		<x:Anchor>3, 15, 1, 10, 5, 15, 2, 64</x:Anchor>
		<x:AutoFill>False</x:AutoFill>
		<x:Row>1</x:Row>
		<x:Column>2</x:Column>
		<x:Visible/>
		<x:Help/>
	</x:ClientData>
</xml>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, vml.ClientData{
		XMLName:       xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ClientData"},
		Type:          vml.ObjectTypeNote,
		MoveWithCells: ml.TriStateBlankTrue(ml.TriStateFalse),
		SizeWithCells: ml.TriStateBlankTrue(ml.TriStateTrue),
		AutoFill:      ml.TriStateBlankTrue(ml.TriStateFalse),
		Anchor:        "3, 15, 1, 10, 5, 15, 2, 64",
		Row:           1,
		Column:        2,
		Visible:       ml.TriStateBlankTrue(ml.TriStateTrue),
		ReservedElements: ml.ReservedElements{
			Nodes: []ml.Reserved{
				{
					XMLName: xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Help"},
				},
			},
		},
	}, entity.ClientData)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)
	require.Equal(t, strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml>
	<x:ClientData ObjectType="Note">
		<x:MoveWithCells>false</x:MoveWithCells>
		<x:SizeWithCells>true</x:SizeWithCells>
		<x:AutoFill>false</x:AutoFill>
		<x:Visible>true</x:Visible>
		<x:Row>1</x:Row>
		<x:Column>2</x:Column>
		<x:Anchor>3, 15, 1, 10, 5, 15, 2, 64</x:Anchor>
		<x:Help></x:Help>
	</x:ClientData>
</xml>
`), string(encoded))
}
