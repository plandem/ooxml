// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package index_test

import (
	"github.com/plandem/ooxml/index"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

type value int
func (v value) Hash() index.Code {
	return index.Hash(strconv.Itoa(int(v)))
}

func TestIndex(t *testing.T) {
	idx := index.Index{}
	require.Nil(t, idx.Add(value(1), 1))
	require.NotNil(t, idx.Add(value(1), 1))

	idx.Remove(value(1))
	require.Nil(t, idx.Add(value(1), 1))

	i, ok := idx.Get(value(1))
	require.Equal(t, 1, i)
	require.Equal(t, true, ok)

	i, ok = idx.Get(value(2))
	require.Equal(t, 0, i)
	require.Equal(t, false, ok)

}