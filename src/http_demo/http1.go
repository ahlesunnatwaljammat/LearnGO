package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	resp, err := http.Get("http://noman/app")

	if err == nil {
		status := resp.StatusCode
		if status == http.StatusOK {
			bytes, _ := ioutil.ReadAll(resp.Body)
			log.Println(string(bytes))
		}
	}
}
