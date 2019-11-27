package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API = "http://noah.up.edu.ph/api/doppler"
)

func main() {
	println("Processing...")

	retrieveRaw()
}

func retrieveRaw() {

	res, err := http.Get(API)
	if err != nil {
		panic(err)
	}
	println(res)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	println(res.StatusCode)
	fmt.Printf("%s\n", body)

}
