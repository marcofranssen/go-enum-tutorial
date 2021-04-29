package car

//go:generate stringer -type=Color
type Color int

const (
	White Color = iota
	Gray
	Red
)

func (c Color) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}
