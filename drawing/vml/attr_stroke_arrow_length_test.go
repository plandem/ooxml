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

func TestStrokeArrowLength(t *testing.T) {
	type Entity struct {
		Attribute vml.StrokeArrowLength `xml:"attribute,attr,omitempty"`
	}

	list := map[vml.StrokeArrowLength]string{
		vml.StrokeArrowLength(0):    "",
		vml.StrokeArrowLengthLong:   vml.StrokeArrowLengthLong.String(),
		vml.StrokeArrowLengthMedium: vml.StrokeArrowLengthMedium.String(),
		vml.StrokeArrowLengthShort:  vml.StrokeArrowLengthShort.String(),
	}

	for k, v := range list {
		t.Run(v, func(tt *testing.T) {
			entity := Entity{Attribute: k}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			if k == 0 {
				require.Equal(tt, `<Entity></Entity>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, v), string(encoded))
			}

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, v, decoded.Attribute.String())
		})
	}
}
