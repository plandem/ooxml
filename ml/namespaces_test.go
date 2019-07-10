// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml_test

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNamespaces(t *testing.T) {
	require.Equal(t, []xml.Attr{
		{Name: xml.Name{Local: "xmlns:v"}, Value: ml.NamespaceVML},
		{Name: xml.Name{Local: "xmlns:o"}, Value: ml.NamespaceVMLOffice},
		{Name: xml.Name{Local: "xmlns:p"}, Value: ml.NamespaceVMLPowerPoint},
		{Name: xml.Name{Local: "xmlns:w"}, Value: ml.NamespaceVMLWord},
		{Name: xml.Name{Local: "xmlns:x"}, Value: ml.NamespaceVMLExcel},
		{Name: xml.Name{Local: "xmlns:r"}, Value: ml.NamespaceRelationships},
	}, ml.Namespaces(
		"unknown",
		ml.NamespaceVML,
		ml.NamespaceVMLOffice,
		ml.NamespaceVMLPowerPoint,
		ml.NamespaceVMLWord,
		ml.NamespaceVMLExcel,
		ml.NamespaceRelationships,
	))
}
