package main

import (
	"examples/mymath"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("\uFDF2")
	log.Println(mymath.Sum(1, 2))
	log.Println(mymath.Subtract(1, 2))
	log.Println(mymath.Multiply(1, 2))
	divide := mymath.Divide(1, 0)
	log.Println("Division: ", divide)

	calculator := mymath.New()
	log.Println(calculator.Sum(5, 5))
	log.Println(calculator.Subtract(5, 5))
	log.Println(calculator.Multiply(5, 5))
	log.Println(calculator.Divide(5, 0))
}
