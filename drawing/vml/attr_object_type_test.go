// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package vml_test

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestObjectType(t *testing.T) {
	type Entity struct {
		Attribute vml.ObjectType `xml:"attribute,attr,omitempty"`
	}

	list := map[vml.ObjectType]string{
		vml.ObjectType(0):      "",
		vml.ObjectTypeButton:   vml.ObjectTypeButton.String(),
		vml.ObjectTypeCheckbox: vml.ObjectTypeCheckbox.String(),
		vml.ObjectTypeDialog:   vml.ObjectTypeDialog.String(),
		vml.ObjectTypeDrop:     vml.ObjectTypeDrop.String(),
		vml.ObjectTypeEdit:     vml.ObjectTypeEdit.String(),
		vml.ObjectTypeGBox:     vml.ObjectTypeGBox.String(),
		vml.ObjectTypeLabel:    vml.ObjectTypeLabel.String(),
		vml.ObjectTypeLineA:    vml.ObjectTypeLineA.String(),
		vml.ObjectTypeList:     vml.ObjectTypeList.String(),
		vml.ObjectTypeMovie:    vml.ObjectTypeMovie.String(),
		vml.ObjectTypeNote:     vml.ObjectTypeNote.String(),
		vml.ObjectTypePict:     vml.ObjectTypePict.String(),
		vml.ObjectTypeRadio:    vml.ObjectTypeRadio.String(),
		vml.ObjectTypeRectA:    vml.ObjectTypeRectA.String(),
		vml.ObjectTypeScroll:   vml.ObjectTypeScroll.String(),
		vml.ObjectTypeSpin:     vml.ObjectTypeSpin.String(),
		vml.ObjectTypeShape:    vml.ObjectTypeShape.String(),
		vml.ObjectTypeGroup:    vml.ObjectTypeGroup.String(),
		vml.ObjectTypeRect:     vml.ObjectTypeRect.String(),
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
