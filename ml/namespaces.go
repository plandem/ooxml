// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml

import (
	"encoding/xml"
	"fmt"
)

const (
	//NamespaceXML is a XML namespace
	NamespaceXML = "http://www.w3.org/XML/1998/namespace"

	//NamespaceRelationships is a namespace for OOXML relationships
	NamespaceRelationships = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"

	//NamespaceVML is general VML drawings namespace
	NamespaceVML = "urn:schemas-microsoft-com:vml"

	//NamespaceVMLOffice is general VML Office drawings namespace
	NamespaceVMLOffice = "urn:schemas-microsoft-com:office:office"

	//NamespaceVMLExcel is Excel related VML drawing namespace
	NamespaceVMLExcel = "urn:schemas-microsoft-com:office:excel"

	//NamespaceVMLWord is Word related VML drawing namespace
	NamespaceVMLWord = "urn:schemas-microsoft-com:office:word"

	//NamespaceVMLPowerPoint is PowerPoint related VML drawing namespace
	NamespaceVMLPowerPoint = "urn:schemas-microsoft-com:office:powerpoint"
)

var (
	namespacePrefixes map[string]string
)

func init() {
	namespacePrefixes = map[string]string{
		NamespaceRelationships: "r",
		NamespaceVML:           "v",
		NamespaceVMLOffice:     "o",
		NamespaceVMLExcel:      "x",
		NamespaceVMLWord:       "w",
		NamespaceVMLPowerPoint: "p",
	}
}

func errorNamespace(namespace string) error {
	return fmt.Errorf("can't resolve prefix for: %s", namespace)
}

//ApplyNamespacePrefix adds namespace prefix to Local name and drops Space name
func ApplyNamespacePrefix(namespace string, name xml.Name) xml.Name {
	if prefix, ok := namespacePrefixes[namespace]; ok {
		return xml.Name{
			Local: prefix + ":" + name.Local,
		}
	}

	return name
}

//Namespaces transform list of namespaces into list of related attributes
func Namespaces(namespaces ...string) []xml.Attr {
	attrs := make([]xml.Attr, 0, len(namespaces))

	for _, namespace := range namespaces {
		if prefix, ok := namespacePrefixes[namespace]; ok {
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "xmlns:" + prefix}, Value: namespace})
		}
	}

	return attrs
}
