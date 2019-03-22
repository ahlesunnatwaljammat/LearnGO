package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"log"
	"os"
	"strconv"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	if atoi, err := strconv.Atoi("-42"); err == nil {
		fmt.Printf("%T, %v\n", atoi, atoi)
	}

	s := strconv.Itoa(-42)
	fmt.Printf("%T, %v\n", s, s)

	if b, err := strconv.ParseBool("true"); err == nil {
		fmt.Printf("%T, %v\n", b, b)
	}

	if f, err := strconv.ParseFloat("3.1415", 64); err == nil {
		fmt.Printf("%T, %v\n", f, f)
	}

	if i, err := strconv.ParseInt("-42", 10, 64); err == nil {
		fmt.Printf("%T, %v\n", i, i)
	}

	if u, err := strconv.ParseUint("42", 10, 64); err == nil {
		fmt.Printf("%T, %v\n", u, u)
	}

	var hex string
	fmt.Print("Please enter hex (0-19| AF | FFFFFF): ")
	fmt.Scan(&hex)

	/*convert hex to decimal*/
	if s, err := strconv.ParseInt(hex, 16, 32); err == nil {
		fmt.Printf("Hex to Dec: %T, %v\n", s, s)
	} else {
		fmt.Println(err)
	}

	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 34); err == nil {
		fmt.Printf("%T, %v\n", s, humanize.Comma(s))
	} else {
		fmt.Println(err)
	}

	num2 := ""
	fmt.Print("Please enter num2: ")
	fmt.Scan(&num2)

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, humanize.Comma(s))
	} else {
		fmt.Println(err)
	}

	if s, err := strconv.ParseInt(num2, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	} else {
		fmt.Println(err)
	}

	os.Exit(0)
}
