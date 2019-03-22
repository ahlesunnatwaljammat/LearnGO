package main

import (
	"log"
	"os"
)

func byval(y int) int {
	y = y + 1
	return y
}

func byref(y *int) int {
	*y = *y + 1
	return *y
}

func main() {
	log.SetOutput(os.Stdout)

	x := 5
	byval(x)
	log.Println(x)

	byref(&x)
	log.Println(x)
}
