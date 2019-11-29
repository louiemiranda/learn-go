package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// const
const (
	API string = "http://noah.up.edu.ph/api/doppler"
)

type res struct {
	body string
}

func init() {
	fmt.Println("learn-go/app/parser -- loaded via init")
	// Retrieve()
}

// Retrieve actual API call
func Retrieve() {

	res, err := http.Get(API)
	if err != nil {
		panic(err)
	}
	// println(res)

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	println("http.status_code: ", res.StatusCode)
	fmt.Printf("%s\n", body)

}
