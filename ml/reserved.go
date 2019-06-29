package ml

import "encoding/xml"

//Reserved is special type that catches all inner content AS IS to save original information - used to mark 'non implemented' elements
type Reserved struct {
	XMLName  xml.Name
	InnerXML string `xml:",innerxml"`
	ReservedAttributes
}

//ReservedAttributes is a special type that catches all not captured attributes AS IS to save original information - used to mark 'non implemented' attributes
type ReservedAttributes struct {
	Attrs []xml.Attr `xml:",any,attr"`
}

//ReservedElements is a special type that catches all not captured nested elements AS IS to save original information - used to mark 'non implemented' elements
type ReservedElements struct {
	Nodes []Reserved `xml:",any"`
}
