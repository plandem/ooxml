// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package index_test

import (
	"github.com/plandem/ooxml/index"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHash(t *testing.T) {
	s := "just a string"

	//same hash for same value
	require.Equal(t, index.Hash(s), index.Hash(s))

	//different hash for different value
	require.NotEqual(t, index.Hash(s), index.Hash("another string"))

	require.Equal(t, "0xe7bffdf20d5b21ad", index.Hash(s).String())
}
