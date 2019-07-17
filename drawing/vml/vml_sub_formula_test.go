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

func TestFormula(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<xml xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
	<v:shape path="m 1,1 l 1,200, 200,200, 200,1 x e" style="position:relative;top:1;left:1;width:400;height:400">
		<v:formulas>
			<v:f eqn="sum 33030 0 #0"></v:f>
			<v:f eqn="prod #0 4 3"></v:f>
			<v:f eqn="prod @0 1 3"></v:f>
			<v:f eqn="sum @1 0 @2"></v:f>
		</v:formulas>
	</v:shape>
</xml>
`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &vml.Excel{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, &vml.Formulas{
		XMLName: xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "formulas"},
		List: []vml.Formula{
			"sum 33030 0 #0",
			"prod #0 4 3",
			"prod @0 1 3",
			"sum @1 0 @2",
		},
	}, entity.Shape[0].Formulas)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, data, string(encoded))
}
