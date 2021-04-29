package car

type Car struct {
	Brand Brand
	Color Color
	Model string
}

func New(brand Brand, color Color, model string) *Car {
	return &Car{
		Brand: brand,
		Color: color,
		Model: model,
	}
}
