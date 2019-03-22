package main

import (
	"encoding/json"
	"log"
	"os"
)

type Foo struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	foo := &Foo{Number: 1, Title: "abbasi"}
	fooJson, _ := json.Marshal(foo)
	log.Println(string(fooJson))

	fooJsn, _ := json.Marshal(Foo{Number: 1, Title: "noman"})
	log.Println(string(fooJsn))
}
