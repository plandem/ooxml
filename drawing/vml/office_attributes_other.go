package vml

//type AltHRef string         //`xml:"althref,attr,omitempty" namespace:"o"`
//type DetectMouseClick *bool //`xml:"detectmouseclick,attr,omitempty" namespace:"o"`
//type HRef string            //`xml:"href,attr,omitempty" namespace:"o"`
//type Opacity2 string        //`xml:"opacity2,attr,omitempty" namespace:"o"`
//type RelID string           //`xml:"relid,attr,omitempty" namespace:"o"`
//type Title string           //`xml:"title,attr,omitempty" namespace:"o"`
//type Movie float64          //`xml:"movie,attr,omitempty" namespace:"o"`
//
//type ConnectType string       //`xml:"connecttype,attr,omitempty" namespace:"o"`
//type ConnectLocs string       //`xml:"connectlocs,attr,omitempty" namespace:"o"`
//type ConnectAngles string     //`xml:"connectangles,attr,omitempty" namespace:"o"`
//type ExtrusionOK *bool        //`xml:"extrusionok,attr,omitempty" namespace:"o"`
//type ForceDash *bool          //`xml:"forcedash,attr,omitempty" namespace:"o"`
//type SingleClick bool         //`xml:"singleclick,attr,omitempty" namespace:"o"`
//type InsetMode string         //`xml:"insetmode,attr,omitempty" namespace:"o"` //enum
//type SpID string              //`xml:"spid,attr,omitempty" namespace:"o"`
//type OnEd bool                //`xml:"oned,attr,omitempty" namespace:"o"`
//type RegroupID int            //`xml:"regroupid,attr,omitempty" namespace:"o"`
//type DoubleClickNotify bool   //`xml:"doubleclicknotify,attr,omitempty" namespace:"o"`
//type Button bool              //`xml:"button,attr,omitempty" namespace:"o"`
//type UserHidden bool          //`xml:"userhidden,attr,omitempty" namespace:"o"`
//type Bullet bool              //`xml:"bullet,attr,omitempty" namespace:"o"`
//type HR bool                  //`xml:"hr,attr,omitempty" namespace:"o"`
//type HRStd bool               //`xml:"hrstd,attr,omitempty" namespace:"o"`
//type HRNoShade bool           //`xml:"hrnoshade,attr,omitempty" namespace:"o"`
//type HRPct float64            //`xml:"hrpct,attr,omitempty" namespace:"o"`
//type HRAlign string           //`xml:"hralign,attr,omitempty" namespace:"o"` //enum
//type AllowInCell bool         //`xml:"allowincell,attr,omitempty" namespace:"o"`
//type AllowOverlap bool        //`xml:"allowoverlap,attr,omitempty" namespace:"o"`
//type UserDrawn bool           //`xml:"userdrawn,attr,omitempty" namespace:"o"`
//type BorderTopColor string    //`xml:"bordertopcolor,attr,omitempty" namespace:"o"`
//type BorderLeftColor string   //`xml:"borderleftcolor,attr,omitempty" namespace:"o"`
//type BorderBottomColor string //`xml:"borderbottomcolor,attr,omitempty" namespace:"o"`
//type BorderRightColor string  //`xml:"borderrightcolor,attr,omitempty" namespace:"o"`
//type DgmLayout byte           //`xml:"dgmlayout,attr,omitempty" namespace:"o"`
//type DgmNodeKind int          //`xml:"dgmnodekind,attr,omitempty" namespace:"o"`
//type DgmLayoutMru byte        //`xml:"dgmlayoutmru,attr,omitempty" namespace:"o"`
//type Master string //`xml:"master,attr,omitempty" namespace:"o"`
//type Spt            string //`xml:"spt,attr,omitempty" namespace:"o"`
//type ConnectorType  string //`xml:"connectortype,attr,omitempty" namespace:"o"` //enum
//type BWMode         string //`xml:"bwmode,attr,omitempty" namespace:"o"`        //enum
//type BWPure         string //`xml:"bwpure,attr,omitempty" namespace:"o"`        //enum
//type BWNormal       string //`xml:"bwnormal,attr,omitempty" namespace:"o"`      //enum
//type ForceDash      bool   //`xml:"forcedash,attr,omitempty" namespace:"o"`
//type OLEIcon        bool   //`xml:"oleicon,attr,omitempty" namespace:"o"`
//type OLE            bool   //`xml:"ole,attr,omitempty" namespace:"o"`
//type PreferRelative bool   //`xml:"preferrelative,attr,omitempty" namespace:"o"`
//type ClipToWrap     bool   //`xml:"cliptowrap,attr,omitempty" namespace:"o"`
//type Clip           bool   //`xml:"clip,attr,omitempty" namespace:"o"`

//
//func (v AltHRef) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: string(v)}, nil
//}
//
//func (v HRef) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: string(v)}, nil
//}
//
//func (v DetectMouseClick) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: "???"}, nil
//}
//
//func (v Opacity2) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: string(v)}, nil
//}
//
//func (v RelID) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: string(v)}, nil
//}
//
//func (v Title) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: string(v)}, nil
//}
//
//func (v Movie) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: strconv.FormatFloat(float64(v), 'g', -1, 64)}, nil
//}
//
//func (v ConnectType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: string(v)}, nil
//}
//
//func (v ConnectLocs) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: string(v)}, nil
//}
//
//func (v ConnectAngles) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: string(v)}, nil
//}
//
//func (v ExtrusionOK) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
//	return xml.Attr{Name: applyNamespacePrefix("o", name), Value: "???"}, nil
//}
//
////func (v HRef) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
////	return xml.Attr{Name: name, Value: string(v)}, nil
////}
