package mymath

import (
	"fmt"
	"log"
)

func Sum(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

var Divide = func(a int, b int) int {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Divide: ", err)
		}
	}()

	return a / b
}

/*If method name is start with lower it wont be expose outside the package*/
func hideme() {
	fmt.Println("ahhhahah")
}
