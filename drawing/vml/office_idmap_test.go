package vml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestIdMap(t *testing.T) {
	type Entity struct {
		XMLName xml.Name
		IdMap   *vml.IdMap `xml:"idmap"`
	}

	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<entity>
	<o:idmap v:ext="edit" data="1"></o:idmap>
</entity>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, "1", entity.IdMap.Data)
	require.Equal(t, vml.ExtTypeEdit, entity.IdMap.Ext)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
