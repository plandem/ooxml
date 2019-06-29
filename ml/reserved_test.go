package ml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestReserved1(t *testing.T) {
	type Entity struct {
		XMLName xml.Name
		Root    *ml.Reserved `xml:"ext"`
	}

	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<entity>
	<ext id="1" title="text title">
		<sub>content</sub>
	</ext>
</entity>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, xml.Name{
		Local: "ext",
	}, entity.Root.XMLName)

	require.Equal(t, []xml.Attr{
		{
			Name:  xml.Name{Local: "id"},
			Value: "1",
		},
		{
			Name:  xml.Name{Local: "title"},
			Value: "text title",
		},
	}, entity.Root.Attrs)
	require.Equal(t, `<sub>content</sub>`, entity.Root.InnerXML)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}

func TestReservedAttributes(t *testing.T) {
	type Entity struct {
		XMLName xml.Name
		ID      int `xml:"id,attr,omitempty"`
		ml.ReservedAttributes
	}

	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<entity id="123" title="text title"></entity>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, 123, entity.ID)

	require.Equal(t, []xml.Attr{
		{
			Name:  xml.Name{Local: "title"},
			Value: "text title",
		},
	}, entity.Attrs)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}

func TestReservedElements(t *testing.T) {
	type SubEntity struct {
		XMLName xml.Name
		Value   string `xml:"value,attr,omitempty"`
		ml.ReservedAttributes
	}

	type Entity struct {
		XMLName xml.Name
		ID      int        `xml:"id,attr,omitempty"`
		Sub     *SubEntity `xml:"sub,omitempty"`
		ml.ReservedElements
		ml.ReservedAttributes
	}

	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<entity id="123" title="text title">
	<sub value="val" mode="x"></sub>
	<ext guid="5"><another/></ext>
</entity>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	//check entity
	require.Equal(t, 123, entity.ID)

	//check reserved attributes of entity
	require.Equal(t, []xml.Attr{
		{
			Name:  xml.Name{Local: "title"},
			Value: "text title",
		},
	}, entity.Attrs)

	//check sub
	require.Equal(t, xml.Name{
		Local: "sub",
	}, entity.Sub.XMLName)

	require.Equal(t, "val", entity.Sub.Value)

	//check reserved attributes of sub
	require.Equal(t, []xml.Attr{
		{
			Name:  xml.Name{Local: "mode"},
			Value: "x",
		},
	}, entity.Sub.Attrs)

	//check reserved elements of entity
	require.Equal(t, 1, len(entity.Nodes))
	require.Equal(t, xml.Name{Local: "ext"}, entity.Nodes[0].XMLName)
	require.Equal(t, `<another/>`, entity.Nodes[0].InnerXML)
	require.Equal(t, []xml.Attr{
		{
			Name:  xml.Name{Local: "guid"},
			Value: "5",
		},
	}, entity.Nodes[0].Attrs)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
