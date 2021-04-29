package car

//go:generate stringer -type=Brand
type Brand int

const (
	Unknown Brand = iota
	BMW
	Mercedes
	Audi
	Toyota
	Volkswagen
	Porsche
	Ferrari
)

func (b Brand) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}
