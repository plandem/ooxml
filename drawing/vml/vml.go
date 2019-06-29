package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/ml"
)

type officeDrawing struct {
	ShapeLayout *ShapeLayout `xml:"shapelayout,omitempty"`
	ShapeType   []*ShapeType `xml:"shapetype,omitempty"`
	Shape       []*Shape     `xml:"shape,omitempty"`
	predefinedShapes
	ml.ReservedElements
}

//Excel is type for Excel VML Drawings
type Excel officeDrawing

//Word is type for Word VML Drawings
type Word officeDrawing

//PowerPoint is type for PowerPoint VML Drawings
type PowerPoint officeDrawing

const (
	NamespaceVML        = "urn:schemas-microsoft-com:vml"
	NamespaceOffice     = "urn:schemas-microsoft-com:office:office"
	NamespaceExcel      = "urn:schemas-microsoft-com:office:excel"
	NamespaceWord       = "urn:schemas-microsoft-com:office:word"
	NamespacePowerPoint = "urn:schemas-microsoft-com:office:powerpoint"

	NamespaceVMLPrefix        = "v"
	NamespaceOfficePrefix     = "o"
	NamespaceExcelPrefix      = "x"
	NamespaceWordPrefix       = "w"
	NamespacePowerPointPrefix = "p"
)

//resolveName tries to resolve namespace and apply prefix for it
func resolveName(a xml.Name) xml.Name {
	switch a.Space {
	case NamespaceVML:
		return ooxml.ApplyNamespacePrefix(NamespaceVMLPrefix, a)
	case NamespaceOffice:
		return ooxml.ApplyNamespacePrefix(NamespaceOfficePrefix, a)
	case NamespaceExcel:
		return ooxml.ApplyNamespacePrefix(NamespaceExcelPrefix, a)
	case NamespaceWord:
		return ooxml.ApplyNamespacePrefix(NamespaceWordPrefix, a)
	case NamespacePowerPoint:
		return ooxml.ApplyNamespacePrefix(NamespacePowerPointPrefix, a)
	case ml.NamespaceRelationships:
		return ooxml.ApplyNamespacePrefix(ml.NamespaceRelationshipsPrefix, a)
	}

	return a
}

//resolveAttributesName tries to resolve namespace and apply prefix for it for all reserved attributes
func resolveAttributesName(reserved ml.ReservedAttributes) {
	for i, attr := range reserved.Attrs {
		reserved.Attrs[i].Name = resolveName(attr.Name)
	}
}

//resolveElementsName tries to resolve namespace and apply prefix for it for all reserved elements
func resolveElementsName(nested ml.ReservedElements) {
	for i, node := range nested.Nodes {
		nested.Nodes[i].XMLName = resolveName(node.XMLName)
		resolveAttributesName(node.ReservedAttributes)
	}
}

//attachNamespaces transform list of namespaces into list of related attributes
func attachNamespaces(namespaces ...string) []xml.Attr {
	attrs := make([]xml.Attr, 0, len(namespaces))

	for _, namespace := range namespaces {
		switch namespace {
		case NamespaceVML:
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespaceVMLPrefix}, Value: NamespaceVML})
		case NamespaceOffice:
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespaceOfficePrefix}, Value: NamespaceOffice})
		case NamespaceExcel:
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespaceExcelPrefix}, Value: NamespaceExcel})
		case NamespaceWord:
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespaceWordPrefix}, Value: NamespaceWord})
		case NamespacePowerPoint:
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "xmlns:" + NamespacePowerPointPrefix}, Value: NamespacePowerPoint})
		case ml.NamespaceRelationships:
			attrs = append(attrs, xml.Attr{Name: xml.Name{Local: "xmlns:" + ml.NamespaceRelationshipsPrefix}, Value: ml.NamespaceRelationships})
		}
	}

	return attrs
}

//MarshalXML marshals Excel Drawings
func (o *Excel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "xml"}
	start.Attr = append(start.Attr, attachNamespaces(
		NamespaceVML,
		NamespaceOffice,
		NamespaceExcel,
		ml.NamespaceRelationships,
	)...)
	resolveElementsName(o.ReservedElements)
	return e.EncodeElement(*o, start)
}

//MarshalXML marshals Word Drawings
func (o *Word) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "xml"}
	start.Attr = append(start.Attr, attachNamespaces(
		NamespaceVML,
		NamespaceOffice,
		NamespaceWord,
		ml.NamespaceRelationships,
	)...)
	resolveElementsName(o.ReservedElements)
	return e.EncodeElement(*o, start)
}

//MarshalXML marshals PowerPoint Drawings
func (o *PowerPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "xml"}
	start.Attr = append(start.Attr, attachNamespaces(
		NamespaceVML,
		NamespaceOffice,
		NamespacePowerPoint,
		ml.NamespaceRelationships,
	)...)
	resolveElementsName(o.ReservedElements)
	return e.EncodeElement(*o, start)
}
