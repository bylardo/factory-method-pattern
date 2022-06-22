package main

type bmw struct {
	car
}

func newBMW() iCar {
	return &bmw{
		car: car{
			Name:           "BMW",
			Model:          "X6",
			ProductionYear: 2020,
			Color:          "Black",
			EngineType:     "Automatic",
			EngineTopSpeed: 340,
			TiresBrand:     "Michelin",
			TiresSize:      265,
		},
	}
}
