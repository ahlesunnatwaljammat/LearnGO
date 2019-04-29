package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)

	go boring("boring!", c)

	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
}

func boring(message string, c chan string) {
	for i := 0; true; i++ {
		c <- fmt.Sprintf("%s %d", message, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
