package css

type position byte

const (
	PositionStatic position = iota
	PositionAbsolute
	PositionRelative
)

var (
	toPosition   map[string]position
	fromPosition map[position]string
)

func init() {
	fromPosition = map[position]string{
		PositionStatic:   "static",
		PositionAbsolute: "absolute",
		PositionRelative: "relative",
	}

	toPosition = make(map[string]position, len(fromPosition))
	for k, v := range fromPosition {
		toPosition[v] = k
	}
}

func (e position) String() string {
	return fromPosition[e]
}
