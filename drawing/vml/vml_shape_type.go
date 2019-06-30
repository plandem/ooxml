package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"strconv"
)

//ShapeType is direct mapping of CT_ShapeType
type ShapeType struct {
	XMLName xml.Name `xml:"shapetype"`
	Path    string   `xml:"path,attr,omitempty"`
	shapeAttributes
	shapeElements
}

func (s *ShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	resolveAttributesName(s.ReservedAttributes)

	//spt has namespace, so better to manually encode it, than create a special type for it
	if s.Spt > 0 {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  ooxml.ApplyNamespacePrefix(NamespaceOffice, xml.Name{Local: "spt"}),
			Value: strconv.Itoa(int(s.Spt)),
		})
	}

	resolveElementsName(s.ReservedElements)
	return e.EncodeElement(*s, xml.StartElement{Name: ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, start.Name)})
}
