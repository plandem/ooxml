package vml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStrokeDashStyle(t *testing.T) {
	type Entity struct {
		Attribute vml.StrokeDashStyle `xml:"attribute,attr,omitempty"`
	}

	list := map[vml.StrokeDashStyle]string{
		vml.StrokeDashStyle(0):             "",
		vml.StrokeDashStyleDash:            vml.StrokeDashStyleDash.String(),
		vml.StrokeDashStyleDashDot:         vml.StrokeDashStyleDashDot.String(),
		vml.StrokeDashStyleDot:             vml.StrokeDashStyleDot.String(),
		vml.StrokeDashStyleLongDash:        vml.StrokeDashStyleLongDash.String(),
		vml.StrokeDashStyleLongDashDot:     vml.StrokeDashStyleLongDashDot.String(),
		vml.StrokeDashStyleLongDashDotDot:  vml.StrokeDashStyleLongDashDotDot.String(),
		vml.StrokeDashStyleShortDash:       vml.StrokeDashStyleShortDash.String(),
		vml.StrokeDashStyleShortDashDot:    vml.StrokeDashStyleShortDashDot.String(),
		vml.StrokeDashStyleShortDashDotDot: vml.StrokeDashStyleShortDashDotDot.String(),
		vml.StrokeDashStyleShortDot:        vml.StrokeDashStyleShortDot.String(),
		vml.StrokeDashStyleSolid:           vml.StrokeDashStyleSolid.String(),
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
