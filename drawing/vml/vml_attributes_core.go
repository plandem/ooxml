package vml

//coreAttributes is direct mapping of AG_AllCoreAttributes
type coreAttributes struct {
	ID          string `xml:"id,attr,omitempty"`
	Style       string `xml:"style,attr,omitempty"`
	Href        string `xml:"href,attr,omitempty"`
	Target      string `xml:"target,attr,omitempty"`
	Class       string `xml:"class,attr,omitempty"`
	Title       string `xml:"title,attr,omitempty"`
	Alt         string `xml:"alt,attr,omitempty"`
	CoordSize   string `xml:"coordsize,attr,omitempty"`
	CoordOrigin string `xml:"coordorigin,attr,omitempty"`
	WrapCoords  string `xml:"wrapcoords,attr,omitempty"`
	Print       bool   `xml:"print,attr,omitempty"`
	coreOfficeAttributes
}
