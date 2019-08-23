package chart

import (
	"encoding/xml"
	"github.com/plandem/ooxml/drawing/dml"
	"github.com/plandem/ooxml/ml"
)

//Title is a direct mapping of XSD CT_Tx
type Text struct {
	Rich *dml.TextBody `xml:"rich,omitempty"`
	ml.ReservedElements
}

func (n *Text) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n.ReservedElements.ResolveNamespacePrefixes()
	return e.Encode(*n)
}
