package css

type visibility byte

const (
	VisibilityInherit visibility = iota
	VisibilityHidden
	VisibilityVisible
	VisibilityCollapse
)

var (
	toVisibility   map[string]visibility
	fromVisibility map[visibility]string
)

func init() {
	fromVisibility = map[visibility]string{
		VisibilityInherit:  "inherit",
		VisibilityHidden:   "hidden",
		VisibilityVisible:  "visible",
		VisibilityCollapse: "collapse",
	}

	toVisibility = make(map[string]visibility, len(fromVisibility))
	for k, v := range fromVisibility {
		toVisibility[v] = k
	}
}

func (e visibility) String() string {
	return fromVisibility[e]
}
