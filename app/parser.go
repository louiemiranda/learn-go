package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var projStatus = "processing"

// API string
// var API string = "http://noah.up.edu.ph/api/doppler"

// const
const (
	API string = "http://noah.up.edu.ph/api/doppler"
)

type res struct {
	body string
}

func init() {
	fmt.Println("parser loaded via init")
	retrieve()
}

func retrieve() {

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
