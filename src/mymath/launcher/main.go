package main

import (
	"../../mymath"
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
}
