package css_test

import (
	"bytes"
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNumber(t *testing.T) {
	//ints by default is in px
	require.Equal(t, css.NewNumber(1, css.UnitPx), css.NewNumber(1))
	require.Equal(t, css.NewNumber(1, css.UnitPx), css.NewNumber("1"))

	//ints by default is in pt
	require.Equal(t, css.NewNumber(1.0, css.UnitPt), css.NewNumber(1.0))
	require.Equal(t, css.NewNumber(1.0, css.UnitPt), css.NewNumber("1.0"))

	//floats can't be px
	require.Equal(t, css.NewNumber(1.0, css.UnitPt), css.NewNumber(1.0, css.UnitPx))

	//ints can be only px and percentage
	require.Equal(t, css.NewNumber(1, css.UnitPx), css.NewNumber(1, css.UnitPt))

	//typed
	require.Equal(t, css.NewNumber(1.0, css.UnitCm), css.NewNumber("1.0cm"))
	require.Equal(t, css.NewNumber(1.0, css.UnitMm), css.NewNumber("1.0mm"))
	require.Equal(t, css.NewNumber(1.0, css.UnitIn), css.NewNumber("1.0in"))
	require.Equal(t, css.NewNumber(1.0, css.UnitPt), css.NewNumber("1.0pt"))
	require.Equal(t, css.NewNumber(1.0, css.UnitPc), css.NewNumber("1.0pc"))
	require.Equal(t, css.NewNumber(1.0, css.UnitPercentage), css.NewNumber("1.0%"))
	require.Equal(t, css.NewNumber(1, css.UnitPercentage), css.NewNumber("1%"))

	//ints can be non 'px'
	require.Equal(t, css.NewNumber(1.0, css.UnitPt), css.NewNumber("1pt"))

	//'px' can be only ints
	require.Equal(t, css.NewNumber(1.0, css.UnitPt), css.NewNumber("1.0px"))
	require.Equal(t, css.NewNumber(1, css.UnitPx), css.NewNumber("1px"))

	type Entity struct {
		XMLName         xml.Name
		Cm              css.Number `xml:"cm,attr"`
		Mm              css.Number `xml:"mm,attr"`
		In              css.Number `xml:"in,attr"`
		Pt              css.Number `xml:"pt,attr"`
		Pc              css.Number `xml:"pc,attr"`
		Px              css.Number `xml:"px,attr"`
		PercentageInt   css.Number `xml:"perInt,attr"`
		PercentageFloat css.Number `xml:"perFloat,attr"`
		IntPx           css.Number `xml:"intPx,attr"`
		FloatPt         css.Number `xml:"floatPt,attr"`
	}

	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<entity cm="1.1cm" mm="1.2mm" in="1.3in" pt="1.4pt" pc="1.5pc" px="2px" perInt="1%" perFloat="1.1%" intPx="1" floatPt="1.0"/>`)

	//decode
	decoder := xml.NewDecoder(bytes.NewReader([]byte(data)))
	entity := &Entity{}
	err := decoder.DecodeElement(entity, nil)
	require.Nil(t, err)

	require.Equal(t, css.NewNumber(1.1, css.UnitCm), entity.Cm)
	require.Equal(t, css.NewNumber(1.2, css.UnitMm), entity.Mm)
	require.Equal(t, css.NewNumber(1.3, css.UnitIn), entity.In)
	require.Equal(t, css.NewNumber(1.4, css.UnitPt), entity.Pt)
	require.Equal(t, css.NewNumber(1.5, css.UnitPc), entity.Pc)
	require.Equal(t, css.NewNumber(2, css.UnitPx), entity.Px)
	require.Equal(t, css.NewNumber(1, css.UnitPercentage), entity.PercentageInt)
	require.Equal(t, css.NewNumber(1.1, css.UnitPercentage), entity.PercentageFloat)
	require.Equal(t, css.NewNumber(1), entity.IntPx)
	require.Equal(t, css.NewNumber(1.0), entity.FloatPt)

	//encode
	encoded, err := xml.Marshal(&entity)
	require.Nil(t, err)

	require.Equal(t, strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
<entity cm="1.1cm" mm="1.2mm" in="1.3in" pt="1.4pt" pc="1.5pc" px="2px" perInt="1%" perFloat="1.1%" intPx="1px" floatPt="1pt"></entity>`), string(encoded))
}
