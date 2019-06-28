package relation

import "github.com/plandem/ooxml/ml"

//Relations is type that holds all relation attributes to simplify logic
//FIXME: right now some relations are not working properly, due to conflicting names
type Relations struct {
	Embed       string `xml:"embed,attr,omitempty" namespace:"r"`
	Link        string `xml:"link,attr,omitempty" namespace:"r"`
	Dm          string `xml:"dm,attr,omitempty" namespace:"r"`
	Lo          string `xml:"lo,attr,omitempty" namespace:"r"`
	Qs          string `xml:"qs,attr,omitempty" namespace:"r"`
	Cs          string `xml:"cs,attr,omitempty" namespace:"r"`
	Blip        string `xml:"blip,attr,omitempty" namespace:"r"`
	Pict        string `xml:"pict,attr,omitempty" namespace:"r"`
	TopLeft     string `xml:"topLeft,attr,omitempty" namespace:"r"`
	TopRight    string `xml:"topRight,attr,omitempty" namespace:"r"`
	BottomLeft  string `xml:"bottomLeft,attr,omitempty" namespace:"r"`
	BottomRight string `xml:"bottomRight,attr,omitempty" namespace:"r"`

	//FIXME: 'r:href' conflicts with 'o:href'
	Href string `xml:"href,attr,omitempty" namespace:"r"`

	//FIXME: 'r:id' conflicts with 'id'
	ID ml.RID `xml:"id,attr,omitempty" namespace:"r"`
}
