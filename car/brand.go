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
