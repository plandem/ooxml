package vml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestOval(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<v:oval style="top:1;left:1;width:50;height:50" fillcolor="green"></v:oval>
</xml>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &vml.Excel{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, "oval", entity.Oval[0].XMLName.Local)
	require.Equal(t, "green", entity.Oval[0].FillColor)
	require.Equal(t, "top:1;left:1;width:50;height:50", entity.Oval[0].Style)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
