// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/index"
	"strings"
)

//Reserved is special type that catches all inner content AS IS to save original information - used to mark 'non implemented' elements
type Reserved struct {
	XMLName  xml.Name
	InnerXML string `xml:",innerxml"`
	ReservedAttributes
}

//ReservedAttributes is a special type that catches all not captured attributes AS IS to save original information
type ReservedAttributes struct {
	Attrs []xml.Attr `xml:",any,attr"`
}

//ReservedElements is a special type that catches all not captured nested elements AS IS to save original information
type ReservedElements struct {
	Nodes []Reserved `xml:",any"`
}

//ResolveNamespacePrefixes transforms namespaces into namespaces prefixes
func (r ReservedAttributes) ResolveNamespacePrefixes() {
	for i, attr := range r.Attrs {
		r.Attrs[i].Name = ApplyNamespacePrefix(attr.Name.Space, attr.Name)
	}
}

//ResolveNamespacePrefixes tries to resolve namespace and apply prefix for it for all reserved elements
func (r ReservedElements) ResolveNamespacePrefixes() {
	for i, node := range r.Nodes {
		r.Nodes[i].XMLName = ApplyNamespacePrefix(node.XMLName.Space, node.XMLName)
		node.ResolveNamespacePrefixes()
	}
}

//Hash builds hash code for all required values of Reserved to use as unique index
func (r *Reserved) Hash() index.Code {
	reserved := r
	if reserved == nil {
		reserved = &Reserved{}
	}

	result := make([]string, 0, len(reserved.Attrs))
	result = append(result, reserved.InnerXML)

	for _, attr := range reserved.Attrs {
		result = append(result,
			attr.Name.Space,
			attr.Name.Local,
			attr.Value,
		)
	}

	return index.Hash(strings.Join(result, ":"))
}
