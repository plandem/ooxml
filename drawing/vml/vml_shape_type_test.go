package vml

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestShapeType(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<v:shapetype o:spt="202" path="m,l,21600r21600,l21600,xe" id="_x0000_t202" coordsize="21600,21600">
		<v:path o:connecttype="rect" gradientshapeok="true"></v:path>
		<v:stroke joinstyle="miter"></v:stroke>
	</v:shapetype>
</xml>
	`)

	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Excel{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	//check decoded shape type
	require.Equal(t, "_x0000_t202", entity.ShapeType[0].ID)
	require.Equal(t, "21600,21600", entity.ShapeType[0].CoordSize)
	require.Equal(t, 202, entity.ShapeType[0].Spt)
	require.Equal(t, "m,l,21600r21600,l21600,xe", entity.ShapeType[0].Path)
	require.Equal(t, StrokeJoinStyleMiter, entity.ShapeType[0].Stroke.JoinStyle)
	require.Equal(t, ml.TriStateTrue, entity.ShapeType[0].PathSettings.GradientShapeOK)
	require.Equal(t, ConnectTypeRect, entity.ShapeType[0].PathSettings.ConnectType)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
