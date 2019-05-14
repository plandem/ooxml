package css_test

import (
	"github.com/plandem/ooxml/vml/css"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCSS(t *testing.T) {
	data := "position:absolute;margin-left:59.25pt;margin-top:1.50pt;width:96px;height:55px;z-index:1;visibility:hidden"
	decoded := css.Decode(data)

	require.Equal(t, css.Style{
		Position: css.PositionAbsolute,
		MarginLeft: 59.25,
		MarginTop: 1.5,
		Width: int64(96),
		Height: int64(55),
		ZIndex: 1,
		Visible: css.VisibilityHidden,
	}, decoded)

	encoded := decoded.Encode()
	require.Equal(t, data, encoded)
	require.Equal(t, decoded, css.Decode(encoded))
}
