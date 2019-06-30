package css_test

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCSS(t *testing.T) {
	data := strings.NewReplacer("\t", "", "\n", "", "\r", "").Replace(`
position:absolute;
margin-left:59.25pt;
margin-top:1.5cm;
width:96px;
height:55px;
z-index:1;
visibility:hidden;
flip:x;
font:normal small-caps normal 36pt Arial;
text-decoration:underline;
trim:true;
xscale:false;
mso-fit-shape-to-text:true;
mso-fit-text-to-shape:false;
mso-text-shadow:true;
mso-direction-alt:auto;
mso-layout-flow-alt:auto;
mso-next-textbox:auto;
mso-rotate:60.1;
mso-text-scale:1.5
`)
	decoded := css.NewStyle(data)

	//all fields should be set
	require.Equal(t, &css.Style{
		Position:          css.PositionAbsolute,
		MarginLeft:        css.NewNumber(59.25),
		MarginTop:         css.NewNumber(1.5, css.UnitCm),
		Width:             css.NewNumber(96),
		Height:            css.NewNumber(55),
		ZIndex:            1,
		Visible:           css.VisibilityHidden,
		Flip:              "x",
		Font:              "normal small-caps normal 36pt Arial",
		TextDecoration:    "underline",
		Trim:              ml.TriStateTrue,
		XScale:            ml.TriStateFalse,
		MSODirectionAlt:   "auto",
		MSOFitShapeToText: ml.TriStateTrue,
		MSOFitTextToShape: ml.TriStateFalse,
		MSOTextShadow:     ml.TriStateTrue,
		MSOLayoutFlowAlt:  "auto",
		MSONextTextbox:    "auto",
		MSORotate:         60.1,
		MSOTextScale:      1.5,
	}, decoded)

	encoded := decoded.String()
	require.Equal(t, data, encoded)
	require.Equal(t, decoded, css.NewStyle(encoded))

	//with trimmed whitespaces
	require.Equal(t, &css.Style{
		Width:  css.NewNumber(96),
		Height: css.NewNumber(55),
		Trim: ml.TriStateTrue,
		XScale: ml.TriStateFalse,
	}, css.NewStyle(`width:96px;height:55px;trim:   true; xscale:false   `))

	//only provided values
	require.Equal(t, `width:96px;height:55px`, css.Style{
		Width:  css.NewNumber(96),
		Height: css.NewNumber(55),
	}.String())
}

func TestXml(t *testing.T) {
	type Entity struct {
		Style  css.Style  `xml:"style,attr,omitempty"`
		PStyle *css.Style `xml:"p_style,attr,omitempty"`
	}

	//empty
	entity := Entity{Style: css.Style{}, PStyle: &css.Style{}}
	encoded, err := xml.Marshal(&entity)
	require.Empty(t, err)
	require.Equal(t, `<Entity></Entity>`, string(encoded))

	//encode
	s := css.Style{
		Position:   css.PositionAbsolute,
		MarginLeft: css.NewNumber(59.25),
		MarginTop:  css.NewNumber(1.5),
		Width:      css.NewNumber(96),
		Height:     css.NewNumber(55),
		ZIndex:     1,
		Visible:    css.VisibilityHidden,
	}
	entity = Entity{Style: s, PStyle: &s}
	encoded, err = xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, `<Entity style="position:absolute;margin-left:59.25pt;margin-top:1.5pt;width:96px;height:55px;z-index:1;visibility:hidden" p_style="position:absolute;margin-left:59.25pt;margin-top:1.5pt;width:96px;height:55px;z-index:1;visibility:hidden"></Entity>`, string(encoded))

	//decode
	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
}
