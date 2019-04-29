package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text (bufio.NewReader) : ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text (bufio.NewScanner) : ")
	scanner.Scan()
	fmt.Println(scanner.Text())

	fmt.Print("Enter text (fmt.Scanln(&input)): ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
