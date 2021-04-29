package car

//go:generate stringer -type=Color
type Color int

const (
	White Color = iota
	Gray
	Red
)
