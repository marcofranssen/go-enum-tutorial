package car

import "fmt"

type Car struct {
	Brand Brand  `json:"brand,omitempty"`
	Color Color  `json:"color,omitempty"`
	Model string `json:"model,omitempty"`
}

func New(brand Brand, color Color, model string) *Car {
	return &Car{
		Brand: brand,
		Color: color,
		Model: model,
	}
}

func (c Car) String() string {
	return fmt.Sprintf("%s %s (%s)", c.Brand, c.Model, c.Color)
}
