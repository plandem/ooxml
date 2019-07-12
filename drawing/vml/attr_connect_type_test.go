// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnectType(t *testing.T) {
	type Entity struct {
		Attribute vml.ConnectType `xml:"attribute,attr,omitempty"`
	}

	list := map[vml.ConnectType]string{
		vml.ConnectType(0):      "",
		vml.ConnectTypeNone:     vml.ConnectTypeNone.String(),
		vml.ConnectTypeRect:     vml.ConnectTypeRect.String(),
		vml.ConnectTypeSegments: vml.ConnectTypeSegments.String(),
		vml.ConnectTypeCustom:   vml.ConnectTypeCustom.String(),
	}

	for k, v := range list {
		t.Run(v, func(tt *testing.T) {
			entity := Entity{Attribute: k}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if k == 0 {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity o:attribute="%s"></Entity>`, v), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, v, decoded.Attribute.String())
		})
	}
}
