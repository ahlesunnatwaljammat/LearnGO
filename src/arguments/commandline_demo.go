package main

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println(len(os.Args))
	if len(os.Args) != 3 {
		log.SetOutput(os.Stderr)
		log.Fatalln("Please provide username and password")
		os.Exit(-1)
	}

	if os.Args[1] == "nabbasi" && os.Args[2] == "x" {
		log.Println("Welcome, ", os.Args[0])
	}
}
