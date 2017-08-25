package ml

type ContentType string

type ContentTypes struct {
	XMLName Name `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`

	Overrides []*TypeOverride `xml:"Override"`
	Defaults  []*TypeDefault  `xml:"Default"`
}

type TypeOverride struct {
	PartName    string `xml:",attr"`
	ContentType ContentType `xml:",attr"`
}

type TypeDefault struct {
	Extension   string `xml:",attr"`
	ContentType ContentType `xml:",attr"`
}
