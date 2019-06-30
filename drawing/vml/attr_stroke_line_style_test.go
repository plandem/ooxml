package vml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStrokeLineStyle(t *testing.T) {
	type Entity struct {
		Attribute vml.StrokeLineStyle `xml:"attribute,attr,omitempty"`
	}

	list := map[vml.StrokeLineStyle]string{
		vml.StrokeLineStyle(0):              "",
		vml.StrokeLineStyleSingle:           vml.StrokeLineStyleSingle.String(),
		vml.StrokeLineStyleThickBetweenThin: vml.StrokeLineStyleThickBetweenThin.String(),
		vml.StrokeLineStyleThickThin:        vml.StrokeLineStyleThickThin.String(),
		vml.StrokeLineStyleThinThin:         vml.StrokeLineStyleThinThin.String(),
		vml.StrokeLineStyleThinThick:        vml.StrokeLineStyleThinThick.String(),
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
