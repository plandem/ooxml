package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//Layout is a direct mapping of XSD CT_Layout
type Layout struct {
	ml.ReservedElements
}

func (n *Layout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.EncodeElement(*n, start)
}
