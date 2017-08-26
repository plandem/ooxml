package ml

type ContentType string

//ContentTypes is a direct mapping of XSD type
type ContentTypes struct {
	XMLName Name `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`

	Overrides []*TypeOverride `xml:"Override"`
	Defaults  []*TypeDefault  `xml:"Default"`
}

//TypeOverride is a direct mapping of XSD type
type TypeOverride struct {
	PartName    string      `xml:",attr"`
	ContentType ContentType `xml:",attr"`
}

//TypeDefault is a direct mapping of XSD type
type TypeDefault struct {
	Extension   string      `xml:",attr"`
	ContentType ContentType `xml:",attr"`
}
