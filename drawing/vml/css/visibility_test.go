package css_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVisibility(t *testing.T) {
	type Entity struct {
		Attribute css.Visibility `xml:"attribute,attr"`
	}

	list := map[css.Visibility]string{
		css.VisibilityHidden:   css.VisibilityHidden.String(),
		css.VisibilityInherit:  css.VisibilityInherit.String(),
		css.VisibilityVisible:  css.VisibilityVisible.String(),
		css.VisibilityCollapse: css.VisibilityCollapse.String(),
	}

	for k, v := range list {
		t.Run(v, func(tt *testing.T) {
			entity := Entity{Attribute: k}
			encoded, err := xml.Marshal(&entity)

			require.Empty(tt, err)
			require.Equal(tt, fmt.Sprintf(`<Entity attribute="%s"></Entity>`, v), string(encoded))

			var decoded Entity
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
			require.Equal(tt, v, decoded.Attribute.String())
		})
	}
}
