package main

import (
	"encoding/json"
	"fmt"

	"github.com/marcofranssen/go-enum-tutorial/car"
)

func main() {
	bmw := car.New(car.BMW, car.Gray, "320i GT")
	ferrari := car.New(car.Ferrari, car.Red, "SF90 Stadala")

	fmt.Printf("I own a '%s' and dream about a '%s'…\n", bmw, ferrari)

	cars := []*car.Car{bmw, ferrari}
	carsJSON, err := json.Marshal(cars)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(carsJSON))

	porscheJSON := []byte(`{"brand":"porsche","color":"White","model":"Taycan"}`)
	var porsche car.Car
	err = json.Unmarshal(porscheJSON, &porsche)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Another Car I love: %s\n", porsche)
}
