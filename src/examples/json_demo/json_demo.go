package main

import (
	"encoding/json"
	"json_demo/json_type"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	foo := &json_type.Foo{Number: 1, Title: "noman"}
	fooJson, _ := json.Marshal(foo)
	log.Println(string(fooJson))

	fooJsn, _ := json.Marshal(json_type.Foo{Number: 2, Title: "farhan"})
	log.Println(string(fooJsn))
}
