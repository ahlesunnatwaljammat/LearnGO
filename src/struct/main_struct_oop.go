package main

import caro "struct/cars"

func main() {
	car := caro.NewCar()
	car.Wheels = 4
	car.HonkTheHorn()
}
