package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"strconv"
)

//Spt is helper type to attach office namespace for Spt attribute
type Spt int

//MarshalXMLAttr marshal Spt
func (t Spt) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: ooxml.ApplyNamespacePrefix(NamespaceOfficePrefix, name)}
	attr.Value = strconv.Itoa(int(t))
	return attr, nil
}
