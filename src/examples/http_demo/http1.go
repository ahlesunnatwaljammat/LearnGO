package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"json_demo/json_type"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	resp, err := http.Get("http://localhost:8090")

	if err == nil {
		status := resp.StatusCode
		if status == http.StatusOK {
			bytes, _ := ioutil.ReadAll(resp.Body)
			log.Println(string(bytes))
		}
	}

	jsonResp, err := http.Get("http://localhost:8090/json")

	if err == nil {
		status := jsonResp.StatusCode
		if status == http.StatusOK {
			bytes, _ := ioutil.ReadAll(jsonResp.Body)
			log.Printf("%+v", bytes)

			var foo json_type.Foo
			err := json.Unmarshal(bytes, &foo)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Printf("%+v", foo)
		}
	}
}
