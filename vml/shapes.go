package vml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/vml/internal"
)

//ShapeElements is direct mapping of EG_ShapeElements
type ShapeElements struct {
	Path          *Path
	Formulas      *Reserved `xml:"formulas,omitempty"`
	Handles       *Reserved `xml:"handles,omitempty"`
	Fill          *Fill
	Stroke        *Stroke
	Shadow        *Reserved `xml:"shadow,omitempty"`
	TextBox       *Reserved `xml:"textbox,omitempty"`
	TextPath      *Reserved `xml:"textpath,omitempty"`
	ImageData     *Reserved `xml:"imagedata,omitempty"`
	Skew          *Reserved `xml:"skew,omitempty" namespace:"o"`
	Extrusion     *Reserved `xml:"extrusion,omitempty" namespace:"o"`
	Callout       *Reserved `xml:"callout,omitempty" namespace:"o"`
	Lock          *Reserved `xml:"lock,omitempty" namespace:"o"`
	ClipPath      *Reserved `xml:"clippath,omitempty" namespace:"o"`
	SignatureLine *Reserved `xml:"signatureline,omitempty" namespace:"o"`
	Wrap          *Reserved `xml:"wrap,omitempty" namespace:"w10"`
	AnchorLock    *Reserved `xml:"anchorlock,omitempty" namespace:"w10"`
	BorderTop     *Reserved `xml:"bordertop,omitempty" namespace:"w10"`
	BorderBottom  *Reserved `xml:"borderbottom,omitempty" namespace:"w10"`
	BorderLeft    *Reserved `xml:"borderleft,omitempty" namespace:"w10"`
	BorderRight   *Reserved `xml:"borderright,omitempty" namespace:"w10"`
	ClientData    *Reserved `xml:"ClientData,omitempty" namespace:"x"`
	TextData      *Reserved `xml:"textdata,omitempty" namespace:"pvml"`
}

//CoreOfficeAttributes is direct mapping of AG_OfficeCoreAttributes
type CoreOfficeAttributes struct {
	SpID              string  `xml:"spid,attr,omitempty" namespace:"o"`
	OnEd              bool    `xml:"oned,attr,omitempty" namespace:"o"`
	RegroupID         int     `xml:"regroupid,attr,omitempty" namespace:"o"`
	DoubleClickNotify bool    `xml:"doubleclicknotify,attr,omitempty" namespace:"o"`
	Button            bool    `xml:"button,attr,omitempty" namespace:"o"`
	UserHidden        bool    `xml:"userhidden,attr,omitempty" namespace:"o"`
	Bullet            bool    `xml:"bullet,attr,omitempty" namespace:"o"`
	HR                bool    `xml:"hr,attr,omitempty" namespace:"o"`
	HRStd             bool    `xml:"hrstd,attr,omitempty" namespace:"o"`
	HRNoShade         bool    `xml:"hrnoshade,attr,omitempty" namespace:"o"`
	HRPct             float64 `xml:"hrpct,attr,omitempty" namespace:"o"`
	HRAlign           string  `xml:"hralign,attr,omitempty" namespace:"o"` //enum
	AllowInCell       bool    `xml:"allowincell,attr,omitempty" namespace:"o"`
	AllowOverlap      bool    `xml:"allowoverlap,attr,omitempty" namespace:"o"`
	UserDrawn         bool    `xml:"userdrawn,attr,omitempty" namespace:"o"`
	BorderTopColor    string  `xml:"bordertopcolor,attr,omitempty" namespace:"o"`
	BorderLeftColor   string  `xml:"borderleftcolor,attr,omitempty" namespace:"o"`
	BorderBottomColor string  `xml:"borderbottomcolor,attr,omitempty" namespace:"o"`
	BorderRightColor  string  `xml:"borderrightcolor,attr,omitempty" namespace:"o"`
	DgmLayout         byte    `xml:"dgmlayout,attr,omitempty" namespace:"o"`
	DgmNodeKind       int     `xml:"dgmnodekind,attr,omitempty" namespace:"o"`
	DgmLayoutMru      byte    `xml:"dgmlayoutmru,attr,omitempty" namespace:"o"`
	InsetMode         string  `xml:"insetmode,attr,omitempty" namespace:"o"` //enum
}

//CoreAttributes is direct mapping of AG_AllCoreAttributes
type CoreAttributes struct {
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
	CoreOfficeAttributes
}

//ShapeOfficeAttributes is direct mapping of AG_OfficeShapeAttributes
type ShapeOfficeAttributes struct {
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

//ShapeAttributes is direct mapping of AG_AllShapeAttributes
type ShapeAttributes struct {
	ChromaKey    string `xml:"chromakey,attr,omitempty"`
	Filled       bool   `xml:"filled,attr,omitempty"`
	FillColor    string `xml:"fillcolor,attr,omitempty"`
	Opacity      string `xml:"opacity,attr,omitempty"` //E.g.: 50%
	Stroked      bool   `xml:"stroked,attr,omitempty"`
	StrokeColor  string `xml:"strokecolor,attr,omitempty"`
	StrokeWeight string `xml:"strokeweight,attr,omitempty"` //E.g.: 1pt
	InsetPen     bool   `xml:"insetpen,attr,omitempty"`
	ShapeOfficeAttributes
}

//ShapeType is direct mapping of CT_ShapeType
type ShapeType struct {
	XMLName xml.Name `xml:"shapetype" namespace:"v"`

	CoreAttributes
	ShapeAttributes
	Master string `xml:"master,attr,omitempty" namespace:"o"`
	Path   string `xml:"path,attr,omitempty"`
	Adj    string `xml:"adj,attr,omitempty"`

	ShapeElements
	Complex *Reserved `xml:"complex,omitempty" namespace:"o"`
}

//Shape is direct mapping of CT_Shape
type Shape struct {
	XMLName xml.Name `xml:"shape" namespace:"v"`
	Type    string   `xml:"type,attr,omitempty"`

	ShapeType
}

//Rect is direct mapping of CT_Rect
type Rect struct {
	XMLName xml.Name `xml:"rect" namespace:"v"`
	Shape
}

//RoundRect is direct mapping of CT_RoundRect
type RoundRect struct {
	XMLName xml.Name `xml:"roundrect" namespace:"v"`
	Shape
}

//Line is direct mapping of CT_Line
type Line struct {
	XMLName xml.Name `xml:"line" namespace:"v"`
	Shape
}

//PolyLine is direct mapping of CT_PolyLine
type PolyLine struct {
	XMLName xml.Name `xml:"polyline" namespace:"v"`
	Shape
}

//Curve is direct mapping of CT_Curve
type Curve struct {
	XMLName xml.Name `xml:"curve" namespace:"v"`
	Shape
}

//Arc is alias of CT_Arc
type Arc struct {
	XMLName xml.Name `xml:"arc" namespace:"v"`
	Shape
}

//Image is direct mapping of CT_Image
type Image struct {
	XMLName xml.Name `xml:"image" namespace:"v"`
	Shape
}

//Oval is direct mapping of CT_Oval
type Oval struct {
	XMLName xml.Name `xml:"oval" namespace:"v"`
	Shape
}

//Group is direct mapping of CT_Group
type Group struct {
	//TODO: implement
}

func (s *ShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *Shape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *Rect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *RoundRect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *Line) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *PolyLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *Curve) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *Arc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *Image) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *Oval) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

func (s *Group) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return internal.Encode(s, e)
}

