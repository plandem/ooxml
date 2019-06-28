package vml

type horizontalRuler struct {
	HR        bool    `xml:"hr,attr,omitempty" namespace:"o"`
	HRStd     bool    `xml:"hrstd,attr,omitempty" namespace:"o"`
	HRNoShade bool    `xml:"hrnoshade,attr,omitempty" namespace:"o"`
	HRPct     float64 `xml:"hrpct,attr,omitempty" namespace:"o"`
	HRAlign   string  `xml:"hralign,attr,omitempty" namespace:"o"` //enum
}

type borderColor struct {
	BorderTopColor    string `xml:"bordertopcolor,attr,omitempty" namespace:"o"`
	BorderLeftColor   string `xml:"borderleftcolor,attr,omitempty" namespace:"o"`
	BorderBottomColor string `xml:"borderbottomcolor,attr,omitempty" namespace:"o"`
	BorderRightColor  string `xml:"borderrightcolor,attr,omitempty" namespace:"o"`
}

type diagramAttributes struct {
	DgmNodeKind  int  `xml:"dgmnodekind,attr,omitempty" namespace:"o"`
	DgmLayout    byte `xml:"dgmlayout,attr,omitempty" namespace:"o"`
	DgmLayoutMru byte `xml:"dgmlayoutmru,attr,omitempty" namespace:"o"`
}

//coreOfficeAttributes is direct mapping of AG_OfficeCoreAttributes
type coreOfficeAttributes struct {
	OnEd              bool   `xml:"oned,attr,omitempty" namespace:"o"`
	DoubleClickNotify bool   `xml:"doubleclicknotify,attr,omitempty" namespace:"o"`
	Button            bool   `xml:"button,attr,omitempty" namespace:"o"`
	UserHidden        bool   `xml:"userhidden,attr,omitempty" namespace:"o"`
	Bullet            bool   `xml:"bullet,attr,omitempty" namespace:"o"`
	AllowInCell       bool   `xml:"allowincell,attr,omitempty" namespace:"o"`
	AllowOverlap      bool   `xml:"allowoverlap,attr,omitempty" namespace:"o"`
	UserDrawn         bool   `xml:"userdrawn,attr,omitempty" namespace:"o"`
	RegroupID         int    `xml:"regroupid,attr,omitempty" namespace:"o"`
	InsetMode         string `xml:"insetmode,attr,omitempty" namespace:"o"` //enum
	SpID              string `xml:"spid,attr,omitempty" namespace:"o"`
	horizontalRuler
	borderColor
	diagramAttributes
}
