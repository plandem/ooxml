package vml

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strings"
)

//ShapeElements is direct mapping of EG_ShapeElements
type ShapeElements struct {
	Path          *Reserved `xml:"path,omitempty"`
	Formulas      *Reserved `xml:"formulas,omitempty"`
	Handles       *Reserved `xml:"handles,omitempty"`
	Fill          *Reserved `xml:"fill,omitempty"`
	Stroke        *Reserved `xml:"stroke,omitempty"`
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

type Shape = ShapeType

//func (r *CoreOfficeAttributes) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{}, nil
//}

//func marshalNamedAttributes() error {
//
//}
type fieldFlags int

const (
	fElement fieldFlags = 1 << iota
	fAttr
	fCDATA
	fCharData
	fInnerXml
	fComment
	fAny

	fOmitEmpty

	fMode = fElement | fAttr | fCDATA | fCharData | fInnerXml | fComment | fAny
)

func getStart(t reflect.Type) xml.StartElement {
	start := xml.StartElement{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("xml")

		if (f.PkgPath != "" && !f.Anonymous) || tag == "-" {
			continue // Private field
		}

		tokens := strings.Split(tag, ",")
		if f.Name == "XMLName" {
			start.Name.Local = tokens[0]

			if namespace := f.Tag.Get("namespace"); len(namespace) > 0 {
				start.Name.Local = namespace + ":" + start.Name.Local
			}
		}
	}

	return start
}

//func addNameSpace(value interface{}, start *xml.StartElement) {
//	v := reflect.Indirect(reflect.ValueOf(value))
//	vt := v.Type()
//
//	for i := 0; i < v.NumField(); i++ {
//		f := vt.Field(i)
//
//		if f.Name == "XMLName" {
//			if namespace := f.Tag.Get("namespace"); len(namespace) > 0 {
//				start.Name.Local = namespace + ":" + start.Name.Local
//			}
//		}
//	}
//}

func marshalNamedElement(value interface{}, e *xml.Encoder) error {
	//if len(namespace) > 0 {
	//	start.Name.Local = namespace + ":" + start.Name.Local
	//}

	v := reflect.Indirect(reflect.ValueOf(value))

	if !v.IsValid() {
		return nil
	}

	vt := v.Type()

	start := getStart(vt)
	err := e.EncodeToken(start)
	_ = err

	for i := 0; i < v.NumField(); i++ {
		f := vt.Field(i)
		tag := f.Tag.Get("xml")

		if (f.PkgPath != "" && !f.Anonymous) || tag == "-" {
			continue // Private field
		}

		flags := fieldFlags(0)
		tokens := strings.Split(tag, ",")
		if len(tokens) == 1 {
			flags = fElement
		} else {
			tag = tokens[0]
			for _, flag := range tokens[1:] {
				switch flag {
				case "attr":
					flags |= fAttr
				case "cdata":
					flags |= fCDATA
				case "chardata":
					flags |= fCharData
				case "innerxml":
					flags |= fInnerXml
				case "comment":
					flags |= fComment
				case "any":
					flags |= fAny
				case "omitempty":
					flags |= fOmitEmpty
				}
			}
		}


		//switch vtt := vt.(type) {
		//
		//}
		//

		fmt.Println(i, tag, tokens, f.Name)
	}

	return e.EncodeToken(start.End())
}
//
//func encodeNamedElement(v interface{}, e *xml.Encoder, start xml.StartElement) error {
//	addNameSpace(v, &start)
//	fmt.Println(start.Attr)
//	return e.EncodeElement(v, start)
//}

func (s *ShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return marshalNamedElement(s, e)
	//addNameSpace(s, &start)
	//return e.EncodeElement(*s, start)
	//fmt.Printf("%+v", s)
	//return encodeNamedElement(*s, e, start)
}

func (r *CoreOfficeAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Space = "o"
	fmt.Println("CoreOfficeAttributes", start.Attr)
	return nil
}

func (r *ShapeOfficeAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Space = "o"
	fmt.Println("ShapeOfficeAttributes", start.Attr)
	return nil
}

func (r *CoreAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	fmt.Println("CoreAttributes", start.Attr)
	return nil
}

func (r *ShapeAttributes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	fmt.Println("ShapeAttributes", start.Attr)
	return nil
}

////Shape is direct mapping of CT_Shape
//type Shape struct {
//	ShapeType
//	Type string
//}
//
////Group is direct mapping of CT_Group
//type Group struct {
//}
//
////Rect is direct mapping of CT_Rect
//type Rect struct {
//	Shape
//}
//
////RoundRect is direct mapping of CT_RoundRect
//type RoundRect struct {
//	Shape
//}
//
////Line is direct mapping of CT_Line
//type Line struct {
//	Shape
//}
//
////PolyLine is direct mapping of CT_PolyLine
//type PolyLine struct {
//	Shape
//}
//
////Curve is direct mapping of CT_Curve
//type Curve struct {
//	Shape
//}
//
////Oval is direct mapping of CT_Oval
//type Oval struct {
//	Shape
//}
//
////Arc is alias of CT_Arc
//type Arc struct {
//	Shape
//}
//
////Image is direct mapping of CT_Image
//type Image struct {
//	Shape
//}
//
////Diagram is direct mapping of CT_Diagram
//type Diagram struct {
//	Shape
//}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
