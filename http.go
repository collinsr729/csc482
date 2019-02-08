package main

import (
	"os"
	"io/ioutil"
	"log"
	"net/http"
	// "loggly.go"
)
	
func main() {
	resp, err := http.Get("https://api.github.com/repos/collinsr729/csc482")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	_, err = os.Stdout.Write(body)

	if err != nil {
		log.Fatal(err)
	}
}