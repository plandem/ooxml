package ml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestReserved(t *testing.T) {
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
