package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

/**
* This program only run in command
 */
func main() {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	fmt.Printf("You pressed: %q\r\n", char)

	erra := keyboard.Open()
	if erra != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			break
		}
		fmt.Printf("You pressed: %q\r\n", char)
	}
}
