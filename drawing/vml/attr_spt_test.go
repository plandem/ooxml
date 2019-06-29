package vml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSpt(t *testing.T) {
	type Entity struct {
		Attribute vml.Spt `xml:"attribute,attr,omitempty"`
	}

	//empty
	entity := Entity{}
	encoded, err := xml.Marshal(&entity)
	require.Empty(t, err)
	require.Equal(t, `<Entity></Entity>`, string(encoded))

	//with value
	entity = Entity{Attribute: 202}
	encoded, err = xml.Marshal(&entity)

	require.Empty(t, err)
	require.Equal(t, fmt.Sprintf(`<Entity o:attribute="%d"></Entity>`, 202), string(encoded))

	var decoded Entity
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)

	require.Equal(t, entity, decoded)
	require.Equal(t, vml.Spt(202), decoded.Attribute)
}
