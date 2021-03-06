package main

import (
	"fmt"
	"github.com/alexbrainman/printer"
	"os"
	"strings"
)

func main() {
	printers, _ := printer.ReadNames()

	defaultPrinter, _ := printer.Default()

	for i, printerName := range printers {
		s := " "
		if printerName == defaultPrinter {
			s = "*"
		}
		fmt.Printf(" %s %d. %s\n", s, i, printerName)
	}

	p, _ := printer.Open("HP LaserJet P205X series PCL6")
	defer p.Close()

	_ = p.StartRawDocument("Test.txt")
	defer p.EndDocument()

	_ = p.StartPage()
	lines := strings.Split("This a \ntest \ndocument", "\n")
	for _, line := range lines {
		fmt.Fprintf(p, "%s\r\n", line)
	}

	os.Exit(0)
}
