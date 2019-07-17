// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<v:shape path="m 1,1 l 1,200, 200,200, 200,1 x e" style="position:relative;top:1;left:1;width:400;height:400">
		<v:handles>
			<v:h position="#0,#1" xrange="0,21600" yrange="0,10800"></v:h>
		</v:handles>
	</v:shape>
</xml>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &vml.Excel{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, &vml.Handles{
		XMLName: xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "handles"},
		List: []vml.Handle{
			{
				Position: "#0,#1",
				XRange:   "0,21600",
				YRange:   "0,10800",
			},
		},
	}, entity.Shape[0].Handles)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
