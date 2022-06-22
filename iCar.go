package main

type iCar interface {
	setCarName(yourCarname string)
	setModelName(yourModelName string)
	setProductionYear(yourProductionYear int)
	setCarColor(yourCarColor string)
	setEngineType(yourEngineType string)
	setEngineTopSpeed(yourEngineTopSpeed int)
	setTiresBrand(yourTiresBrand string)
	setTiresSize(yourTiresSize int)
	doAssemble()
	showProductionResult() string
}
