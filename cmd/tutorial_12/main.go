package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}
type electricEngine struct {
	kmh   float32
	mpkmh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Tayoutaaa ya bachaaa",
		carModel: "DMAX W RA5 LE",
		engine: gasEngine{
			gallons: 15,
			mpg:     40,
		},
	}
	var electricCar = car[electricEngine]{
		carMake:  "Tesla wila chnw",
		carModel: "MModel S",
		engine: electricEngine{
			kmh:   88,
			mpkmh: 12,
		},
	}
	fmt.Println("Gas Car: ", gasCar)
	fmt.Println("Electric Car: ", electricCar)
	fmt.Println("Gas Car Engine: ", gasCar.engine)
}
