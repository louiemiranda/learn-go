package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var status = "processing"

// const
const (
	API = "http://noah.up.edu.ph/api/doppler"
)

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
