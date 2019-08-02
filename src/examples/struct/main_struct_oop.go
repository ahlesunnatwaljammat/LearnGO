package main

import car "examples/struct/cars"

func main() {
	car := car.NewCar()
	car.Wheels = 4
	car.HonkTheHorn()
}
