// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package index

import (
	"hash/fnv"
	"strconv"
)

//Code is type for result of Hash method to unify logic around
type Code uint64

//Indexer is interface for objects that will be used to build indexes
type Indexer interface {
	Hash() Code
}

//Hash returns FNV1 hash of string key
func Hash(key string) Code {
	h := fnv.New64a()
	h.Write([]byte(key))
	return Code(h.Sum64())
}

//String return string version (16-base) of hash code
func (c Code) String() string {
	return "0x" + strconv.FormatUint(uint64(c), 16)
}
