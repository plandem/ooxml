package vml_test

import (
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnchorClientData(t *testing.T) {
	a, err := vml.StringToClientDataAnchor("1, 15, 0, 2, 3, 15, 3, 16")
	require.Nil(t, err)
	require.Equal(t, vml.ClientDataAnchor{
		LeftColumn:   1,
		LeftOffset:   15,
		TopRow:       0,
		TopOffset:    2,
		RightColumn:  3,
		RightOffset:  15,
		BottomRow:    3,
		BottomOffset: 16,
	}, a)

	require.Equal(t, "1, 15, 0, 2, 3, 15, 3, 16", a.String())

	a, err = vml.StringToClientDataAnchor(" 1,15,0,2, 3, 15,3, 16 ")
	require.Nil(t, err)
	require.Equal(t, vml.ClientDataAnchor{
		LeftColumn:   1,
		LeftOffset:   15,
		TopRow:       0,
		TopOffset:    2,
		RightColumn:  3,
		RightOffset:  15,
		BottomRow:    3,
		BottomOffset: 16,
	}, a)

	_, err = vml.StringToClientDataAnchor(" x,15,0,2, 3, 15,3, 16 ")
	require.NotNil(t, err)
}
