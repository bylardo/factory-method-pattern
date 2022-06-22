package main

import "fmt"

func getCar(carType string) (iCar, error) {
	if carType == "bmw" {
		return newBMW(), nil
	}

	return nil, fmt.Errorf("Wrong car type passed")
}
