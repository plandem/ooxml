package ml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTriStateAttrPersist(t *testing.T) {
	type Element struct {
		State ml.TriStateType `xml:"state,attr"`
	}

	entity := Element{}
	encoded, err := xml.Marshal(&entity)
	require.Empty(t, err)
	require.Equal(t, `<Element state=""></Element>`, string(encoded))

	var decoded Element
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)
	require.Equal(t, entity, decoded)
}

func TestTriStateAttr(t *testing.T) {
	type Element struct {
		State ml.TriStateType `xml:"state,attr,omitempty"`
	}

	list := map[ml.TriStateType]string{
		ml.TriStateBlank: ml.TriStateBlank.String(),
		ml.TriStateTrue:  ml.TriStateTrue.String(),
		ml.TriStateFalse: ml.TriStateFalse.String(),
	}

	for k, v := range list {
		t.Run(v, func(tt *testing.T) {
			entity := Element{State: k}
			encoded, err := xml.Marshal(&entity)
			require.Empty(tt, err)

			if k == ml.TriStateBlank {
				require.Equal(tt, `<Element></Element>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Element state="%s"></Element>`, v), string(encoded))
			}

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}

func TestTriStateElementPersist(t *testing.T) {
	type Element struct {
		State ml.TriStateType `xml:"state"`
	}

	entity := Element{}
	encoded, err := xml.Marshal(&entity)
	require.Empty(t, err)
	require.Equal(t, `<Element><state></state></Element>`, string(encoded))

	var decoded Element
	err = xml.Unmarshal(encoded, &decoded)
	require.Empty(t, err)
	require.Equal(t, entity, decoded)
}

func TestTriStateElement(t *testing.T) {
	type Element struct {
		State ml.TriStateType `xml:"state,omitempty"`
	}

	list := map[ml.TriStateType]string{
		ml.TriStateBlank: ml.TriStateBlank.String(),
		ml.TriStateTrue:  ml.TriStateTrue.String(),
		ml.TriStateFalse: ml.TriStateFalse.String(),
	}

	for k, v := range list {
		t.Run(v, func(tt *testing.T) {
			entity := Element{State: k}
			encoded, err := xml.Marshal(&entity)
			require.Empty(tt, err)

			if k == ml.TriStateBlank {
				require.Equal(tt, `<Element></Element>`, string(encoded))
			} else {
				require.Equal(tt, fmt.Sprintf(`<Element><state>%s</state></Element>`, v), string(encoded))
			}

			var decoded Element
			err = xml.Unmarshal(encoded, &decoded)
			require.Empty(tt, err)

			require.Equal(tt, entity, decoded)
		})
	}
}
