package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("channelAsReturn: %s\n", <-channelAsReturn("This is a message"))
	}
}

func channelAsReturn(message string) <-chan string {
	c := make(chan string)

	go func() {
		for j := 0; j < 10; j++ {
			c <- fmt.Sprintf("%s %d", message, j)
			time.Sleep(time.Duration(1 * time.Second))
		}
	}()

	return c
}
