package main

import (
	"fmt"

	"github.com/marcofranssen/go-enum-tutorial/car"
)

func main() {
	bmw := car.New(car.BMW, car.Gray, "320i GT")
	ferrari := car.New(car.Ferrari, car.Red, "SF90 Stadala")

	fmt.Printf("I own a '%s' and dream about a '%s'…\n", bmw, ferrari)
}
