package main

import "strconv"

type car struct {
	Name           string
	Model          string
	ProductionYear int
	Color          string
	EngineType     string
	EngineTopSpeed int
	TiresBrand     string
	TiresSize      int
	AssembleResult string
}

func (c *car) setCarName(yourCarname string) {
	c.Name = yourCarname
}

func (c *car) setModelName(yourModelName string) {
	c.Model = yourModelName
}

func (c *car) setProductionYear(yourProductionYear int) {
	c.ProductionYear = yourProductionYear
}

func (c *car) setCarColor(yourCarColor string) {
	c.Color = yourCarColor
}

func (c *car) setEngineType(yourEngineType string) {
	c.EngineType = yourEngineType
}

func (c *car) setEngineTopSpeed(yourEngineTopSpeed int) {
	c.EngineTopSpeed = yourEngineTopSpeed
}

func (c *car) setTiresBrand(yourTiresBrand string) {
	c.TiresBrand = yourTiresBrand
}

func (c *car) setTiresSize(yourTiresSize int) {
	c.TiresSize = yourTiresSize
}

func (c *car) doAssemble() {
	c.AssembleResult = "Your Car Name is " + c.Name + " , model " + c.Model + " has a " + c.Color + " color with production year in " + strconv.Itoa(c.ProductionYear) + ". The Enginee is " + c.EngineType + " with " + strconv.Itoa(c.EngineTopSpeed) + " HP top speed" + ". The Tires Brand is " + c.TiresBrand + " with the size " + strconv.Itoa(c.TiresSize)
}

func (c *car) showProductionResult() string {
	return c.AssembleResult
}
