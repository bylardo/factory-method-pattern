package main

import "fmt"

func main() {
	bmwCar, _ := getCar("bmw")
	showAssembleResult(bmwCar)
}

func showAssembleResult(c iCar) {
	fmt.Println("Assembling Car ...")
	c.doAssemble()
	fmt.Println("---- Result ------ ")
	fmt.Printf("%s ", c.showProductionResult())
	fmt.Println()
}
