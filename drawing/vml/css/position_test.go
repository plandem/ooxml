// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package css_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPosition(t *testing.T) {
	type Entity struct {
		Attribute css.Position `xml:"attribute,attr"`
	}

	list := map[css.Position]string{
		css.PositionRelative: css.PositionRelative.String(),
		css.PositionAbsolute: css.PositionAbsolute.String(),
		css.PositionStatic:   css.PositionStatic.String(),
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
