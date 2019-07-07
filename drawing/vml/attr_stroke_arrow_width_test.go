package vml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStrokeArrowWidth(t *testing.T) {
	type Entity struct {
		Attribute vml.StrokeArrowWidth `xml:"attribute,attr,omitempty"`
	}

	list := map[vml.StrokeArrowWidth]string{
		vml.StrokeArrowWidth(0):    "",
		vml.StrokeArrowWidthMedium: vml.StrokeArrowWidthMedium.String(),
		vml.StrokeArrowWidthNarrow: vml.StrokeArrowWidthNarrow.String(),
		vml.StrokeArrowWidthWide:   vml.StrokeArrowWidthWide.String(),
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
