package vml

//shapeOfficeAttributes is direct mapping of AG_OfficeShapeAttributes
type shapeOfficeAttributes struct {
	Spt            string `xml:"spt,attr,omitempty" namespace:"o"`
	ConnectorType  string `xml:"connectortype,attr,omitempty" namespace:"o"` //enum
	BWMode         string `xml:"bwmode,attr,omitempty" namespace:"o"`        //enum
	BWPure         string `xml:"bwpure,attr,omitempty" namespace:"o"`        //enum
	BWNormal       string `xml:"bwnormal,attr,omitempty" namespace:"o"`      //enum
	ForceDash      bool   `xml:"forcedash,attr,omitempty" namespace:"o"`
	OLEIcon        bool   `xml:"oleicon,attr,omitempty" namespace:"o"`
	OLE            bool   `xml:"ole,attr,omitempty" namespace:"o"`
	PreferRelative bool   `xml:"preferrelative,attr,omitempty" namespace:"o"`
	ClipToWrap     bool   `xml:"cliptowrap,attr,omitempty" namespace:"o"`
	Clip           bool   `xml:"clip,attr,omitempty" namespace:"o"`
}
