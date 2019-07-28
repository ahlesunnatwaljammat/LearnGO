package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dirs, e := ioutil.ReadDir("C:")

	if e != nil {
		fmt.Println("Unable to read directories")
		os.Exit(1)
	}

	for _, dir := range dirs {
		fmt.Printf("Name: %s and Size: %d", dir.Name(), dir.Size())
		fmt.Println()
	}
}
