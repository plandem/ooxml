// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package index

import "fmt"

//Index is helper type for internal indexes
type Index struct {
	idx map[Code]int
}

//Add object's hash to index
func (i *Index) Add(o Indexer, idx int) error {
	c := o.Hash()
	if _, ok := i.idx[c]; ok {
		return fmt.Errorf("there is already object with same hash and index=%d", idx)
	}

	if i.idx == nil {
		i.idx = make(map[Code]int)
	}

	i.idx[o.Hash()] = idx
	return nil
}

//Remove object's hash from index
func (i *Index) Remove(o Indexer) {
	delete(i.idx, o.Hash())
}

//Get object's hash and state - same as regular map result
func (i *Index) Get(o Indexer) (idx int, ok bool) {
	idx, ok = i.idx[o.Hash()]
	return
}

//Has returns true if object's hash was found
func (i *Index) Has(o Indexer) bool {
	_, ok := i.idx[o.Hash()]
	return ok
}

//Count returns total number of hashed objects in index
func (i *Index) Count() int {
	return len(i.idx)
}
