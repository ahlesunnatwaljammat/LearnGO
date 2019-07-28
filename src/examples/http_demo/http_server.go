package main

import (
	"encoding/json"
	"json_demo/json_type"
	"log"
	"net/http"
)

func main() {
	// All URLs will be handled by this function
	// http.HandleFunc uses the DefaultServeMux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		arrayFoo := [...]json_type.Foo{{Number: 1, Title: "noman"}, {Number: 2, Title: "farhan"}, {Number: 3, Title: "arsalan"}}
		bytes, e := json.Marshal(arrayFoo)
		if e == nil {
			w.Write([]byte(bytes))
		}
	})

	// Continue to process new requests until an error occurs
	err := http.ListenAndServe(":8090", nil)

	if err == nil {
		log.Println("Server started at http://localhost:8090")
	}

	log.Fatal(err)
}
